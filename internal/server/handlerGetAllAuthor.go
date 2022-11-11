package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetAllAuthor(c *fiber.Ctx) error {
	author, err := s.store.GetAllAuthors()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(200).JSON(author)
}
