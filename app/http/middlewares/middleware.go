package middlewares

import "net/http"

// HttpHandlerFunc
type HttpHandlerFunc func(http.ResponseWriter, *http.Request)
