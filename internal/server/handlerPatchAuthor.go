package server

import (
	"RestApiLibrary/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) UpdateAuthor(c *fiber.Ctx) error {
	param := struct {
		ID int `params:"id"`
	}{}
	var err error
	c.ParamsParser(&param)
	author := new(model.Author)
	if err = c.BodyParser(author); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	//
	authors := make([]model.Author, 0, 5)
	authors, err = s.store.GetAllAuthors()
	flag := false
	for index := range authors {
		if authors[index].AuthorId == param.ID {
			flag = true
		}
	}
	if !flag {
		return fiber.NewError(fiber.StatusNotFound)
	}
	//
	s.store.UpdateAuthor(author, param.ID)
	return c.Status(fiber.StatusOK).JSON(author)
}
