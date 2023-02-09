// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// Book defines model for book.
type Book struct {
	// BookId the id of the book
	BookId openapi_types.UUID `json:"bookId"`

	// Code the code of the book
	Code string `json:"code"`

	// PublicDate date of publication
	PublicDate *openapi_types.Date `json:"publicDate,omitempty"`

	// Publisher name of the publisher
	Publisher *string `json:"publisher,omitempty"`

	// Title the id of the book
	Title *string `json:"title,omitempty"`
}

// Title defines model for title.
type Title = string

// GetBooksParams defines parameters for GetBooks.
type GetBooksParams struct {
	// Title the title of the book
	Title *Title `form:"title,omitempty" json:"title,omitempty"`
}

// CreateBookJSONRequestBody defines body for CreateBook for application/json ContentType.
type CreateBookJSONRequestBody = Book

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of books
	// (GET /api/v1/books)
	GetBooks(c *gin.Context, params GetBooksParams)
	// Add a new book
	// (POST /api/v1/books)
	CreateBook(c *gin.Context)
	// Get a book by id
	// (GET /api/v1/books/{bookId})
	GetBook(c *gin.Context, bookId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetBooks operation middleware
func (siw *ServerInterfaceWrapper) GetBooks(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBooksParams

	// ------------- Optional query parameter "title" -------------

	err = runtime.BindQueryParameter("form", true, false, "title", c.Request.URL.Query(), &params.Title)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter title: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetBooks(c, params)
}

// CreateBook operation middleware
func (siw *ServerInterfaceWrapper) CreateBook(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreateBook(c)
}

// GetBook operation middleware
func (siw *ServerInterfaceWrapper) GetBook(c *gin.Context) {

	var err error

	// ------------- Path parameter "bookId" -------------
	var bookId string

	err = runtime.BindStyledParameter("simple", false, "bookId", c.Param("bookId"), &bookId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter bookId: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetBook(c, bookId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/v1/books", wrapper.GetBooks)

	router.POST(options.BaseURL+"/api/v1/books", wrapper.CreateBook)

	router.GET(options.BaseURL+"/api/v1/books/:bookId", wrapper.GetBook)

	return router
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RUPW/bMBD9K8S1o2rJbYdCW9wCQYAOHboFHs7i2WIikQx5cmEY+u/FUXZs2UrSokMn",
	"m7rH+3jv8fZQudY7S5YjlHvwGLAlppBObLgh+aMpVsF4Ns5CCVyTSiHl1koOK+ceIQMjwaeOwg4ysNiS",
	"YFOKDGJVU4sp6c5LIHIwdgN93x+DqWRKJY0E5ymwoeevd3q6E6Mv2li70CJDCV1nNGSXBTOonH5hKolc",
	"ZLu67btVY6pvyBM5NHK6P2AwfT1rSMIvpow1heuMQuOxoxNsIscrWl0xdClBBoGeOhNIQ3l/JPvA0/I5",
	"NSxGt93qgSoeBDR27c4MA7dOLbB6JKvVT2p9I7Tc/LiDDLYU4tBZMStmc+ncebLoDZTwaVbMCsjAI9dJ",
	"9xwr7pBdyGvChmso930GOXqTb+e5NJpgG2L5EcckzsUqcEu8SIBs5Or7PbwPtIYS3uUn7+cnSD4M0S+F",
	"leidjYMHPxaF/FTOMtlUD70/qpw/RBlqf+Zzw9Smi1PlDo5PIyT1Bk4xBNwNlI5l3BCrxkQWIdPYCrdo",
	"GlxJp/KCurbFsBvGVjjGima4iUdtP2znsBTXuThB29dAyHSQWmxBkRdO7/5q9rdHHpuOQ0f9P/L9ZzXH",
	"tFZpWIXK0i911OLzUHiMXKFWBzYyJQeNjBfU32h9nmqK9Uv35vvhtfVv2fjaxWnbyks5Ldvnlztm9rXt",
	"u/wPrIuZnR22kVphJK2cVfJelNGTdh6QOzXs82tW5QqF7ZGacbnvrsJGadpS43xLltWAhQy60EAJNbMv",
	"87wRXO0il1+KYg79sv8dAAD//+14wYUeBwAA",
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