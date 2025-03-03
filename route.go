package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/mnshah219/go_net_http/auth"
	"github.com/mnshah219/go_net_http/middleware"
)

func registerRoutes() *http.ServeMux {
	middlewareChain := alice.New(middleware.LoggingMiddleware, middleware.RecoveryMiddleware)

	router := http.NewServeMux()
	router.Handle("/auth/", middlewareChain.Then(http.StripPrefix("/auth", auth.RegisterRoutes())))
	return router
}
