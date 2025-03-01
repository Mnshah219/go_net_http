package auth

import (
	"log/slog"
	"net/http"

	"github.com/Mnshah219/go_net_http/auth/dto"
	"github.com/Mnshah219/go_net_http/utils"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var payload dto.SignupDto
	err := utils.UnmarshalJSON(w, r, &payload)
	if err != nil {
		http.Error(w, err.Error(), err.Status)
		return
	}
	slog.Info("Signup request received", "payload", payload)
}

func login(w http.ResponseWriter, r *http.Request) {}
