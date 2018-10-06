package common

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// RespondWithJSON is the common function to be used by all the handlers while
// returning JSON data to the caller.
func RespondWithJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(w, ErrDataEncoding.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, err = io.Copy(w, &buf)
	if err != nil {
		log.Println("RespondWithJSON:", err)
	}
}

// DecodeFromJSON is the common function to be used by all the POST handlers for
// reading JSON data from the request body and performing input validation.
func DecodeFromJSON(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
