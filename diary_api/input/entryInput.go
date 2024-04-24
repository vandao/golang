package input

type EntryInput struct {
	Content string `json:"content" binding:"required"`
}
