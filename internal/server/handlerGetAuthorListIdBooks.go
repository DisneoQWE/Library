package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetAuthorListIdBook(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	c.ParamsParser(&param)
	result, err := s.store.GetAuthorListIdBook(param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(200).JSON(result)
}
