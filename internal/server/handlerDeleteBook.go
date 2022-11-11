package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) DeleteBook(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	c.ParamsParser(&param)
	err := s.store.DeleteBook(param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
