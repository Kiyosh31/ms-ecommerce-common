package customlogger

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ReadRequestPayload(r *http.Request) string {
	// Read the request body
	bodyBytes, _ := io.ReadAll(r.Body)
	// if err != nil {
	// 	return "nil", fmt.Errorf("error reading req body: %v", err)
	// }

	// Reset the request body so it can be read again later
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Format the request body for logging
	var bodyMap map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &bodyMap); err == nil {
		compactBody, _ := json.Marshal(bodyMap) // Compact JSON format
		return string(compactBody)
	} else {
		// Log as raw bytes if unmarshalling fails
		return string(bodyBytes)
	}
}
