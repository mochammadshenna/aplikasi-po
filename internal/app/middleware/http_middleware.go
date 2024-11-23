package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	nr "github.com/mochammadshenna/aplikasi-po/internal/app/newrelic"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/cors"
)

type HttpMiddleware struct {
	Router  *httprouter.Router
	Handler http.Handler
}

type Middleware interface {
	Chain(http.Handler) http.Handler
}

func NewHttpMiddleware(router *httprouter.Router) *HttpMiddleware {
	return &HttpMiddleware{
		Router:  router,
		Handler: cors.AllowAll().Handler(router),
	}
}

func MultipleMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	if len(middlewares) < 1 {
		return h
	}

	for i := range middlewares {
		h = middlewares[len(middlewares)-1-i].Chain(h)
	}

	return h
}

func (middleware *HttpMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var txn *newrelic.Transaction
	// check for not found router
	h, _, _ := middleware.Router.Lookup(request.Method, request.URL.Path)
	if nil == h {
		txn = nr.App.StartTransaction("NotFound")
	} else {
		txn = nr.App.StartTransaction(txnName(request.Method, request.URL.Path))
	}
	txn.AcceptDistributedTraceHeaders(newrelic.TransportHTTPS, request.Header)
	defer txn.End()

	txn.SetWebRequestHTTP(request)
	writer = txn.SetWebResponse(writer)
	request = newrelic.RequestWithTransactionContext(request, txn)

	txn.AddAttribute("request.headers.xForwardedFor", request.Header.Get("X-Forwarded-For"))

	// wrap writer to capture response code
	// set default status 200, for more information why visit https://pkg.go.dev/net/http#ResponseWriter.WriteHeader
	wrappeWriter := wrapResponseWriter(writer, 200)
	w, r := PreHandle(wrappeWriter, request)
	middleware.Handler.ServeHTTP(w, r)
	PostHandle(w, r)
}

func txnName(method, path string) string {
	return method + " " + path
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter, status int) *responseWriter {
	return &responseWriter{ResponseWriter: w, status: status}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}
