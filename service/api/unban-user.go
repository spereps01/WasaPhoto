package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id1, err := strconv.ParseUint(ps.ByName("id1"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id2, err := strconv.ParseUint(ps.ByName("id2"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbuser, err := rt.db.UnbanUser(int(id1), int(id2))
	if err != nil {
		ctx.Logger.WithError(err).WithField("id1", id1).Error("can't unban the user")
		ctx.Logger.WithError(err).WithField("id2", id2).Error("can't unban the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dbuser)
}
