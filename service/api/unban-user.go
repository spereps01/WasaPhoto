package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	idu, err := strconv.ParseUint(ps.ByName("idu"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idb, err := strconv.ParseUint(ps.ByName("idb"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbuser, err := rt.db.UnbanUser(int(idu), int(idb))
	if err != nil {
		ctx.Logger.WithError(err).WithField("idu", idu).Error("can't delete the follower")
		ctx.Logger.WithError(err).WithField("idb", idb).Error("can't delete the follower")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dbuser)
}
