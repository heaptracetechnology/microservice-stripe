package PaymentIntentOperation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/oms-services/stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type StripeObject struct {
	Paymentintent     string
	PaymentMethodType string
}

type Message struct {
	Success string      `json:"success"`
	Message interface{} `json:"Message"`
}

//Create PaymentIntent
func CreatePaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(request.Body)
	var param *stripe.PaymentIntentParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}

	pi, err := paymentintent.New(param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}

	bytes, err := json.Marshal(pi)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusCreated)
}

//Retrieve PaymentIntent
func RetrievePaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(request)
	var id = vars["paymentintentid"]

	pi, err := paymentintent.Get(id, nil)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	bytes, err := json.Marshal(pi)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Update PaymentIntent
func UpdatePaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(request)
	var id = vars["paymentintentid"]

	decoderpi := json.NewDecoder(request.Body)
	var param *stripe.PaymentIntentParams
	err := decoderpi.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	o, err := paymentintent.Update(id, param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	bytes, err := json.Marshal(o)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Capture PaymentIntent
func CapturePaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(request)
	var id = vars["paymentintentid"]

	decoderpi := json.NewDecoder(request.Body)
	var param *stripe.PaymentIntentCaptureParams
	err := decoderpi.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	intent, err := paymentintent.Capture(id, param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}

	bytes, err := json.Marshal(intent)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Cancel PaymentIntent
func CancelPaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(request)
	var id = vars["paymentintentid"]

	intent, err := paymentintent.Cancel(id, nil)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	bytes, err := json.Marshal(intent)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//List All PaymentIntent
func ListAllPaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	params := stripe.PaymentIntentListParams{}
	res := paymentintent.List(&params)

	var listpaymentintent []stripe.PaymentIntent

	for res.Next() {
		paymentintent := stripe.PaymentIntent{}
		listpaymentintent = append(listpaymentintent, paymentintent)
	}
	fmt.Println("res :::", listpaymentintent)
	bytes, err := json.Marshal(&listpaymentintent)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
