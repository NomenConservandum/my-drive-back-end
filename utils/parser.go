package utils

import (
	"encoding/json"
	"errors"
	"io"
	"myDrive/db"
	"net/http"
)

func JsonToUser(r *http.Request) (db.User, error) {
	var u db.User

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return u, errors.New("error reading request body")
	}

	// Convert bytes to string
	bodyString := string(bodyBytes)

	err = json.Unmarshal([]byte(bodyString), &u)
	if err != nil {
		return u, errors.New("error parsing JSON: " + err.Error())
	}
	return u, nil

}

func JsonToTokens(r *http.Request) (db.Tokens, error) {
	var t db.Tokens

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return t, errors.New("error reading request body")
	}

	// Convert bytes to string
	bodyString := string(bodyBytes)

	err = json.Unmarshal([]byte(bodyString), &t)
	if err != nil {
		return t, errors.New("error parsing JSON: " + err.Error())
	}
	return t, nil

}
