// Code generated by go-swagger; DO NOT EDIT.

package export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetExportSidecarHandlerFunc turns a function with the right signature into a get export sidecar handler
type GetExportSidecarHandlerFunc func(GetExportSidecarParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetExportSidecarHandlerFunc) Handle(params GetExportSidecarParams) middleware.Responder {
	return fn(params)
}

// GetExportSidecarHandler interface for that can handle valid get export sidecar params
type GetExportSidecarHandler interface {
	Handle(GetExportSidecarParams) middleware.Responder
}

// NewGetExportSidecar creates a new http.Handler for the get export sidecar operation
func NewGetExportSidecar(ctx *middleware.Context, handler GetExportSidecarHandler) *GetExportSidecar {
	return &GetExportSidecar{Context: ctx, Handler: handler}
}

/*GetExportSidecar swagger:route GET /export/sidecar export getExportSidecar

Export a JSON file that can be used for flagr's sidecar mode

*/
type GetExportSidecar struct {
	Context *middleware.Context
	Handler GetExportSidecarHandler
}

func (o *GetExportSidecar) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetExportSidecarParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
