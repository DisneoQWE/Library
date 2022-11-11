package server

import (
	"RestApiLibrary/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) PostNewMember(c *fiber.Ctx) error {
	member := new(model.Member)
	if err := c.BodyParser(member); err != nil {
		return err
	}
	s.store.PostNewMember(member)
	return c.Status(fiber.StatusCreated).JSON(member)
}
