FROM golang

RUN go get github.com/stripe/stripe-go

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/oms-services/stripe

ADD . /go/src/github.com/oms-services/stripe

RUN go install github.com/oms-services/stripe

ENTRYPOINT stripe

EXPOSE 3000