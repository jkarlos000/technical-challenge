// Package doc API Falabella FIF
//
// Esta API esta diseñada para ser una prueba para los nuevos candidatos al equipo.
//
//     Schemes: http
//     BasePath: /v1
//     Version: 1.0.0
//     Host: localhost:8080
//     Contact: ugaetea@falabella.cl
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package doc

import (
	"github.com/jkarlos000/technical-challenge/beer-api/internal/beer"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/errors"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/pagination"
)

// El Id de la cerveza no existe
//
// swagger:response notFound
type notFoundWrapper struct {

	// A Error entity
	// in: body
	Body errors.ErrorResponse
}

// El ID de la cerveza ya existe
//
// swagger:response conflictError
type conflictdWrapper struct {

	// A Error entity
	// in: body
	Body errors.ErrorResponse
}

// Ingresa una nueva cerveza
//
// swagger:parameters AddBeers
type beerCreateRequestWrapper struct {
	// Parameters to create an Admin
	// in: body
	Body struct{
		// swagger:allOf
		beer.CreateBeerRequest
	}
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
//
// swagger:response internalServerError
type internalServerErrorWrapper struct {
	// A Error entity
	// in: body
	Body errors.ErrorResponse
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
//
// swagger:response unauthorized
type unauthorizedWrapper struct {
	// A Error entity
	// in: body
	Body errors.ErrorResponse
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
//
// swagger:response forbidden
type forbiddenWrapper struct {

	// A Error entity
	// in: body
	Body errors.ErrorResponse
}


// Request invalida
//
// swagger:response badRequest
type badRequestWrapper struct {

	// A Error entity
	// in: body
	Body errors.ErrorResponse
}

// InvalidInput creates a new error response representing a data validation error (HTTP 400).
//
// swagger:response invalidInput
type invalidInputWrapper struct {

	// A Error entity
	// in: body
	Body errors.ErrorResponse
}

// swagger:parameters SearchBeerById BoxBeerPriceById
type beerIDParameterWrapper struct {
	// Busca una cerveza por su Id
	// in: path
	// required: true
	ID string `json:"beerID"`
}

// Operacion exitosa
//
// swagger:response beersItemResponse
type beerItems struct {
	// The error message
	// in: body
	Body struct {
		// swagger:allOf
		pagination.Pages
		// swagger:allOf
		Items []beer.BeerItem `json:"items"`
	}
}

// Operacion exitosa
//
// swagger:response beerItemResponse
type beerItem struct {
	// The error message
	// in: body
	Body struct {
		// swagger:allOf
		beer.BeerItem
	}
}

// Cerveza creada
//
// swagger:response beerItemCreatedResponse
type beerItemCreated struct {
	// The error message
	// in: body
	Body struct {
		// swagger:allOf
		beer.BeerItem
	}
}

// Operacion exitosa
//
// swagger:response beerBoxResponse
type beerBoxItem struct {
	// The error message
	// in: body
	Body struct {
		// swagger:allOf
		beer.BeerBox
	}
}

// swagger:parameters BoxBeerPriceById
type beerQueryParamsWrapper struct {
	// Query-Params in URL
	// in: query

	//Tipo de moneda con la que pagará
	//
	Currency string `json:"currency"`

	//La cantidad de cervezas a comprar
	//
	// default: 6
	Quantity int `json:"quantity"`
}
