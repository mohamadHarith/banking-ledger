package dto

type ResponseMessage struct {
	ErrorCode int32  `json:"errorCode,omitempty"`
	Item      any    `json:"item,omitempty"`
	Items     any    `json:"items,omitempty"`
	Message   string `json:"message"`
}
