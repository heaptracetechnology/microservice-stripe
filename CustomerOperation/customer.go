package CustomerOperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/oms-services/stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

//Create Cusytomer
func CreateCustomer(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(request.Body)
	var param *stripe.CustomerParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	errr := param.SetSource("tok_amex")
	if errr != nil {
		result.WriteErrorResponse(nil, errr)
	}

	cus, err := customer.New(param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	bytes, _ := json.Marshal(cus)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusCreated)
}
