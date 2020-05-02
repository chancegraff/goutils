package loggers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HTTPLogger logs interactions
type HTTPLogger struct {
	name string
	wt   *http.ResponseWriter
	je   *json.Encoder
}

// NewHTTPLogger creates a new logger
func NewHTTPLogger(name string, wt *http.ResponseWriter) HTTPLogger {
	(*wt).Header().Set("Content-Type", "text/html; charset=utf-8")
	(*wt).Header().Set("Access-Control-Allow-Origin", "*")
	(*wt).Header().Set("Access-Control-Allow-Headers", "X-Token-Auth")

	return HTTPLogger{
		name: name,
		wt:   wt,
		je:   json.NewEncoder(*wt),
	}
}

// Okay records success
func (l *HTTPLogger) Okay(v interface{}) {
	log.Printf("%s::Success\r\n", l.name)
	(*l.wt).WriteHeader(http.StatusOK)
	l.je.Encode(v)
}

// Error records malfunctions
func (l *HTTPLogger) Error(err error, status int) {
	log.Printf("%s::Error:::%s, %s\r\n", l.name, fmt.Sprint(status), err)
	(*l.wt).WriteHeader(status)
	l.je.Encode(map[string]string{"status": "error"})
}
