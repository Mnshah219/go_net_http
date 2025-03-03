package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mnshah219/go_net_http/auth/dto"
	"github.com/mnshah219/go_net_http/auth/utils"
	commonUtils "github.com/mnshah219/go_net_http/utils"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var payload dto.SignupDto
	validatonError := commonUtils.UnmarshalJSON(w, r, &payload)
	if validatonError != nil {
		http.Error(w, validatonError.Error(), validatonError.Status)
		return
	}
	existingUser := findOneUser(map[string]any{"email": payload.Email})
	if existingUser.ID != "" {
		http.Error(w, "Err: User with given email exists!", http.StatusBadRequest)
		return
	}
	// hash passwd
	// for more std implementation https://snyk.io/blog/secure-password-hashing-in-go/
	salt := os.Getenv("SALT")
	payload.Password = utils.GenerateHash([]byte(payload.Password), []byte(salt))
	_, err := createUser(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	resp, _ := json.Marshal(struct {
		User dto.SignupDto `json:"user"`
	}{User: payload})
	w.Write(resp)
}

func login(w http.ResponseWriter, r *http.Request) {
	var cred dto.LoginRequestDto
	validatonError := commonUtils.UnmarshalJSON(w, r, &cred)
	if validatonError != nil {
		http.Error(w, validatonError.Error(), validatonError.Status)
		return
	}
	user := findOneUser(map[string]any{"email": cred.Email})
	if user.ID == "" {
		http.Error(w, "Err: User with given email does not exists!", http.StatusNotFound)
		return
	}
	salt := os.Getenv("SALT")
	err := utils.Compare(user.Password, []byte(salt), []byte(cred.Password))
	if err != nil {
		http.Error(w, "Err: Incorrect credentials!", http.StatusBadRequest)
		return
	}
	jwt := utils.IssueJWT(user.ID)
	resp, _ := json.Marshal(dto.LoginResponseDto{Token: jwt})
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
