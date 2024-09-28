package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username  string
	AuthToken string
}

type NoteDetails struct {
	Username string
	Notes    []string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserNotes(username string) *NoteDetails
	SetupDatabase() error
	// AddNote(username string, note string) error
	// RemoveNote(username string, note string) error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
