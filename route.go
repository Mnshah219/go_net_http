package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/mnshah219/go_net_http/auth"
	"github.com/mnshah219/go_net_http/middleware"
)

func registerRoutes() *http.ServeMux {
	middlewareChain := alice.New(middleware.RecoveryMiddleware, middleware.LoggingMiddleware)
	authMiddlewareChain := middlewareChain.Append(middleware.AuthMiddleware)
	router := http.NewServeMux()
	router.Handle("/auth/", middlewareChain.Then(http.StripPrefix("/auth", auth.RegisterRoutes())))
	router.Handle("/feature/", authMiddlewareChain.Then(http.StripPrefix("/feature", auth.RegisterRoutes())))
	return router
}
