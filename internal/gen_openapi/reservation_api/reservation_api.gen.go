// Package reservation_api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package reservation_api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// CreateReservationRequest defines model for CreateReservationRequest.
type CreateReservationRequest struct {
	End   *string `json:"end,omitempty"`
	Name  *string `json:"name,omitempty"`
	Start *string `json:"start,omitempty"`
}

// ListReservationRequest defines model for ListReservationRequest.
type ListReservationRequest struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

// ListReservationResponse defines model for ListReservationResponse.
type ListReservationResponse struct {
	Reservations *[]Reservation `json:"reservations,omitempty"`
}

// Placement defines model for Placement.
type Placement struct {
	Height *int    `json:"height,omitempty"`
	Seats  *[]Seat `json:"seats,omitempty"`
	Width  *int    `json:"width,omitempty"`
}

// Reservation defines model for Reservation.
type Reservation struct {
	End       *string    `json:"end,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Placement *Placement `json:"placement,omitempty"`
	Start     *string    `json:"start,omitempty"`
}

// Seat defines model for Seat.
type Seat struct {
	Floor    *int     `json:"floor,omitempty"`
	Height   *int     `json:"height,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Rotation *float32 `json:"rotation,omitempty"`
	Space    *int     `json:"space,omitempty"`
	Status   *string  `json:"status,omitempty"`
	True     *float32 `json:"true,omitempty"`
	User     *string  `json:"user,omitempty"`
	Width    *int     `json:"width,omitempty"`
	X        *float32 `json:"x,omitempty"`
}

// GetReservationBusinessIdJSONBody defines parameters for GetReservationBusinessId.
type GetReservationBusinessIdJSONBody ListReservationRequest

// PostReservationBusinessIdJSONBody defines parameters for PostReservationBusinessId.
type PostReservationBusinessIdJSONBody CreateReservationRequest

// GetReservationBusinessIdRequestBody defines body for GetReservationBusinessId for application/json ContentType.
type GetReservationBusinessIdJSONRequestBody GetReservationBusinessIdJSONBody

// PostReservationBusinessIdRequestBody defines body for PostReservationBusinessId for application/json ContentType.
type PostReservationBusinessIdJSONRequestBody PostReservationBusinessIdJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /reservation/{businessId})
	GetReservationBusinessId(ctx echo.Context, businessId string) error

	// (POST /reservation/{businessId})
	PostReservationBusinessId(ctx echo.Context, businessId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetReservationBusinessId converts echo context to params.
func (w *ServerInterfaceWrapper) GetReservationBusinessId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "businessId" -------------
	var businessId string

	err = runtime.BindStyledParameter("simple", false, "businessId", ctx.Param("businessId"), &businessId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter businessId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReservationBusinessId(ctx, businessId)
	return err
}

// PostReservationBusinessId converts echo context to params.
func (w *ServerInterfaceWrapper) PostReservationBusinessId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "businessId" -------------
	var businessId string

	err = runtime.BindStyledParameter("simple", false, "businessId", ctx.Param("businessId"), &businessId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter businessId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostReservationBusinessId(ctx, businessId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/reservation/:businessId", wrapper.GetReservationBusinessId)
	router.POST(baseURL+"/reservation/:businessId", wrapper.PostReservationBusinessId)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RVQW/bPAz9KwK/72jUaXeaj92GoUAPRXccemBsOlYhSypFpwuC/PdBctI4tVJ0KHba",
	"KbFEPT4+PlFbqF3vnSUrAaothLqjHtPfL0wodE+BeI2inb2np4GCxD3PzhOLphRJtok/svEEFQRhbVew",
	"K8BiT9mNIMiS2dkVhxW3fKRaYuytDvIREh/KFbyzgebJ+BiUvrVQn/78z9RCBf+VR1nLvablBBmO2ZEZ",
	"N3k6dwZr6slmqu1Ir7ppWdoKrYhTxYTyflY/CGVOp4Bn3UiXS5BjOi3t4/bw07rf4n4U6A8bnYqeEW2N",
	"c5zX9C29z9bBTl4k2W/aoV/uu+SxpjMNFJQhZCGFB8rCDYE4e+JsHwv4lUGaqxWXtG1dDG4o1Kz9WBR8",
	"s4132opqHasoqaFaFDa9tqpHi6vUG+UNSuu4hwJEi4nYU78UsCYOI+LlxSISc54seg0VfLpYXFxBAR6l",
	"S4qUk6tXbpdD0JZCuGl2cXNFMmcZL7VCdQhVJ3c35eL0cdNABd9pOgCuX+ATBcaehDhA9fN1kpuvyrVK",
	"OlKHMxBVgyoxh4NJYDlFZHoaNFMDVWxrsR++OQs/jMEU5No1mxhROyv7G4LeG12PijyG0W5HqLfuz5nZ",
	"mlp+Si4tjNMwteFqsfh7LPZTN9E4lfmehDWtqVEmdtW1J91UYahrCqEdjNmMx70LGUuMT5tCZel5ijCz",
	"w50L/5Ifzj7573LE5VzoCZQawZuItUuvFK8P2g1soIJOxFdlaVyNpnNBqq13LLsSvS7Xl/DaC7cxTo0w",
	"cYgga1yakUw8OLJpcTACFXxeLBYx9cPudwAAAP//8d2jO/MIAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
