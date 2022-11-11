package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) DeleteAuthor(c *fiber.Ctx) error {
	var err error
	param := struct {
		ID int `params:"id"`
	}{}
	c.ParamsParser(&param)
	err = s.store.DeleteAuthor(param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
