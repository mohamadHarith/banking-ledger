package dto

type Pagination struct {
	TotalRecords uint32 `json:"totalRecords"`
	CurrentPage  uint32 `json:"currentPage"`
	TotalPages   uint32 `json:"totalPages"`
	NextPage     uint32 `json:"nextPage"`
	PrevPage     uint32 `json:"prevPage"`
}

type ResponseMessage struct {
	ErrorCode  int32       `json:"errorCode,omitempty"`
	Item       any         `json:"item,omitempty"`
	Items      []any       `json:"items,omitempty"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
