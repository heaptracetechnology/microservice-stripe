package TokenOperation

import (
	"encoding/json"
	"github.com/oms-services/stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/token"
	"net/http"
	"os"
)

//CreateToken
func CreateToken(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.TokenParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	token, err := token.New(param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}
	bytes, _ := json.Marshal(token)
	result.WriteJsonResponse(w, bytes, http.StatusCreated)
}
