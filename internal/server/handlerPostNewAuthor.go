package server

import (
	"RestApiLibrary/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) PostNewAuthor(c *fiber.Ctx) error {
	author := new(model.Author)
	if err := c.BodyParser(author); err != nil {
		return err
	}
	if author.AuthorId == 0 && author.AuthorSpecialization == "" && author.AuthorPseudonym == "" && author.AuthorFio == "" {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	s.store.PostNewAuthor(author)
	return c.Status(fiber.StatusCreated).JSON(author)
}
