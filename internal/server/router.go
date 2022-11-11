package server

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
)

func (s *Server) NewRouter() {
	//Автор
	s.app.Get("/authors", s.GetAllAuthor)
	s.app.Get("/authors/:id", s.GetAuthorByIdAndBook)
	s.app.Post("/authors/", s.PostNewAuthor)
	s.app.Patch("/authors/:id", s.UpdateAuthor)
	s.app.Delete("authors/:id", s.DeleteAuthor)
	//Книги
	s.app.Get("/books", s.getAllBooks)
	s.app.Post("books", s.postNewBook)
	s.app.Patch("/books/:id", s.patchBook)
	s.app.Delete("/books/:id", s.DeleteBook)
	s.app.Get("/books/:id", s.GetBookById)
	//Участники
	s.app.Get("/members/", s.GetMembers)
	s.app.Post("/members/", s.PostNewMember)
	s.app.Patch("/members/:id", s.PatchNewMember)
	s.app.Delete("/members/:id", s.DeleteMember)
	s.app.Get("/members/:id", s.GetMemberById)
	//Специальные
	s.app.Get("/authors/:id/books", s.GetAuthorListIdBook)
	s.app.Get("/members/:id/books", s.GetMemberListIdBooks)
	s.app.Get("/*", s.AllOtherRequests)
	//swagger oauth
	s.app.Get("/swagger/*", swagger.HandlerDefault)
}
