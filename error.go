package tvdb

import (
	"fmt"
)

// JSONError is a generic type for errors
type JSONError struct {
	Message string `json:"error,omitempty"`
}

func (e JSONError) Error() string {
	return fmt.Sprintf("tvdb: %v", e.Message)
}
