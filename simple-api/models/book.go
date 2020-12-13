package models

// Book struct
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
