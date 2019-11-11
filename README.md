# Stripe as a microservice
An OMG service for Stripe, it allows individuals and businesses to receive payments over the Internet. Stripe provides the technical, fraud prevention, and banking infrastructure required to operate on-line payment systems.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.com/oms-services/stripe.svg?branch=master)](https://travis-ci.com/oms-services/stripe)
[![codecov](https://codecov.io/gh/oms-services/stripe/branch/master/graph/badge.svg)](https://codecov.io/gh/oms-services/stripe)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Create Charge
```sh
$ omg run create_charge -a amount=<AMOUNT> -a currency=<CURRENCY> -a description=<DESCRIPTION> -a capture=<TRUE/FALSE>  -e SECRET_KEY=<SECRET_KEY>
```
##### Capture Charge
```sh
$ omg run capture_charge -a charge=<CHARGE_ID> -e SECRET_KEY=<SECRET_KEY>
```
##### Get Balance
```sh
$ omg run get_balance -e SECRET_KEY=<SECRET_KEY>
```
##### Create Paymentintent
```sh
$ omg run create_payment_intent -a amount=<AMOUNT> -a currency=<CURRENCY> -a customer=<CUSTOMER_ID> -a savepaymentmethod=<TRUE/FALSE> -a capturemethod=<AUTOMATIC/MANUAL> -a paymentmethodtypes=<METHOD_LIST> -e SECRET_KEY=<SECRET_KEY>
```
##### Retrive Paymentintent
```sh
$ omg run retrieve_payment_intent -a paymentintentid=<PAYMENTINTENT_ID> -e SECRET_KEY=<SECRET_KEY>
```
##### Update Paymentintent
```sh
$ omg run update_payment_intent -a amount=<AMOUNT> -a paymentintentid=<PAYMENTINTENT_ID> -a shipping=<SHIPPING_OBJECT> -e SECRET_KEY=<SECRET_KEY>
```
##### Capture Paymentintent
```sh
$ omg run capture_payment_intent -a paymentintentid=<PAYMENTINTENT_ID> -a amounttocapture=<AMOUNT_TO_CAPTURE> -e SECRET_KEY=sk_test_gENQu8ecxwwMUsWlgsQeqbgI
```
##### Cancel PaymentIntent
```sh
$ omg run cancel_payment_intent -a paymentintentid=<PAYMENTINTENT_ID> -e SECRET_KEY=<SECRET_KEY>
```
##### List PaymentIntent
```sh
omg run list_all_payment_intent -e SECRET_KEY=<SECRET_KEY>
```
##### Create Customer
```sh
$ omg run create_customer -a description=<DESCRIPTION> -a email=<CUSTOMER_EMAIL> -e SECRET_KEY=<SECRET_KEY>
```
##### Create Card
```sh
$ omg run create_card -a source=<SOURCE_DETAILS> -a customer=<CUSTOMER_ID> -a metadata=<METADATA> -a token=<CARD_TOKEN>  -e SECRET_KEY=<SECRET_KEY>
```
##### Create Refund
```sh
$ omg run create_refund -a charge=<CHARGE_ID> -e SECRET_KEY=<SECRET_KEY>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-stripe .
```
### RUN
```
docker run -p 3000:3000 microservice-stripe
```
