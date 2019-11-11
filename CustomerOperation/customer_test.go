package CustomerOperation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/oms-services/stripe/result"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Customer struct {
	Email       string `json:"email"`
	Description string `json:"description"`
}

var _ = Describe("Create Customer operations", func() {

	customer := Customer{Email: "testcust@demo.com", Description: "Test Demo Customer"}
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(customer)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")

	req, err := http.NewRequest("POST", "/createcustomer", reqbody)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCustomer)
	handler.ServeHTTP(recorder, req)

	Describe("create customer", func() {
		Context("CreateCustomer", func() {
			It("Should result http.StatusCreated", func() {
				Expect(recorder.Code).To(Equal(http.StatusCreated))
			})
		})
	})
})

var _ = Describe("Create Customer operations invalid data", func() {
	//invalid emailID
	customer := Customer{Email: "testcustdemocom", Description: "Test Demo Customer"}
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(customer)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")

	req, err := http.NewRequest("POST", "/createcustomer", reqbody)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCustomer)
	handler.ServeHTTP(recorder, req)

	It("Should result http.StatusBadRequest", func() {
		Expect(recorder.Code).To(Equal(http.StatusBadRequest))
	})
})
