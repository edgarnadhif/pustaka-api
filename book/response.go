package book

// BookRequest contains response body of the Book API
type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Price       int    `json:"price"`
	Discount    int    `json:"discount"`
}
