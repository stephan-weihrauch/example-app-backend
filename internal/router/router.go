package router

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const appVersion = "1.0.0"
const logoPath = ""

// CORS options
// https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
// https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
var allowedOrigins = []string{
	"*",
}
var allowedHeaders = []string{
	"Authorization",
	"Content-Type",
}

var handler http.Handler

func init() {
	r := mux.NewRouter()
	// Products api
	r.HandleFunc("/api/products", getProducts).Methods(http.MethodGet)
	r.HandleFunc("/api/products", createProduct).Methods(http.MethodPost)
	r.HandleFunc("/api/products", deleteProduct).Methods(http.MethodDelete)
	r.HandleFunc("/api/product-fields", getProductFields).Methods(http.MethodGet)
	// Other endpoints
	r.HandleFunc("/api/logo", getLogo).Methods(http.MethodGet)
	r.HandleFunc("/api/version", getVersion).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))

	handler = handlers.CORS(
		handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowedOrigins(allowedOrigins),
	)(r)
}

// HandleRequest handles the request based on the configured
func HandleRequest(res http.ResponseWriter, req *http.Request) {
	handler.ServeHTTP(res, req)
}

func getLogo(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	io.WriteString(res, fmt.Sprintf("Not implemented yet! Comming soon"))
	//http.ServeFile(res, req, logoPath)
}

func getVersion(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, fmt.Sprintf(`{"version": "%s"}`, appVersion))
}

// TODO implement
func getProducts(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	io.WriteString(res, fmt.Sprintf("Not implemented yet! Comming soon"))
	//products.GetProducts()
}

// TODO implement
func createProduct(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	io.WriteString(res, fmt.Sprintf("Not implemented yet! Comming soon"))
	//products.CreateProduct()
}

// TODO implement
func deleteProduct(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	io.WriteString(res, fmt.Sprintf("Not implemented yet! Comming soon"))
	//products.DeleteProduct()
}

// TODO implement
func getProductFields(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	io.WriteString(res, fmt.Sprintf("Not implemented yet! Comming soon"))
	//products.GetFields()
}
