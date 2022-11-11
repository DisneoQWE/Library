package server

import (
	"RestApiLibrary/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) PatchNewMember(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	var err error
	c.ParamsParser(&param)
	member := new(model.Member)
	if err = c.BodyParser(member); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	//
	members := make([]model.Member, 0, 5)
	members, err = s.store.GetMembers()
	flag := false
	for index := range members {
		if members[index].MemberId == param.ID {
			flag = true
		}
	}
	if !flag {
		return fiber.NewError(fiber.StatusNotFound)
	}
	//
	err = s.store.PatchNewMember(member, param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(member)
}
