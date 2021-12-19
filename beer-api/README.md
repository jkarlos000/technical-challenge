# Microservices-test

[![GoDoc](https://godoc.org/github.com/qiangxue/go-rest-api?status.png)](https://bitbucket.org/lgaetecl/microservices-test/src/master/)
[![Build Status](https://github.com/qiangxue/go-rest-api/workflows/build/badge.svg)](https://github.com/jkarlos000/technical-challenge/actions?query=workflow%3Abuild)
[![Code Coverage](https://codecov.io/gh/qiangxue/go-rest-api/branch/master/graph/badge.svg)](https://app.codecov.io/gh/jkarlos000/technical-challenge)
[![Go Report](https://goreportcard.com/badge/github.com/qiangxue/go-rest-api)](https://goreportcard.com/report/github.com/qiangxue/go-rest-api)

Go based Beer API built using the Ozzo Routing.

## Documentation

OpenAPI documentation can be found in the [swagger.yaml](./swaggerui/swagger.yaml) file

## Running

The applicaiton can be run with `go run`

```
➜ make run
go run -ldflags "-X main.Version=0f7570a" cmd/server/main.go
{"level":"info","ts":1639930252.2630742,"caller":"server/main.go:93","msg":"server 0f7570a is running at :8080","version":"0f7570a"}

curl localhost:8080/beers
```

The RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `GET /healthcheck`: a healthcheck service provided for health checking purpose (needed when implementing a server cluster)
* `GET /v1/beers`: returns a paginated list of the beers
* `GET /v1/beers/:id`: returns the detailed information of a beer
* `POST /v1/beers`: creates a new beer
* `GET /v1/albums/:id/beerbox`: returns the price of beer box

Try the URL `http://localhost:8080/healthcheck` in a browser, and you should see something like `"OK v1.0.0"` displayed.