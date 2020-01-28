package middleware

import (
	"github.com/go-pg/pg"
	"github.com/pedro823/maratona-runtime/errors"
	"github.com/pedro823/maratona-runtime/model"
	"github.com/pedro823/maratona-runtime/util"
	"net/http"
)

func RequireCTFID(req *http.Request, res *util.JSONRenderer, db *pg.DB, context util.ContextMap) {
	CTFID := req.Header.Get("CTFID")

	user, httpErr := verify(CTFID, db)
	if httpErr != nil {
		httpErr.WriteJSON(res)
	}

	context[util.UserContextKey] = user
}

func verify(CTFID string, db *pg.DB) (*model.User, errors.HTTPError) {

}
