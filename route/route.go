package route

import (
    "github.com/gorilla/mux"
    "github.com/user/stripe/StripeOperation"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "CreateCharge",
        "POST",
        "/createcharge",
        StripeOperation.CreateCharge,
    },
    Route{
        "GetBalance",
        "GET",
        "/getbalance",
        StripeOperation.GetBalance,
    },
    Route{
        "CreatePaymentIntent",
        "POST",
        "/createpaymentintent",
        StripeOperation.CreatePaymentIntent,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
