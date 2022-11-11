package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetAuthorByIdAndBook(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	c.ParamsParser(&param)
	author, err := s.store.GetAllAuthorsById(param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	if author.AuthorId == 0 && author.AuthorPseudonym == "" && author.AuthorSpecialization == "" && author.AuthorFio == "" {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.Status(200).JSON(author)
}
