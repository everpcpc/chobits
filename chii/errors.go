package chii

import (
	"fmt"
)

// https://github.com/bangumi/api/blob/master/docs-raw/Define.md#status-code

// APIError represents a Bangumi API Error response
type APIError struct {
	Message string `json:"error"`
	Code    int    `json:"code"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("bangumi: %d %v", e.Code, e.Message)
}

// Empty returns true if empty.
// FIXME(everpcpc)
func (e APIError) Empty() bool {
	if e.Message == "" {
		return true
	}
	return false
}

func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}
	return apiError
}
