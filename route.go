package main

import (
	"net/http"

	"github.com/Mnshah219/go_net_http/auth"
	"github.com/Mnshah219/go_net_http/middleware"
	"github.com/justinas/alice"
)

func registerRoutes() *http.ServeMux {
	middlewareChain := alice.New(middleware.LoggingMiddleware, middleware.RecoveryMiddleware)

	router := http.NewServeMux()
	router.Handle("/auth/", middlewareChain.Then(http.StripPrefix("/auth", auth.RegisterRoutes())))
	return router
}
