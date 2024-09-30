package Util

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBodyFromJson(r *http.Response, x interface{}) {
	body, err := io.ReadAll(r.Body)
	CheckErrorAndLog(err, "Failed to read response body")

	err = json.Unmarshal(body, &x)
	CheckErrorAndLog(err, "Failed to unmarshal response body")
}
