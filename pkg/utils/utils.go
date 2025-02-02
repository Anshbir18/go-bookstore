// we will have code that would unmarshal the data that we recieve from the request

package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the request body and unmarshals it into the provided interface.
func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	json.Unmarshal(body, x)
}
