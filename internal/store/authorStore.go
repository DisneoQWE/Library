package store

import (
	"RestApiLibrary/internal/model"
)

// GET: Возвращает список всех авторов
func (a *DBConnection) GetAllAuthors() ([]model.Author, error) {
	authors := make([]model.Author, 0, 5) //make create a heap in memory for a slice maps and initializes and puts zero or empty
	rows, err := a.db.NamedQuery("SELECT * FROM authors", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		author := model.Author{}
		err = rows.Scan(&author.AuthorId, &author.AuthorFio, &author.AuthorPseudonym, &author.AuthorSpecialization)

		authors = append(authors, author)
	}
	if err != nil {
		return nil, err
	}
	return authors, nil
}

// GET: Показывает автора по ID
func (a *DBConnection) GetAllAuthorsById(authorId int) (*model.Author, error) {
	author := new(model.Author)
	rows, err := a.db.NamedQuery(`SELECT * FROM authors WHERE authorid=:authorId`, map[string]interface{}{
		"authorId": authorId})
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&author.AuthorId, &author.AuthorFio, &author.AuthorPseudonym, &author.AuthorSpecialization)
	}
	if err != nil {
		return nil, err
	}
	return author, nil
}

// POST: Создает нового автора
func (a *DBConnection) PostNewAuthor(author *model.Author) error {
	_, err := a.db.NamedExec(`INSERT INTO authors (authorfio, authorpseudonym, authorspecialization)
		VALUES (:authorfio,:authorpseudonym,:authorspecialization)`,
		map[string]interface{}{
			"authorfio":            author.AuthorFio,
			"authorpseudonym":      author.AuthorPseudonym,
			"authorspecialization": author.AuthorSpecialization})
	if err != nil {
		return err
	}

	return nil
}

// PATCH: Обновляет сущность автора
func (a *DBConnection) UpdateAuthor(author *model.Author, authorId int) error {
	var err error
	if author.AuthorFio != "" {
		_, err = a.db.NamedExec(`UPDATE authors SET authorfio=:authorfio WHERE authorid=:authorid`, map[string]interface{}{
			"authorid":  authorId,
			"authorfio": author.AuthorFio})
	}
	if author.AuthorSpecialization != "" {
		_, err = a.db.NamedExec(`UPDATE authors SET authorpseudonym=:authorpseudonym WHERE authorid=:authorid`, map[string]interface{}{
			"authorid":        authorId,
			"authorpseudonym": author.AuthorPseudonym})
	}
	if author.AuthorSpecialization != "" {
		_, err = a.db.NamedExec(`UPDATE authors SET authorspecialization=:authorspecialization WHERE authorid=:authorid`, map[string]interface{}{
			"authorid":             authorId,
			"authorspecialization": author.AuthorSpecialization})
	}
	if err != nil {
		return err
	}
	return nil
}

// DELETE: Удаляет ряд в таблице Author
func (a *DBConnection) DeleteAuthor(authorId int) error {
	_, err := a.db.NamedQuery(`UPDATE books SET author_id=null WHERE author_id=:authorId`, map[string]interface{}{
		"authorId": authorId,
	})
	_, err = a.db.NamedQuery(`DELETE FROM authors WHERE authorid=:authorid`, map[string]interface{}{
		"authorid": authorId,
	})
	if err != nil {
		return err
	}
	return nil
}
