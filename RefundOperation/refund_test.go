package RefundOperation

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

type Refund struct {
	ChargeId string `json:"charge"`
}

var _ = Describe("Create Refund operations", func() {

	refund := Refund{ChargeId: "ch_1EAIJQJytX7n0OoXhSP6RTc9"}
	reqbody := new(bytes.Buffer)
	err := json.NewEncoder(reqbody).Encode(refund)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("POST", "/createrefund", reqbody)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRefund)
	handler.ServeHTTP(recorder, req)

	Describe("create refund", func() {
		Context("CreateRefund", func() {
			It("Should result http.StatusBadRequest already refunded", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
