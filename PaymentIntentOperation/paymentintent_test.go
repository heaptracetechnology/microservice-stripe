package PaymentIntentOperation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gorilla/mux"
	"github.com/oms-services/stripe/result"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type PaymentIntent struct {
	Amount             int64    `json:"amount"`
	Currency           string   `json:"currency"`
	Paymentmethodtypes []string `json:"paymentmethodtypes"`
}

var _ = Describe("Create PaymentIntent operations", func() {

	pmt := []string{"card"}
	paymentintent := PaymentIntent{Amount: 1000, Currency: "usd", Paymentmethodtypes: pmt}
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(paymentintent)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("POST", "/createpaymentintent", reqbody)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePaymentIntent)
	handler.ServeHTTP(recorder, req)
	Describe("create payment intent", func() {
		Context("CreatePaymentIntent", func() {
			It("Should result http.StatusCreated", func() {
				Expect(recorder.Code).To(Equal(http.StatusCreated))
			})
		})
	})
})

var _ = Describe("Retrive PaymentIntent operations", func() {

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("GET", "/retrievepaymentIntent/paymentintentid", nil)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	vars := map[string]string{
		"paymentintentid": "pi_1EAHHeJytX7n0OoXQnPHro2u",
	}

	req = mux.SetURLVars(req, vars)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RetrievePaymentIntent)
	handler.ServeHTTP(recorder, req)
	Describe("retrive payment intent", func() {
		Context("RetrivePaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Update PaymentIntent operations", func() {

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	paymentintent := PaymentIntent{Amount: 2000, Currency: "usd"}
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(paymentintent)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	req, err := http.NewRequest("PUT", "/updatepaymentintent/paymentintentid", reqbody)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	vars := map[string]string{
		"paymentintentid": "pi_1EAHHeJytX7n0OoXQnPHro2u",
	}
	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdatePaymentIntent)
	handler.ServeHTTP(recorder, req)
	Describe("update payment intent", func() {
		Context("UpdatePaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Capture PaymentIntent operations", func() {

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	paymentintent := PaymentIntent{}
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(paymentintent)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	req, err := http.NewRequest("PUT", "/capturecharge/paymentintentid", reqbody)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	vars := map[string]string{
		"paymentintentid": "pi_1EAHHeJytX7n0OoXQnPHro2u",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CapturePaymentIntent)
	handler.ServeHTTP(recorder, req)
	Describe("capture payment intent", func() {
		Context("CapturePaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Cancel PaymentIntent operations", func() {

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	paymentintent := PaymentIntent{}
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(paymentintent)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	req, err := http.NewRequest("POST", "/cancelpaymentintent/paymentintentid", nil)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	vars := map[string]string{
		"paymentintentid": "pi_1EAHHeJytX7n0OoXQnPHro2u",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CancelPaymentIntent)
	handler.ServeHTTP(recorder, req)
	Describe("cancel payment intent", func() {
		Context("CancelPaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("List All PaymentIntent operations", func() {

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("GET", "/listallpaymentintent", nil)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListAllPaymentIntent)
	handler.ServeHTTP(recorder, req)
	Describe("list all payment intent", func() {
		Context("ListAllPaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
