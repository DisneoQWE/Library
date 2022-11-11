package server

import (
	"RestApiLibrary/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) patchBook(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	c.ParamsParser(&param)
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
	//
	flag := false
	s.store.GetAllBooks()
	books, err := s.store.GetAllBooks()
	for i := range books {
		if model.ConvertFromUllIntToInt(books[i].MemberID) == model.ConvertFromUllIntToInt(book.MemberID) {
			s.store.CreateZeroMember()
			flag = true
		}
	}
	for i := range books {
		if books[i].AuthorId == book.AuthorId {
			flag = true
		}
	}
	if !flag {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	//
	err = s.store.PatchNewBook(book, param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
