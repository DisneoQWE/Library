package store

import (
	"RestApiLibrary/internal/model"
)

// GET: Get All Books
func (a *DBConnection) GetAllBooks() ([]model.Book, error) {
	books := make([]model.Book, 0, 5) //make create a heap in memory for a slice maps and initializes and puts zero or empty
	rows, err := a.db.NamedQuery("SELECT * FROM books", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(&book.BookId, &book.BookName, &book.BookGenre, &book.IsbnCode, &book.AuthorId, &book.MemberID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// POST : Create new book
func (a *DBConnection) PostNewBook(book *model.Book) error {
	bookResult := model.BookResult{
		BookId:    book.BookId,
		BookName:  book.BookName,
		BookGenre: book.BookGenre,
		IsbnCode:  book.IsbnCode,
		AuthorId:  book.AuthorId,
		MemberID:  model.ConvertFromUllIntToInt(book.MemberID),
	}
	_, err := a.db.NamedExec(`INSERT INTO books (bookname, bookgenre, bookisbn, author_id, memberid)
		VALUES (:bookname,:bookgenre,:bookisbn,:author_id,:memberid)`,
		map[string]interface{}{
			"bookname":  bookResult.BookName,
			"bookgenre": bookResult.BookGenre,
			"bookisbn":  bookResult.IsbnCode,
			"author_id": bookResult.AuthorId,
			"memberid":  bookResult.MemberID})
	if err != nil {
		return err
	}
	return nil
}

// PATCH: Update a book
func (a *DBConnection) PatchNewBook(book *model.Book, bookId int) error {
	var err error
	if book.BookName != "" {
		_, err = a.db.NamedExec(`UPDATE books SET bookname=:bookname WHERE bookid=:bookid`, map[string]interface{}{
			"bookid":   book.BookId,
			"bookname": book.BookName})
	}
	if book.BookGenre != "" {
		_, err = a.db.NamedExec(`UPDATE books SET bookgenre=:bookgenre WHERE bookid=:bookid`, map[string]interface{}{
			"bookid":    book.BookId,
			"bookgenre": book.BookGenre})
	}
	if book.IsbnCode != 0 {
		_, err = a.db.NamedExec(`UPDATE books SET bookisbn=:bookisbn WHERE bookid=:bookid`, map[string]interface{}{
			"bookid":   book.BookId,
			"bookisbn": book.IsbnCode})
	}
	_, err = a.db.NamedExec(`UPDATE books SET author_id=:author_id WHERE bookid=:bookid`, map[string]interface{}{
		"author_id": book.AuthorId,
		"bookid":    book.BookId})

	a.CreateZeroMember()
	_, err = a.db.NamedExec(`UPDATE books SET memberid=:member_id WHERE bookid=:bookid`, map[string]interface{}{
		"member_id": model.ConvertFromUllIntToInt(book.MemberID),
		"bookid":    book.BookId})
	if err != nil {
		return err
	}
	return nil
}

// DELETE: Delete book
func (a *DBConnection) DeleteBook(bookId int) error {
	_, err := a.db.NamedQuery(`DELETE FROM books WHERE bookid=:bookId`, map[string]interface{}{
		"bookId": bookId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *DBConnection) GetBookById(bookId int) (*model.Book, error) {
	bookRes := new(model.BookResult)
	rows, err := a.db.NamedQuery(`SELECT * FROM books WHERE bookid=:bookid`, map[string]interface{}{
		"bookid": bookId})
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&bookRes.BookId, &bookRes.BookName, &bookRes.BookGenre, &bookRes.IsbnCode, &bookRes.AuthorId, &bookRes.MemberID)
	}
	book := new(model.Book)
	book.BookId = bookRes.BookId
	book.BookName = bookRes.BookName
	book.BookGenre = bookRes.BookGenre
	book.IsbnCode = bookRes.IsbnCode
	book.AuthorId = bookRes.AuthorId
	book.MemberID = model.ConvertNullInt(bookRes.MemberID)
	if err != nil {
		return nil, err
	}
	return book, nil
}
