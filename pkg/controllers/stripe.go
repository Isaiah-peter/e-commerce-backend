package controllers

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/models"
	utils "github.com/Isaiah-peter/e-commerce-backend/pkg/util"
	"github.com/stripe/stripe-go"
	"net/http"
)

var (
	NewCharge models.Charge
)

func CreateCharge(w http.ResponseWriter, r *http.Request) {
	utils.ParseBody(r, &NewCharge)
	apikey := ""
	stripe.Key = apikey

}
