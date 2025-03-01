package auth

import (
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /signup", signup)
	router.HandleFunc("POST /login", login)
	return router
}
