package dto

type ErrorMessage struct {
	ErrorCode int32  `json:"errorCode"`
	Message   string `json:"message"`
}

type ResponseMessage struct {
	Item    any    `json:"item,omitempty"`
	Items   any    `json:"items,omitempty"`
	Message string `json:"message"`
}
