package controllers

import (
	"encoding/json"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/models"
	utils "github.com/Isaiah-peter/e-commerce-backend/pkg/util"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"net/http"
)

var (
	NewCharge models.Charge
)

func CreateCharge(w http.ResponseWriter, r *http.Request) {
	charges := models.Charge{}
	utils.ParseBody(r, &charges)
	apikey := "sk_test_51IcK2XJ61o5BqUjzBde4uxzQB139FKEl6yAZnQJeVqu5UEihSB89KEqOFrrQPaXXVve38ehd7uoqSGX5tKtD5EiT00dbMVkvW1"
	stripe.Key = apikey
	res, err := charge.New(&stripe.ChargeParams{
		Amount:   stripe.Int64(charges.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Source:   &stripe.SourceParams{Token: stripe.String(charges.Source)}})

	if err != nil {
		// Handle any errors from attempt to charge
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
   w.Write(u)
}