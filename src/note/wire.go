package note

type CreateNoteRequest struct {
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type Err struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
