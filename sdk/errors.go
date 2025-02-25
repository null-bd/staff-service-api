package sdk

import "fmt"

type APIError struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Details []APIErrorDetail `json:"details,omitempty"`
}

type APIErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *APIError) IsNotFound() bool {
	return e.Code == "NOT_FOUND"
}

func (e *APIError) IsConflict() bool {
	return e.Code == "CONFLICT"
}

func (e *APIError) IsBadRequest() bool {
	return e.Code == "BAD_REQUEST"
}
