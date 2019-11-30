# _Stripe_ Open Microservice

> This is a Stripe service

[![Open Microservice Specification Version](https://img.shields.io/badge/Open%20Microservice-1.0-477bf3.svg)](https://openmicroservices.org) [![Open Microservices Spectrum Chat](https://withspectrum.github.io/badge/badge.svg)](https://spectrum.chat/open-microservices) [![Open Microservices Code of Conduct](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](https://github.com/oms-services/.github/blob/master/CODE_OF_CONDUCT.md) [![Open Microservices Commitzen](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com) 
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Introduction

This project is an example implementation of the [Open Microservice Specification](https://openmicroservices.org), a standard originally created at [Storyscript](https://storyscript.io) for building highly-portable "microservices" that expose the events, actions, and APIs inside containerized software.

## Getting Started

The `oms` command-line interface allows you to interact with Open Microservices. If you're interested in creating an Open Microservice the CLI also helps validate, test, and debug your `oms.yml` implementation!

See the [oms-cli](https://github.com/microservices/oms) project to learn more!

### Installation

```
npm install -g @microservices/oms
```

## Usage

### Open Microservices CLI Usage

Once you have the [oms-cli](https://github.com/microservices/oms) installed, you can run any of the following commands from within this project's root directory:

#### Actions

##### create_charge

> Create a New Charge
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| amount | `int` | `true` | None | The amount of charge to be created |
| currency | `string` | `true` | None | The type of currency |
| description | `string` | `false` | None | The description for charge |
| applicationfeeamount | `int` | `false` | None | A fee in cents that will be applied to the charge and transferred to the application owner’s Stripe account |
| capture | `boolean` | `false` | None | Whether to immediately capture the charge |
| customer | `string` | `false` | None | The ID of an existing customer that will be charged in this request. |
| metadata | `object` | `false` | None | Set of key-value pairs that you can attach to an object |
| onbehalfof | `string` | `false` | None | The Stripe account ID for which these funds are intended. Automatically set if you use the destination parameter |
| receiptemail | `string` | `false` | None | The email address to which this charge’s receipt will be sent |
| shipping | `object` | `false` | None | Shipping information for the charge. Helps prevent fraud on charges for physical goods. |
| source | `object` | `false` | None | A payment source to be charged |
| transferdata | `string` | `false` | None | An optional dictionary including the account to automatically transfer to as part of a destination charge (required destination feild) |
| transfergroup | `string` | `false` | None | A string that identifies this transaction as part of a group |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run create_charge \ 
    -a amount='*****' \ 
    -a currency='*****' \ 
    -a description='*****' \ 
    -a applicationfeeamount='*****' \ 
    -a capture='*****' \ 
    -a customer='*****' \ 
    -a metadata='*****' \ 
    -a onbehalfof='*****' \ 
    -a receiptemail='*****' \ 
    -a shipping='*****' \ 
    -a source='*****' \ 
    -a transferdata='*****' \ 
    -a transfergroup='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### capture_charge

> Capture Charge

##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| charge | `string` | `true` | None | The ID of charge to capture |
| amount | `int` | `false` | None | The amount of charge to be created |
| applicationfeeamount | `int` | `false` | None | A fee in cents that will be applied to the charge and transferred to the application owner’s Stripe account |
| receiptemail | `string` | `false` | None | The email address to which this charge’s receipt will be sent |
| transferdata | `string` | `false` | None | An optional dictionary including the account to automatically transfer to as part of a destination charge (required destination feild) |
| transfergroup | `string` | `false` | None | A string that identifies this transaction as part of a group |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run capture_charge \ 
    -a charge='*****' \ 
    -a amount='*****' \ 
    -a applicationfeeamount='*****' \ 
    -a receiptemail='*****' \ 
    -a transferdata='*****' \ 
    -a transfergroup='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### get_balance

> Get your stripe balance
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run get_balance \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### create_payment_intent

> Create a new payment intent
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| amount | `int` | `true` | None | The amount of payment to be created |
| currency | `string` | `true` | None | The type of currency |
| paymentmethodtypes | `list` | `false` | None | The type of payment method |
| description | `string` | `false` | None | The description for paymentintent |
| applicationfeeamount | `int` | `false` | None | A fee in cents that will be applied to the paymentintent and transferred to the application owner’s Stripe account |
| capturemethod | `string` | `false` | None | Capture method of this PaymentIntent, one of automatic or manual |
| confirm | `boolean` | `false` | None | Attempt to confirm this PaymentIntent immediately |
| customer | `string` | `false` | None | ID of the customer this PaymentIntent is for if one exists. |
| metadata | `object` | `false` | None | Set of key-value pairs that you can attach to an object |
| onbehalfof | `string` | `false` | None | The Stripe account ID for which these funds are intended. Automatically set if you use the destination parameter |
| receiptemail | `string` | `false` | None | The email address to which this charge’s receipt will be sent |
| returnurl | `string` | `false` | None | The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method’s app or site |
| savepaymentmethod | `boolean` | `false` | None | Set to true to save this PaymentIntent’s payment method to the associated Customer, if the payment method is not already attached |
| shipping | `object` | `false` | None | Shipping information for the PaymentIntent. Helps prevent fraud on charges for physical goods. |
| source | `string` | `false` | None | A payment source to be PaymentIntent |
| transferdata | `string` | `false` | None | The parameters used to automatically create a Transfer when the payment succeeds |
| transfergroup | `string` | `false` | None | A string that identifies this transaction as part of a group |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run create_payment_intent \ 
    -a amount='*****' \ 
    -a currency='*****' \ 
    -a paymentmethodtypes='*****' \ 
    -a description='*****' \ 
    -a applicationfeeamount='*****' \ 
    -a capturemethod='*****' \ 
    -a confirm='*****' \ 
    -a customer='*****' \ 
    -a metadata='*****' \ 
    -a onbehalfof='*****' \ 
    -a receiptemail='*****' \ 
    -a returnurl='*****' \ 
    -a savepaymentmethod='*****' \ 
    -a shipping='*****' \ 
    -a source='*****' \ 
    -a transferdata='*****' \ 
    -a transfergroup='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### retrieve_payment_intent

> Retrieves the details of a PaymentIntent that has previously been created
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| paymentintentid | `string` | `true` | None | The payment intent code |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run retrieve_payment_intent \ 
    -a paymentintentid='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### update_payment_intent

> Updates a PaymentIntent object
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| paymentintentid | `string` | `true` | None | The payment intent code |
| amount | `int` | `false` | None | The payment amount |
| applicationfeeamount | `int` | `false` | None | The application fee amount |
| currency | `string` | `false` | None | The type of currency |
| customer | `string` | `false` | None | The customer id payment |
| description | `string` | `false` | None | The payment description |
| receiptemail | `string` | `false` | None | The receipt email address |
| savepaymentmethod | `boolean` | `false` | None | The payment method to save (true/false) |
| shipping | `object` | `false` | None | Shipping information for this PaymentIntent |
| source | `string` | `false` | None | The ID of the Source object to attach to this PaymentIntent |
| transfergroup | `string` | `false` | None | A string that identifies the resulting payment as part of a group. transfer_group may only be provided if it has not been set |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run update_payment_intent \ 
    -a paymentintentid='*****' \ 
    -a amount='*****' \ 
    -a applicationfeeamount='*****' \ 
    -a currency='*****' \ 
    -a customer='*****' \ 
    -a description='*****' \ 
    -a receiptemail='*****' \ 
    -a savepaymentmethod='*****' \ 
    -a shipping='*****' \ 
    -a source='*****' \ 
    -a transfergroup='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### capture_payment_intent

> Capture a PaymentIntent object
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| paymentintentid | `string` | `true` | None | The payment intent code |
| amounttocapture | `int` | `false` | None | The payment intent code |
| applicationfeeamount | `int` | `false` | None | The payment intent code |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run capture_payment_intent \ 
    -a paymentintentid='*****' \ 
    -a amounttocapture='*****' \ 
    -a applicationfeeamount='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### cancel_payment_intent

> Cancel PaymentIntent object
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| paymentintentid | `string` | `true` | None | The payment intent code |
| cancellationreason | `string` | `false` | None | Reason for canceling this PaymentIntent. If set, possible values are duplicate, fraudulent, requested_by_customer, or failed_invoice |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run cancel_payment_intent \ 
    -a paymentintentid='*****' \ 
    -a cancellationreason='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### list_all_payment_intent

> List all PaymentIntent object
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run list_all_payment_intent \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### create_source

> Create source for stripe
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| transfertype | `string` | `true` | None | The payment intent code |
| currency | `string` | `false` | None | The payment intent code |
| owner | `string` | `false` | None | The payment intent code |
| amount | `int` | `false` | None | Amount associated with the source. This is the amount for which the source will be chargeable once ready |
| flow | `string` | `false` | None | The authentication flow of the source to create. flow is one of redirect, receiver, code_verification, none. It is generally inferred unless a type supports multiple flows |
| mandate | `object` | `false` | None | Information about a mandate possibility attached to a source object (generally for bank debits) as well as its acceptance status |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run create_source \ 
    -a transfertype='*****' \ 
    -a currency='*****' \ 
    -a owner='*****' \ 
    -a amount='*****' \ 
    -a flow='*****' \ 
    -a mandate='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### create_customer

> Create customer for stripe
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| description | `string` | `false` | None | The description for customer |
| email | `string` | `false` | None | The customer email address |
| accountbalance | `int` | `false` | None | The customer account balance |
| invoiceprefix | `string` | `false` | None | The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers |
| invoicesettings | `object` | `false` | None | Default invoice settings for this customer |
| metadata | `object` | `false` | None | A set of key-value pairs that you can attach to a customer object. It can be useful for storing additional information about the customer in a structured format |
| shipping | `object` | `false` | None | The customer’s shipping information. Appears on invoices emailed to this customer |
| taxinfo | `object` | `false` | None | The customer’s tax information. Appears on invoices emailed to this customer |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run create_customer \ 
    -a description='*****' \ 
    -a email='*****' \ 
    -a accountbalance='*****' \ 
    -a invoiceprefix='*****' \ 
    -a invoicesettings='*****' \ 
    -a metadata='*****' \ 
    -a shipping='*****' \ 
    -a taxinfo='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### create_card

> Create card for customer.
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| source | `object` | `true` | None | The object of card details, required fields are object(type of payment source. Should be card),number(card number), exp_month, exp_year |
| customer | `string` | `true` | None | The created card will associate with customer ID |
| token | `string` | `false` | None | The type of token eg- "tok_visa" |
| metadata | `object` | `false` | None | A set of key-value pairs that you can attach to a card object |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run create_card \ 
    -a source='*****' \ 
    -a customer='*****' \ 
    -a token='*****' \ 
    -a metadata='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### create_token

> Create token for card.
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| card | `object` | `false` | None | The object of card details, required fields are object(type of payment source. Should be card),number(card number), exp_month, exp_year |
| customer | `string` | `false` | None | The token for card will associate with customer ID |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run create_token \ 
    -a card='*****' \ 
    -a customer='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

##### create_refund

> Create refund for charge.
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| charge | `string` | `true` | None | The identifier of the charge to refund |
| amount | `int` | `false` | None | positive integer in cents representing how much of this charge to refund |
| metadata | `object` | `false` | None | A set of key-value pairs that you can attach to a Refund object |
| refundapplicationfee | `boolean` | `false` | None | Boolean indicating whether the application fee should be refunded when refunding this charge |
| reversetransfer | `boolean` | `false` | None | Boolean indicating whether the transfer should be reversed when refunding this charge |
| SECRET_KEY | `string` | `true` | None | secret key for stripe |

``` shell
oms run create_refund \ 
    -a charge='*****' \ 
    -a amount='*****' \ 
    -a metadata='*****' \ 
    -a refundapplicationfee='*****' \ 
    -a reversetransfer='*****' \ 
    -e SECRET_KEY=$SECRET_KEY
```

## Contributing

All suggestions in how to improve the specification and this guide are very welcome. Feel free share your thoughts in the Issue tracker, or even better, fork the repository to implement your own ideas and submit a pull request.

[![Edit stripe on CodeSandbox](https://codesandbox.io/static/img/play-codesandbox.svg)](https://codesandbox.io/s/github/oms-services/stripe)

This project is guided by [Contributor Covenant](https://github.com/oms-services/.github/blob/master/CODE_OF_CONDUCT.md). Please read out full [Contribution Guidelines](https://github.com/oms-services/.github/blob/master/CONTRIBUTING.md).

## Additional Resources

* [Install the CLI](https://github.com/microservices/oms) - The OMS CLI helps developers create, test, validate, and build microservices.
* [Example OMS Services](https://github.com/oms-services) - Examples of OMS-compliant services written in a variety of languages.
* [Example Language Implementations](https://github.com/microservices) - Find tooling & language implementations in Node, Python, Scala, Java, Clojure.
* [Storyscript Hub](https://hub.storyscript.io) - A public registry of OMS services.
* [Community Chat](https://spectrum.chat/open-microservices) - Have ideas? Questions? Join us on Spectrum.
