package server

import (
	"RestApiLibrary/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) postNewBook(c *fiber.Ctx) error {

	bookResult := new(model.BookResult)
	if err := c.BodyParser(bookResult); err != nil {
		return err
	}
	book := new(model.Book)
	book.BookId = bookResult.BookId
	book.BookName = bookResult.BookName
	book.BookGenre = bookResult.BookGenre
	book.IsbnCode = bookResult.IsbnCode
	book.AuthorId = bookResult.AuthorId
	book.MemberID = model.ConvertNullInt(bookResult.MemberID)
	s.store.PostNewBook(book)
	return c.Status(fiber.StatusCreated).JSON(book)
}
