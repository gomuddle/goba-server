package goba_server

// ErrorResponse represents a response to a failed request.
type ErrorResponse struct {
	// Error describes the reason why the request failed.
	Error string `json:"error,omitempty"`
}
