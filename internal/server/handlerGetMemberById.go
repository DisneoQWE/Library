package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetMemberById(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	c.ParamsParser(&param)
	member, err := s.store.GetMemberById(param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(200).JSON(member)
}
