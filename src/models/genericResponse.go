package models

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorResponse struct {
	Message      []ErrorMessage `json:"errors,omitempty"`
	ErrorMessage string         `json:"error,omitempty"`
}
