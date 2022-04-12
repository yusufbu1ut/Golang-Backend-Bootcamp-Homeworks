package bookAuthor

import (
	"errors"

	"gorm.io/gorm"
)

var (
	declaredElement = errors.New("Element has been declared")
)

type BookAuthRepository struct {
	db *gorm.DB
}

func NewBookAuthRepository(db *gorm.DB) *BookAuthRepository {
	return &BookAuthRepository{
		db: db,
	}
}

func (r *BookAuthRepository) Migration() {
	r.db.AutoMigrate(&Book_Author{})
}

//InsertSampleData creates datas for book_author
func (r *BookAuthRepository) InsertSampleData(bookAuthhors []Book_Author) {
	for _, c := range bookAuthhors {
		r.db.Create(&c)
	}
}

// FindByISBN finds elements with isbn count it is for book connections
func (r *BookAuthRepository) FindByISBN(isbn int) ([]Book_Author, error) {
	var book_authors []Book_Author
	result := r.db.Find(&book_authors, "book_id = ?", isbn)

	if result.Error != nil {
		return nil, result.Error
	}

	return book_authors, nil
}

//FindByAuthorID finds elements with using author id
func (r *BookAuthRepository) FindByAuthorID(id int) ([]Book_Author, error) {
	var book_authors []Book_Author
	result := r.db.Find(&book_authors, "author_id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return book_authors, nil
}

//DeleteByISBN deletes items looking isbn colmn returns authors that have no book which as not deleted
func (r *BookAuthRepository) DeleteByISBN(isbn int) ([]int, error) {
	ba, err := r.FindByISBN(isbn)
	if err != nil {
		return nil, err
	}

	result := r.db.Delete(ba)
	if result.Error != nil {
		return nil, result.Error
	}

	var authorsID []int
	for _, v := range ba {
		// var book_authors []Book_Author
		// rslt := r.db.Find(&book_authors, "author_id = ?", v.AuthorID)
		// if rslt.Error != nil {
		// 	return nil, rslt.Error
		// }
		book_authors, err := r.FindByAuthorID(int(v.AuthorID))
		if err != nil {
			return nil, err
		}
		if len(book_authors) == 0 {
			authorsID = append(authorsID, int(v.AuthorID))
		}
	}

	return authorsID, nil
}

func (r *BookAuthRepository) DeleteByAuthorID(id int) ([]int, error) {
	ba, err := r.FindByAuthorID(id)
	if err != nil {
		return nil, err
	}

	result := r.db.Delete(ba)
	if result.Error != nil {
		return nil, result.Error
	}

	var booksId []int
	for _, v := range ba {

		book_authors, err := r.FindByAuthorID(v.BookID)
		if err != nil {
			return nil, err
		}
		if len(book_authors) == 0 {
			booksId = append(booksId, v.BookID)
		}
	}

	return booksId, nil
}

// Create creates new book_authors element
func (r *BookAuthRepository) Create(bauthors Book_Author) error {
	result := r.db.Where(&Book_Author{BookID: bauthors.BookID, AuthorID: bauthors.AuthorID}).First(&bauthors)
	if result.Error != nil {
		r.db.Create(&bauthors)
		return nil
	}

	return declaredElement
}
