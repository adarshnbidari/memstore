package response

//JSONResponse for response format in json
type JSONResponse struct {
	Status string `json:"status"`

	Message string `json:"message,omitempty"`

	Body interface{} `json:"body,omitempty"`
}
