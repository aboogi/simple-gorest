// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /school/materials)
	GetSchoolMaterialArray(ctx echo.Context, params GetSchoolMaterialArrayParams) error

	// (POST /school/materials)
	CreateSchoolMaterial(ctx echo.Context) error

	// (GET /school/materials/{material_id})
	GetSchoolMaterial(ctx echo.Context, materialId openapi_types.UUID) error

	// (PUT /school/materials/{material_id})
	UpdateSchoolMaterial(ctx echo.Context, materialId openapi_types.UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetSchoolMaterialArray converts echo context to params.
func (w *ServerInterfaceWrapper) GetSchoolMaterialArray(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetSchoolMaterialArrayParams
	// ------------- Optional query parameter "material_type" -------------

	err = runtime.BindQueryParameter("form", true, false, "material_type", ctx.QueryParams(), &params.MaterialType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter material_type: %s", err))
	}

	// ------------- Optional query parameter "from_created_at" -------------

	err = runtime.BindQueryParameter("form", true, false, "from_created_at", ctx.QueryParams(), &params.FromCreatedAt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter from_created_at: %s", err))
	}

	// ------------- Optional query parameter "to_created_at" -------------

	err = runtime.BindQueryParameter("form", true, false, "to_created_at", ctx.QueryParams(), &params.ToCreatedAt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter to_created_at: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "page_size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_size", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page_size: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetSchoolMaterialArray(ctx, params)
	return err
}

// CreateSchoolMaterial converts echo context to params.
func (w *ServerInterfaceWrapper) CreateSchoolMaterial(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateSchoolMaterial(ctx)
	return err
}

// GetSchoolMaterial converts echo context to params.
func (w *ServerInterfaceWrapper) GetSchoolMaterial(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "material_id" -------------
	var materialId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "material_id", ctx.Param("material_id"), &materialId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter material_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetSchoolMaterial(ctx, materialId)
	return err
}

// UpdateSchoolMaterial converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateSchoolMaterial(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "material_id" -------------
	var materialId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "material_id", ctx.Param("material_id"), &materialId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter material_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateSchoolMaterial(ctx, materialId)
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

	router.GET(baseURL+"/school/materials", wrapper.GetSchoolMaterialArray)
	router.POST(baseURL+"/school/materials", wrapper.CreateSchoolMaterial)
	router.GET(baseURL+"/school/materials/:material_id", wrapper.GetSchoolMaterial)
	router.PUT(baseURL+"/school/materials/:material_id", wrapper.UpdateSchoolMaterial)

}
