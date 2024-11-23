package nrhttprouter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Router should be used in place of httprouter.Router.  Create it using
// New().
type Router struct {
	*httprouter.Router

	application *newrelic.Application
}

// New creates a new Router to be used in place of httprouter.Router.
func New(app *newrelic.Application) *Router {
	return &Router{
		Router:      httprouter.New(),
		application: app,
	}
}

func txnName(method, path string) string {
	return method + " " + path
}

func (r *Router) handle(method string, path string, original httprouter.Handle) {
	handle := original
	if nil != r.application {
		handle = func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
			txn := r.application.StartTransaction(txnName(req.Method, req.URL.Path))
			defer txn.End()

			txn.SetWebRequestHTTP(req)
			txn.SetWebResponse(w)

			req = newrelic.RequestWithTransactionContext(req, txn)

			original(w, req, ps)
		}
	}
	r.Router.Handle(method, path, handle)
}

// DELETE replaces httprouter.Router.DELETE.
func (r *Router) DELETE(path string, h httprouter.Handle) {
	r.handle(http.MethodDelete, path, h)
}

// GET replaces httprouter.Router.GET.
func (r *Router) GET(path string, h httprouter.Handle) {
	r.handle(http.MethodGet, path, h)
}

// HEAD replaces httprouter.Router.HEAD.
func (r *Router) HEAD(path string, h httprouter.Handle) {
	r.handle(http.MethodHead, path, h)
}

// OPTIONS replaces httprouter.Router.OPTIONS.
func (r *Router) OPTIONS(path string, h httprouter.Handle) {
	r.handle(http.MethodOptions, path, h)
}

// PATCH replaces httprouter.Router.PATCH.
func (r *Router) PATCH(path string, h httprouter.Handle) {
	r.handle(http.MethodPatch, path, h)
}

// POST replaces httprouter.Router.POST.
func (r *Router) POST(path string, h httprouter.Handle) {
	r.handle(http.MethodPost, path, h)
}

// PUT replaces httprouter.Router.PUT.
func (r *Router) PUT(path string, h httprouter.Handle) {
	r.handle(http.MethodPut, path, h)
}

// Handle replaces httprouter.Router.Handle.
func (r *Router) Handle(method, path string, h httprouter.Handle) {
	r.handle(method, path, h)
}

// Handler replaces httprouter.Router.Handler.
func (r *Router) Handler(method, path string, handler http.Handler) {
	_, h := newrelic.WrapHandle(r.application, txnName(method, path), handler)
	r.Router.Handler(method, path, h)
}

// HandlerFunc replaces httprouter.Router.HandlerFunc.
func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.Handler(method, path, handler)
}

// ServeHTTP replaces httprouter.Router.ServeHTTP.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if nil != r.application {
		h, _, _ := r.Router.Lookup(req.Method, req.URL.Path)
		if nil == h {
			txn := r.application.StartTransaction("NotFound")
			defer txn.End()

			txn.SetWebRequestHTTP(req)
			txn.SetWebResponse(w)

			req = newrelic.RequestWithTransactionContext(req, txn)
		}
	}

	r.Router.ServeHTTP(w, req)
}
