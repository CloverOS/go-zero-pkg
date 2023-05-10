package docs

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func Run(config Config, server *rest.Server, handler http.HandlerFunc) {
	if config.Enable {
		server.AddRoute(rest.Route{
			Method:  http.MethodGet,
			Path:    "/swagger/*any",
			Handler: handler,
		})
	}
}
