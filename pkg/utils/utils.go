package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return errors.New("failed to read request body")
	}

	if err := json.Unmarshal(body, x); err != nil {
		return errors.New("invalid JSON format")
	}

	return nil
}