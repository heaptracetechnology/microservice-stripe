package BalanceOperation

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/oms-services/stripe/result"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Balance operations with valid data", func() {

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("GET", "/getbalance", nil)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBalance)
	handler.ServeHTTP(recorder, req)

	It("Should result http.StatusOK", func() {
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})
})

var _ = Describe("Get Balance operations with invalid data", func() {

	//Invalid SECRET_KEY
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI11")
	req, err := http.NewRequest("GET", "/getbalance", nil)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBalance)
	handler.ServeHTTP(recorder, req)

	It("Should result http.StatusBadRequest", func() {
		Expect(recorder.Code).To(Equal(http.StatusBadRequest))
	})
})
