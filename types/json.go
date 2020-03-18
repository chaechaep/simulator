package types

// JSONResponse represents an HTTP response which contains a JSON body.
type JSONResponse struct {
	// HTTP status code.
	Code int
	// JSON represents the JSON that should be serialized and sent to the client
	JSON interface{}
	// Headers represent any headers that should be sent to the client
	Headers map[string]string
}

type JSONEmpty struct {
}
