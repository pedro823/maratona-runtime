package middleware

import (
	"net/http"
	"crypto/subtle"
	"os"
)

var masterToken = []byte(os.Getenv("CHALLENGE_MASTER_TOKEN"))

func RequireToken(req *http.Request, res http.ResponseWriter) {
	authToken := []byte(req.Header.Get("Authorization"))
	
	if subtle.ConstantTimeCompare(masterToken, authToken) != 1 {
		res.WriteHeader(http.StatusUnauthorized)
	}
}