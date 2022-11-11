package server

import (
	"RestApiLibrary/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetBookById(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	c.ParamsParser(&param)
	book, err := s.store.GetBookById(param.ID)
	bookResult := new(model.BookResult)
	bookResult.BookId = book.BookId
	bookResult.BookName = book.BookName
	bookResult.BookGenre = book.BookGenre
	bookResult.IsbnCode = book.IsbnCode
	bookResult.AuthorId = book.AuthorId
	bookResult.MemberID = model.ConvertFromUllIntToInt(book.MemberID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(200).JSON(bookResult)
}
