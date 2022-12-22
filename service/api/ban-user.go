package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The Fountain ID in the path is a 64-bit unsigned integer. Let's parse it.
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
	// Read the new content for the fountain from the request body.

	/*	err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			// The body was not a parseable JSON, reject it
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	*/
	// Update the fountain in the database.

	dbuser, err := rt.db.BanUser(int(idu), int(idb))
	rt.db.UnfollowUser(int(idb), int(idu))
	rt.db.UnfollowUser(int(idu), int(idb))
	if err != nil {
		ctx.Logger.WithError(err).WithField("idb", idb).Error("can't ban the user")
		ctx.Logger.WithError(err).WithField("idu", idu).Error("can't ban the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dbuser)
}
