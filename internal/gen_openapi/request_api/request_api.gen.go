// Package request_api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package request_api

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

// ChangeRequest defines model for ChangeRequest.
type ChangeRequest struct {
	Id           *string   `json:"_id,omitempty"`
	Address      *string   `json:"address,omitempty"`
	BusinessName *string   `json:"businessName,omitempty"`
	Description  *string   `json:"description,omitempty"`
	Location     *Location `json:"location,omitempty"`
	Tags         *[]string `json:"tags,omitempty"`
	Type         *string   `json:"type,omitempty"`
}

// ListRequestResponse defines model for ListRequestResponse.
type ListRequestResponse struct {
	Request *[]ChangeRequest `json:"request,omitempty"`
}

// Location defines model for Location.
type Location struct {
	Latitude  *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
}

// GetRequestParams defines parameters for GetRequest.
type GetRequestParams struct {
	Page int `json:"page"`
}

// PostRequestBusinessIdJSONBody defines parameters for PostRequestBusinessId.
type PostRequestBusinessIdJSONBody ChangeRequest

// PostRequestBusinessIdRequestBody defines body for PostRequestBusinessId for application/json ContentType.
type PostRequestBusinessIdJSONRequestBody PostRequestBusinessIdJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /request)
	GetRequest(ctx echo.Context, params GetRequestParams) error

	// (DELETE /request/{businessId})
	DeleteRequestBusinessId(ctx echo.Context, businessId string) error

	// (GET /request/{businessId})
	GetRequestBusinessId(ctx echo.Context, businessId string) error

	// (POST /request/{businessId})
	PostRequestBusinessId(ctx echo.Context, businessId string) error

	// (POST /request/{businessId}/approve)
	PostRequestBusinessIdApprove(ctx echo.Context, businessId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetRequest converts echo context to params.
func (w *ServerInterfaceWrapper) GetRequest(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRequestParams
	// ------------- Required query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, true, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRequest(ctx, params)
	return err
}

// DeleteRequestBusinessId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteRequestBusinessId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "businessId" -------------
	var businessId string

	err = runtime.BindStyledParameter("simple", false, "businessId", ctx.Param("businessId"), &businessId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter businessId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteRequestBusinessId(ctx, businessId)
	return err
}

// GetRequestBusinessId converts echo context to params.
func (w *ServerInterfaceWrapper) GetRequestBusinessId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "businessId" -------------
	var businessId string

	err = runtime.BindStyledParameter("simple", false, "businessId", ctx.Param("businessId"), &businessId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter businessId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRequestBusinessId(ctx, businessId)
	return err
}

// PostRequestBusinessId converts echo context to params.
func (w *ServerInterfaceWrapper) PostRequestBusinessId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "businessId" -------------
	var businessId string

	err = runtime.BindStyledParameter("simple", false, "businessId", ctx.Param("businessId"), &businessId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter businessId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRequestBusinessId(ctx, businessId)
	return err
}

// PostRequestBusinessIdApprove converts echo context to params.
func (w *ServerInterfaceWrapper) PostRequestBusinessIdApprove(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "businessId" -------------
	var businessId string

	err = runtime.BindStyledParameter("simple", false, "businessId", ctx.Param("businessId"), &businessId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter businessId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRequestBusinessIdApprove(ctx, businessId)
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

	router.GET(baseURL+"/request", wrapper.GetRequest)
	router.DELETE(baseURL+"/request/:businessId", wrapper.DeleteRequestBusinessId)
	router.GET(baseURL+"/request/:businessId", wrapper.GetRequestBusinessId)
	router.POST(baseURL+"/request/:businessId", wrapper.PostRequestBusinessId)
	router.POST(baseURL+"/request/:businessId/approve", wrapper.PostRequestBusinessIdApprove)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RWQWvbTBD9K8t831FYStpLdWscCIFQiq8llLU0ljdIu5vZkYsx/u9lV5JleZWElhbS",
	"k2zN7Myb995IOkBhGms0anaQH8AVW2xk+LncSl3hCp9bdOxvWDIWiRWG8HdV+gvvLUIOjknpCo4JyLIk",
	"dG42tm6d0ujcF9ngbEKJriBlWRk9G69NIYfg/4QbyOG/dJwg7eGnD0PeMQGWVUCjGJt5WP0NSST34/8o",
	"ccw06ycs2Kc+KMc9RSt01miHMVU0cngC8Rr6KfMRwFkcZ8RMm9eSFbfl+UC6bdZIHZ+6eikat/G3lN6Y",
	"kKy49rEBZQI7JBcQwNUiW2S+vLGopVWQw4dFtriGBKzkbYCVnpFSYbhMxA/MisEwQupSKHaiCNSI/rDw",
	"aKjpJg/tKPy+LyGHO+QRnJUkG2QkB/m3Ayjf4blF2kMCOrgRrKwQkiCWIiwhZ2ox6TfijB+lGatA0KPP",
	"7kQPM11nmb8URjPqMJK0tladMumT6+QZC77q4BljBQGmLK2QSeEORe3pMpspYxdsubYo0LlNW9feRr7a",
	"IEN6GA7el8dOjBoZY1lW6L0gpHAWC7VRxUWTSIbbUKgf5ebUJNZk2uf+1k/DWxTDGUg62byFRtXW5xXf",
	"1O60yrF0H+NZl1P6KIyOZcRjMu/gO/wVnka7vmOS/py/L55xsbOXL266oN70c0pY42akWBJKRiGFxh+X",
	"W7ExJORpbyJdvhr3PoXpIJly/zc1mWI6Roa4epHr33z2pNJaMrvuLTqr5ecu4Uyzt1ZrVsK+zL/wHLqZ",
	"n1P0VJWBz2MCDmk3TNFSDTlsmW2epv6zqd4ax/nBGmJPskp3V3C5c/4rohZdGf9Gl6Tkuu6Q+YMdtI1s",
	"a4YcPmVZ5ls/Hn8GAAD//+f47SFGCgAA",
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
