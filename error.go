package tvdb

import (
	"encoding/json"
	"fmt"
)

// JSONError is a generic type for errors
type JSONError struct {
	Message string `json:"error,omitempty"`
}

func (j *JSONError) UnmarshalJSON(p []byte) error {
	type MessageErr struct {
		Message string `json:"error,omitempty"`
	}
	var v MessageErr
	if err := json.Unmarshal(p, &v); err == nil {
		j.Message = v.Message
		return nil
	}
	/*
	   "errors": {
	   		 "invalidLanguage": "Incomplete or no translation for the given language"
	 	 }
	*/
	type ErrorMap map[string]string
	var v2 ErrorMap
	err := json.Unmarshal(p, &v2)
	if err == nil {
		j.Message = fmt.Sprintf("%q", v2)
	}
	return err
}

func (e JSONError) Error() string {
	return fmt.Sprintf("tvdb: %v", e.Message)
}

// Empty checks if an error message is empty
func (e JSONError) Empty() bool {
	if len(e.Message) == 0 {
		return true
	}
	return false
}

// relevantError returns an error or nil
// selects the right error based on the Empty() result
func relevantError(httpError error, jsonError *JSONError) error {
	if httpError != nil {
		return httpError
	}
	if jsonError != nil && !jsonError.Empty() {
		return jsonError
	}
	return nil
}
