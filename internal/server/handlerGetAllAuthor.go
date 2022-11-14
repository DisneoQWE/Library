package server

import (
	"github.com/gofiber/fiber/v2"
)

// GetAllAuthor godoc
// @Summary     Get author
// @Description get all authors
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       id  path     int true "Account ID"
// @Success     200 {object} model.Author
// @Failure     400 {object} model.Author
// @Failure     404 {object} model.Author
// @Failure     500 {object} model.Author
// @Router      /authors [get]
func (s *Server) GetAllAuthor(c *fiber.Ctx) error {
	author, err := s.store.GetAllAuthors()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(200).JSON(author)
}
