package server

import "github.com/gofiber/fiber/v2"

func (s *Server) GetMembers(c *fiber.Ctx) error {
	author, err := s.store.GetMembers()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(200).JSON(author)
}
