package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) AllOtherRequests(c *fiber.Ctx) error {
	return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{})
}
