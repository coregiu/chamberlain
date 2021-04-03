package server

type ChamberlainWebError struct {
	HttpResponseStatus int
	ErrorMessage string
}

func NewChamberlainError(httpResponseStatus int, errorMessage string) *ChamberlainWebError {
	return &ChamberlainWebError{HttpResponseStatus: httpResponseStatus, ErrorMessage: errorMessage}
}