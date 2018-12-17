# myRetail
 
[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

myRetail is a microservice architecture application that enables the CRUD operations for basic  for retail platform.

# APIs services! 
  - GET /heartbeat,   check system status. 
  - POST /products,   creates new product.
  ```sh
 $curl  -X POST -H "Content-Type: application/json" -d '{ "current_prise" : { "value" : 300000, "currency_code" : "INR" }, "product_details" : { "name" : "BTC", "description" : "bitcoin", "type" : "currency" } }' https://localhost:3000/products
```
  - GET /products?id=10,  get product details.
```sh
$curl -i  https://localhost:3000/products?id=50
```
- PUT /products, performs full update of product details.
```sh
 $curl  -X PUT -H "Content-Type: application/json" -d '{ "product_id" : 16, "current_prise" : { "value" : 56, "currency_code" : "INR" }, "product_details" : { "name" : "BTC", "description" : "bitcoin", "type" : "currency" } }' https://localhost:3000/products
```

- DELETE /products?id=10,, performs delete operation of product.
```sh
$curl -X DELETE  https://localhost:3000/products?id=50
``` 

### Installation

myRetail requires [Go](https://golang.org/) go1.10+ to run.

Install the dependencies and also run the DB scripts at 'DB_Scripts' and start the server.

```sh
$ git clone git@github.com:chilakantip/my_retail.git
$ cd my_retail
$ go build
$ ./my_retail
```
