package SourceOperation

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oms-services/stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/source"
	"net/http"
	"os"
)

//Create Source
func CreateSource(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var tranftype = vars["transfertype"]

	decoder := json.NewDecoder(r.Body)
	var param *stripe.SourceObjectParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	param.Type = stripe.String(tranftype)
	s, errr := source.New(param)
	if errr != nil {
		result.WriteErrorResponse(w, err)
	}
	fmt.Println(s)
	bytes, _ := json.Marshal(s)
	result.WriteJsonResponse(w, bytes, http.StatusOK)

}
