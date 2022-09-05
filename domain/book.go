package domain

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Book struct {
	ID     uint    `json:"id"`
	Author Author  `json:"author"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	ISBN   string  `json:"isbn"`
	Stock  int64   `json:"stock"`
}
