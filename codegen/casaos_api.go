// Package codegen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package codegen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	Access_tokenScopes = "access_token.Scopes"
)

// BaseResponse defines model for BaseResponse.
type BaseResponse struct {
	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// HealthServices defines model for HealthServices.
type HealthServices struct {
	NotRunning *[]string `json:"not_running,omitempty"`
	Running    *[]string `json:"running,omitempty"`
}

// GetHealthServicesOK defines model for GetHealthServicesOK.
type GetHealthServicesOK struct {
	Data *HealthServices `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// ResponseInternalServerError defines model for ResponseInternalServerError.
type ResponseInternalServerError = BaseResponse

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get service status
	// (GET /health/services)
	GetHealth(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetHealth converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealth(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetHealth(ctx)
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

	router.GET(baseURL+"/health/services", wrapper.GetHealth)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xVYWsjNxD9K2LaD0nZeE1CoSzch1yvzYVQXJqDFmLjk7XjXV12JXVGcrIN+9+LpM35",
	"ErchLbSfbEmj997Mzjw9gLK9swaNZ6gegJCdNYxpcYH+PcrOt9dIO62QF1dxW1nj0fj4VzrXaSW9tqb8",
	"xNbEPVYt9jKddt1iC9XNA3xNuIUKvir3bGWO4/KtZPxlooWxeABH1iF5nUXU0iewlyCeqoRxHFfjOBZQ",
	"IyvSLsqDChZXMBbwSHVpPJKRXbyF9AORpX+U3OtTOlTyyC0yucjsMW5CiARPQKrnZemRWTbp4Cn2dCAI",
	"fSCDtdgMgjMN6xqF3grfIqHQLKQZoAC8l73rECqAAghlvTDdAJWngAX4wcUT9qRNk3N5Vu0Dacb6NQVj",
	"4oXqAbTHPu3veZRkaXnGGQIOWD5vSCI5xPVr8E4a6fFODq/HTekwqkDaD9ex9DkDqRQyr729xfTVdSxs",
	"i7JGggKM7CPGefCtJf1HapA9l3T6CodcKW229vALLcN8fqacVj4QpgUujRBC5AO2gRSKHmst3yzhyBFu",
	"kfhE2c7SSWoQrEQt6fZ4CYJJMfo3S2i9d1yVJcm7WaN9GzaBkaZ2ninbl5cKf21lhx9QtWVnG1v2Upsy",
	"F2/6WW+kMUjrCL82umn9+rv53N3PnGmW8G/FdhHoP1Tr73SiWG+6gC8L1n0jZBclfC9ZLq6zqP9fUVZT",
	"PuuCpcmqxPnPl8KR3ekaWfSaFXadNGgDix59a2sWW0ui1tstEhovWKGRpC3PIsqPloRmDhhnvBa1ZhWY",
	"tTVcCNehZBQ7zdpHKxA3F9q/DxtB6Cxrb2lYHT1WI1fiMP0s81hYEp+sNuLGBhLvNCtL9f52nTdmTVPe",
	"mt/PN5u3G/zteLZM46J9mt19wlDADonzkOxO47hah0Y6DRWczeazMyjASd+mGS3bZEMlf+FDDfrDcbtA",
	"Lyb7EOylDyzsVqBUrfg4Gcc3H8UEM4NESmmoL+t8PRteMscvXsfT+fzv3oHPceVfPaFjAd++5u5LL1Uy",
	"rtD3koYpxSmBKcVYYNkwVDeTXT+2DayemF56nZ/a3c1qXMWASMbpPFAHFZS706nJIQZM8M+rffRh8W5x",
	"vHfJZ+zx5v2Jl80F2eAywBT50xTygvLV+GcAAAD//zg681OwCAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
