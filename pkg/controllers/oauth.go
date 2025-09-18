package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/models"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/util"
	"github.com/jinzhu/gorm"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	state := utils.GenerateState()

	http.SetCookie(w, &http.Cookie{
		Name: "google_oauth_state",
		Value: state,
		Expires: time.Now().Add(1 * time.Minute),
		HttpOnly: true,
		Secure: true,
		Path: "/",
	})

	url := config.GoogleOAuthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("google_oauth_state")

	if err != nil {
		http.Error(w ,"Missing state cookie", http.StatusBadRequest)
		return
	}

	stateFromCookies := cookie.Value
	stateFromGoogle := r.URL.Query().Get("state")

	if stateFromCookies != stateFromGoogle {
		http.Error(w, "Invalid state parameter", http.StatusUnauthorized)
	}

	code := r.URL.Query().Get("code")

	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)

	if err != nil {
		http.Error(w, "fail to exchange token"+err.Error(), http.StatusBadRequest)
		return
	}

	client := config.GoogleOAuthConfig.Client(context.Background(), token)

	oauthService, _ := oauth2.NewService(context.Background(),option.WithHTTPClient(client))

	userinfo, err := oauthService.Userinfo.Get().Do()

	if err != nil {
		http.Error(w, "Fail to get user info"+err.Error(), http.StatusBadRequest)
		return
	}

	userData, _ := json.MarshalIndent(userinfo, "", " ")
	fmt.Println("Google userinfo ", string(userData))

	var user models.User

	result := db.Where("emails = ?", userinfo.Email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			user = models.User{
				Email: userinfo.Email,
				UserName: userinfo.Name,
				GoogleID: userinfo.Id,
				IsAdmin: false,
			}

			db.Create(&user)
		} else {
			http.Error(w, "Database Error"+result.Error.Error(), http.StatusInternalServerError)
			return
		}
	}

	tokenString := GenerateToken(&user)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-Type", "application/json")
	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user

	json.NewEncoder(w).Encode(resp)
}