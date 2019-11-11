package CardOperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/oms-services/stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
)

func CreateCard(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(request.Body)
	var param *stripe.CardParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	card, err := card.New(param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	bytes, _ := json.Marshal(card)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusCreated)
}
