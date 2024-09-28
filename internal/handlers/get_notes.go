package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MouraJr/firstgoapi/api"
	"github.com/MouraJr/firstgoapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var params = api.UserNotesParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.NoteDetails
	tokenDetails = (*database).GetUserNotes(params.Username)

	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var resp = api.UserNotesResponse{
		Notes: (*tokenDetails).Notes,
		Code:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
