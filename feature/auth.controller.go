package feature

import (
	"fmt"
	"net/http"

	"github.com/mnshah219/go_net_http/middleware"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(middleware.UserIDKey("userID")).(string)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Authenticated pong to userID %s", userID)))
	})
	return router
}
