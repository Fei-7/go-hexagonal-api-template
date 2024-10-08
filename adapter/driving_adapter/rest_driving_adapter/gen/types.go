// Package api_gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api_gen

// Book defines model for Book.
type Book struct {
	Id    *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
}

// GetBooksResponse defines model for GetBooksResponse.
type GetBooksResponse struct {
	Data []Book `json:"data"`
}

// Pong defines model for Pong.
type Pong struct {
	Ping string `json:"ping"`
}
