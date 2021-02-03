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
	Id           *string `json:"_id,omitempty"`
	Address      *string `json:"address,omitempty"`
	BusinessName *string `json:"businessName,omitempty"`
	Description  *string `json:"description,omitempty"`
	Location     *struct {
		Latitude  *float32 `json:"latitude,omitempty"`
		Longitude *float32 `json:"longitude,omitempty"`
	} `json:"location,omitempty"`
	Policy *struct {
		MinAge *int `json:"minAge,omitempty"`
	} `json:"policy,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Type *string   `json:"type,omitempty"`
}

// ListRequestResponse defines model for ListRequestResponse.
type ListRequestResponse struct {
	MaxPage *int             `json:"maxPage,omitempty"`
	Request *[]ChangeRequest `json:"request,omitempty"`
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
	router.GET(baseURL+"/request/:businessId", wrapper.GetRequestBusinessId)
	router.POST(baseURL+"/request/:businessId", wrapper.PostRequestBusinessId)
	router.POST(baseURL+"/request/:businessId/approve", wrapper.PostRequestBusinessIdApprove)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8yW0WvbPhDH/xVxv9+jid1uL/Nb10EplFHyOspQ7LOjYkvq6ZwthPzvQ5KdxLHTjkGh",
	"T3Yl3d1X9/memx0UprVGo2YH+Q5cscZWhtfbtdQ1LvGlQ8d+wZKxSKwwbP9UpX/w1iLk4JiUrmGfgCxL",
	"Qudm91adUxqd+y5bnD1QoitIWVZGz+43ppDD5lhOI1lxV56m1V27Qophur60u0+GFbN6xoL9eWsaVWyn",
	"RVqlb+rTJEoz1peysKxDlGJs5/vRL0giuT3+PTk4l/xBOe7ZLNFZox3O6JW/H+W84AToSPag8H/CCnL4",
	"Lz26Iu0tkY79MFE/FemXlK5MKK+48XtDeAIbJBdIwtUiW2Q+obGopVWQw6dFtriGBKzkdRCWnqitMTxG",
	"Xgn9EIO/hNSlUOxEETSLPlh4NdRGB4VyFN7vS8jhDvkozkqSLTKSg/zHDpSv8NIhbSEBHcwL1jc2dlER",
	"lpAzdZj0AzRrkSd/OqIKd7rOMv8ojGbU4UrS2kZFh6fPLtr8mPA1OHN2CADGXVoik8INisa3y1Tjjp11",
	"y3VFgc5VXdN4vj7bgCHdDYH35f4ikztkIYWzWKhKFWfpXwHw9ZB7imJc4f6bvwSvUQwxkERa3jlHWKvT",
	"jG8iO8zdexI7G6cpq9uL3hXUYywnjPy3y82guCWUjEIKjb/OOVeGhDw4YcLl0biPCSZKMuX2PZmMNe0n",
	"hri62Ot/nKZUWktmE7/msyxv4oETZm+N1izCPs2HG7HP0xsPlf9yCHxvHdJmuFBHDeSwZrZ5mvqfEM3a",
	"OM531hD7fqt0cwXn4/fgz4mYxv+7kqTkqokifWBUWcmuYcjhS5ZlvvTT/k8AAAD//1bdhLJSCQAA",
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
