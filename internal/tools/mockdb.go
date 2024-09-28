package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"first": {
		Username:  "first",
		AuthToken: "firstauthtoken",
	},
	"second": {
		Username:  "second",
		AuthToken: "secondauthtoken",
	},
	"third": {
		Username:  "third",
		AuthToken: "thirdauthtoken",
	},
}

var mockNotes = map[string]NoteDetails{
	"first": {
		Username: "first",
		Note:     "first first",
	},
	"second": {
		Username: "second",
		Note:     "second second",
	},
	"third": {
		Username: "third",
		Note:     "third third",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserNotes(username string) *NoteDetails {
	time.Sleep(time.Second * 1)

	var clientData = NoteDetails{}
	clientData, ok := mockNotes[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
