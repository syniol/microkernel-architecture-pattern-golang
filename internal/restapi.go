package internal

import (
	"architecture/pkg"
	"net/http"
)

func NewRESTfulAPI(rest *http.ServeMux) {
	for _, routeHandler := range []pkg.RestAPIPlugin{} {
		rest.HandleFunc(routeHandler())
	}
}
