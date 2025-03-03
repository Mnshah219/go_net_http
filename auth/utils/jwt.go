package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IssueJWT(userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    os.Getenv("JWT_ISSUER"),
		Subject:   userID,
		Audience:  []string{"WEB_APP"},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(12))),
	})
	secret := os.Getenv("JWT_SECRET")
	jwt, _ := token.SignedString([]byte(secret))
	return jwt
}

func VerifyJWT(jwtToken string) (userID string, err error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	if err != nil {
		return "", err
	}
	return token.Claims.GetSubject()
}
