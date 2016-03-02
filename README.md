# go-rabbitmq
[![Build Status](https://travis-ci.org/cilindrox/go-rabbitmq.svg?branch=master)](https://travis-ci.org/cilindrox/go-rabbitmq) [![Code Climate](https://codeclimate.com/repos/56cb2250f5d6231ca700dc2f/badges/04d036cd3c37a0707b2a/gpa.svg)](https://codeclimate.com/repos/56cb2250f5d6231ca700dc2f/feed)

This is basically a sample implementation, mostly taken from the [RabbitMQ tutorials](http://www.rabbitmq.com/getstarted.html).


## Requirements

To run this code you need [Go RabbitMQ client](https://github.com/streadway/amqp):

    go get github.com/streadway/amqp

`RABBIT_URL="amqp://guest:guest@localhost:5672"`


## Code

Code examples are executed via `go run`:

[Tutorial one: "Hello World!"](http://www.rabbitmq.com/tutorial-one-go.html):

    go run send.go
    go run receive.go


To learn more, see [Go RabbitMQ client](https://github.com/streadway/amqp).
