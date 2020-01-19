package middleware

import (
	"crypto/subtle"
	"github.com/pedro823/maratona-runtime/errors"
	"net/http"
	"os"
)

var masterToken = []byte(os.Getenv("CHALLENGE_MASTER_TOKEN"))

func RequireToken(req *http.Request, res http.ResponseWriter) {
	authToken := []byte(req.Header.Get("Authorization"))

	if subtle.ConstantTimeCompare(masterToken, authToken) != 1 {
		errors.NewHTTPError(http.StatusUnauthorized, "Unauthorized").WriteResponse(res)
	}
}