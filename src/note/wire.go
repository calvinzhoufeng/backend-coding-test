package note

type CreateNoteRequest struct {
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type Err struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
