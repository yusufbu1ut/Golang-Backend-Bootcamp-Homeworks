package book

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

var (
	hihgerAmount = errors.New("Given count is higher than book' amount.")
	declaredBook = errors.New("Book has been declared")
)

type BookRepository struct {
	db *gorm.DB
	sync.RWMutex
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&Book{})
}

// InsertSampleData creates for db books when program runs
func (r *BookRepository) InsertSampleData(books []Book) {
	for _, c := range books {
		r.db.Create(&c)
	}
}

// GetAllBooks gives all books
func (r *BookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// FindByBookName finds books with given input
func (r *BookRepository) FindByName(input string) ([]Book, error) {
	var books []Book
	result := r.db.Where("LOWER(name) LIKE LOWER(?)", "%"+input+"%").Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// FindByID finds book with parameter id
func (r *BookRepository) FindByID(id int) ([]Book, error) {
	var books []Book
	result := r.db.Find(&books, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

//FindByISBN finds books with ISBN uniqe count it is for book_authors base
func (r *BookRepository) FindByISBN(isbn int) ([]Book, error) {
	var books []Book
	result := r.db.Find(&books, "isbn = ?", isbn)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// Deletes book with given book
func (r *BookRepository) Delete(b Book) error {
	_, err := r.FindByID(int(b.ID))
	if err != nil {
		return err
	}
	result := r.db.Delete(b)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Deletes book with given id
func (r *BookRepository) DeleteByID(id int) ([]Book, error) {
	book, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}

	result := r.db.Delete(&Book{}, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

// Create creates book if it is not exist and uniqe looking on isbn name publisher
func (r *BookRepository) Create(book Book) error {
	result, _ := r.FindByISBN(book.ISBN)
	if len(result) != 0 {
		return declaredBook
	}
	r.db.Create(&book)
	return nil
}

// Update saves new changed book writing on db same id
func (r *BookRepository) Update(book Book) error {
	r.db.Save(book)
	return nil
}

// Buy decrease book stock amount with given count parameter
func (r *BookRepository) Buy(book Book, count int) error {
	r.Lock()
	defer r.Unlock()

	if book.StockAmount < count {
		return hihgerAmount
	}

	book.StockAmount -= count

	r.Update(book)

	return nil
}
