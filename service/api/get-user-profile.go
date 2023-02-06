package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	username := ps.ByName("username")
	tk, err := strconv.Atoi(r.Header.Get("Authorization"))
	if err != nil {
		return
	}
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbuser, err := rt.db.GetUserProfile(username, tk)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't find the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dbuser)
}
