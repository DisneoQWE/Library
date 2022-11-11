package server

import (
	"RestApiLibrary/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) getAllBooks(c *fiber.Ctx) error {
	books, err := s.store.GetAllBooks()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	bookResults := make([]model.BookResult, 0, 5)
	for i := range books {
		bookResult := model.BookResult{}
		bookResult.BookId = books[i].BookId
		bookResult.BookName = books[i].BookName
		bookResult.BookGenre = books[i].BookGenre
		bookResult.IsbnCode = books[i].IsbnCode
		bookResult.AuthorId = books[i].AuthorId
		bookResult.MemberID = model.ConvertFromUllIntToInt(books[i].MemberID)
		bookResults = append(bookResults, bookResult)
	}
	return c.Status(200).JSON(bookResults)
}
