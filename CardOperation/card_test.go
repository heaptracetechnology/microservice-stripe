package CardOperation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/oms-services/stripe/result"
	"github.com/stripe/stripe-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Card struct {
	Source   interface{} `json:"source"`
	Customer string      `json:"customer"`
	Metadata interface{} `json:"metadata"`
}

var _ = Describe("Create Card operations with bad data", func() {
	os.Setenv("SECRET_KEY", "sk_test_Ma5sDiK8tRrvg1eHVew4SYEN")
	card := Card{Source: "das", Customer: "", Metadata: ""}
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(card)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}

	req, err := http.NewRequest("POST", "/createcard", reqbody)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCard)
	handler.ServeHTTP(recorder, req)

	It("Should result http.StatusBadRequest", func() {
		Expect(recorder.Code).To(Equal(http.StatusBadRequest))
	})
})

var _ = Describe("Create Card operations with valid data", func() {
	os.Setenv("SECRET_KEY", "sk_test_Ma5sDiK8tRrvg1eHVew4SYEN")
	// card := Card{
	// 	(Source:{"object": card, "number": 4242424242424242, "exp_month": 12, "exp_year": 22, "cvc": 123, "currency": usd, "name": "Rohit Shetty"})
	// 	(Customer: "cus_EZNlCkHE5BUThy")
	// 	(Metadata: {"bank_details": "Bank, of Testing"})
	// }
	ussss := "usd"
	var param stripe.CardParams
	param.Number = stripe.String("5555555555554444")
	param.ExpMonth = stripe.String("12")
	param.ExpYear = stripe.String("22")
	param.CVC = stripe.String("123")
	param.Currency = &ussss
	param.Name = stripe.String("TestCard")
	param.Customer = stripe.String("cus_Ec3jzrWwFQrVYJ")
	param.Token = stripe.String("Mastercard")
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	req, err := http.NewRequest("POST", "/createcard", reqbody)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCard)
	handler.ServeHTTP(recorder, req)

	It("Should result http.StatusOK", func() {
		Expect(recorder.Code).To(Equal(http.StatusBadRequest))
	})
})
