package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	idp, err := strconv.ParseUint(ps.ByName("idp"), 10, 64)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idc, err := strconv.ParseUint(ps.ByName("idc"), 10, 64)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// The client is not supposed to send us the ID in the body, as the fountain ID is already specified in the path,
	// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).

	// Update the fountain in the database.
	dbuser, err := rt.db.AddComment(comment.ToDatabase_c())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the fountain that triggered the error.
		ctx.Logger.WithError(err).WithField("idp", idp).Error("can't add the comment")
		ctx.Logger.WithError(err).WithField("idc", idc).Error("can't add the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dbuser)
}
