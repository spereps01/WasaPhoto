package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the fountain from the request body.
	var photo Photo
	var err error
	/*
		err := json.NewDecoder(r.Body).Decode(&photo)
		if err != nil {
			// The body was not a parseable JSON, reject it
			w.WriteHeader(http.StatusBadRequest)
			return
		}

			tk, err := strconv.Atoi(r.Header.Get("Authorization"))
			if err != nil {
				return
			}
	*/
	photo.User_id, err = strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo.Data = time.Now().Format("2006-01-02 15:04:05")
	photo.Photo, err = io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the fountain in the database. Note that this function will return a new instance of the fountain with the
	// same information, plus the ID.
	dbuser, err := rt.db.AddPhoto(photo.ToDatabase_p())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
	photo.FromDatabase_p(dbuser)

	// Send the output to the user.
	w.Header().Set("Content-Type", "image/png")
	_ = json.NewEncoder(w).Encode(photo)
}
