package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostSourcesIDKapacitorsKapaIDProxyHandlerFunc turns a function with the right signature into a post sources ID kapacitors kapa ID proxy handler
type PostSourcesIDKapacitorsKapaIDProxyHandlerFunc func(context.Context, PostSourcesIDKapacitorsKapaIDProxyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostSourcesIDKapacitorsKapaIDProxyHandlerFunc) Handle(ctx context.Context, params PostSourcesIDKapacitorsKapaIDProxyParams) middleware.Responder {
	return fn(ctx, params)
}

// PostSourcesIDKapacitorsKapaIDProxyHandler interface for that can handle valid post sources ID kapacitors kapa ID proxy params
type PostSourcesIDKapacitorsKapaIDProxyHandler interface {
	Handle(context.Context, PostSourcesIDKapacitorsKapaIDProxyParams) middleware.Responder
}

// NewPostSourcesIDKapacitorsKapaIDProxy creates a new http.Handler for the post sources ID kapacitors kapa ID proxy operation
func NewPostSourcesIDKapacitorsKapaIDProxy(ctx *middleware.Context, handler PostSourcesIDKapacitorsKapaIDProxyHandler) *PostSourcesIDKapacitorsKapaIDProxy {
	return &PostSourcesIDKapacitorsKapaIDProxy{Context: ctx, Handler: handler}
}

/*PostSourcesIDKapacitorsKapaIDProxy swagger:route POST /sources/{id}/kapacitors/{kapa_id}/proxy postSourcesIdKapacitorsKapaIdProxy

POST body directly to configured kapacitor.  The response and status code from kapacitor is directly returned.

*/
type PostSourcesIDKapacitorsKapaIDProxy struct {
	Context *middleware.Context
	Handler PostSourcesIDKapacitorsKapaIDProxyHandler
}

func (o *PostSourcesIDKapacitorsKapaIDProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostSourcesIDKapacitorsKapaIDProxyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
