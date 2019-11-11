package BalanceOperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/oms-services/stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/balance"
)

//GetBalance
func GetBalance(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")
	b, err := balance.Get(nil)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	bytes, _ := json.Marshal(b)
	responseWriter.WriteHeader(http.StatusOK)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
