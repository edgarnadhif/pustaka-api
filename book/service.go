package book

// Repository implements CRUD methods for Books table from the Repository struct
type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

// service implements Repository struct
type service struct {
	repository Repository
}

// NewService creates a new Service instance
func NewService(repository Repository) *service {
	return &service{repository}
}

// FindAll returns all data from the Books table
func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

// FindByID returns a single data from Books table based on the provided ID
func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

// Create creates a new data to the Books table based on the provided parameters
func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Rating:      int(rating),
		Price:       int(price),
		Discount:    int(discount),
	}

	return s.repository.Create(book)
}

// Update changes the existing data in the Books table based on the provided parameters
func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Price = int(price)
	book.Discount = int(discount)

	return s.repository.Update(book)
}

// Delete removes data from the Books table based on the provided parameters
func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	return s.repository.Delete(book)
}
