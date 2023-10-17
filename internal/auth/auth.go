package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts API key from headers
// of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authantication header found")
	}

	 vals := strings.Split(val, " ")
	 if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	 }

	 if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of api key")
	 }
	 return vals[1], nil
}