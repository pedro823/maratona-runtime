package errors

import (
	"encoding/json"
	"github.com/pedro823/maratona-runtime/util"
	"net/http"
)

type HTTPError interface {
	WriteResponse(res http.ResponseWriter)
	WriteJSON(res *util.JSONRenderer)
}

type genericHTTPError struct {
	status int
	reason string
}

func (e genericHTTPError) WriteResponse(res http.ResponseWriter) {
	body := map[string]string{"error": e.reason}
	marshaledBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	res.WriteHeader(e.status)
	_, err = res.Write(marshaledBody)
	if err != nil {
		panic(err)
	}
}

func (e genericHTTPError) WriteJSON(res *util.JSONRenderer) {
	res.JSON(e.status, map[string]string{"error": e.reason})
}

func NewHTTPError(status int, reason string) HTTPError {
	return &genericHTTPError{status: status, reason: reason}
}
