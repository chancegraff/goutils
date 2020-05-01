package decoders

import (
	"encoding/json"
	"net/http"
)

// RequestBody returns the body from a request
func RequestBody(i interface{}, rq *http.Request) error {
	err := json.NewDecoder(rq.Body).Decode(&i)
	if err != nil {
		return err
	}
	return nil
}
