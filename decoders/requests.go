package decoders

import (
	"encoding/json"
	"net/http"
)

// RequestBody returns the body from a request
func RequestBody(rq *http.Request) (interface{}, error) {
	// Decode body
	var i interface{}
	err := json.NewDecoder(rq.Body).Decode(&i)
	if err != nil {
		return nil, err
	}
	return i, nil
}
