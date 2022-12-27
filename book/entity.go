package book

import "time"

// Book contains attributes for Books table
type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
