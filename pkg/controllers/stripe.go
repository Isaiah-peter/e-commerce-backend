package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Isaiah-peter/e-commerce-backend/pkg/models"
	utils "github.com/Isaiah-peter/e-commerce-backend/pkg/util"
)

var (
	NewCharge models.Charge
)

func CreateCharge(w http.ResponseWriter, r *http.Request) {
	charges := models.Charge{}
	utils.ParseBody(r, &charges)
	u := charges.CreateStripe()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
