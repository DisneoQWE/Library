package server

import (
	_ "RestApiLibrary/docs"
	"RestApiLibrary/internal/model"
	store2 "RestApiLibrary/internal/store"
	"RestApiLibrary/pkg/config"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"log"
)

type Store interface {
	//Авторы
	GetAllAuthors() ([]model.Author, error)
	GetAllAuthorsById(authorId int) (*model.Author, error)
	PostNewAuthor(*model.Author) error
	UpdateAuthor(*model.Author, int) error
	DeleteAuthor(authorId int) error
	//Книги
	GetAllBooks() ([]model.Book, error)
	PostNewBook(*model.Book) error
	PatchNewBook(*model.Book, int) error
	DeleteBook(int) error
	GetBookById(int) (*model.Book, error)
	//Читатели
	GetMembers() ([]model.Member, error)
	PostNewMember(*model.Member) error
	PatchNewMember(*model.Member, int) error
	DeleteMember(int) error
	CreateZeroMember() error
	GetMemberById(int) (*model.Member, error)
	//Специальные пути
	GetAuthorListIdBook(authorId int) ([]store2.Result, error)
	GetMemberListIDBook(memberId int) ([]store2.StoreResult, error)
}

type Server struct {
	app   *fiber.App
	store Store
}

func NewServer(db *sqlx.DB) *Server {
	app := fiber.New()
	store := store2.NewStore(db)
	return &Server{
		app:   app,
		store: store,
	}
}

// Start Server
func (s *Server) ServerRun(c *config.Config) {
	s.app.Get("/swagger/*", swagger.HandlerDefault)
	s.app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	s.NewRouter()
	err := s.app.Listen(":" + c.Port)
	if err != nil {
		log.Fatal(err)
	}
}
