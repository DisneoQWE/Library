package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) DeleteMember(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	c.ParamsParser(&param)
	err := s.store.DeleteMember(param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
