package store

type Result struct {
	Author     string `json:"author" db:"author" params:"author"`
	Books_name string `json:"books_name" db:"books_name" params:"books_name"`
}
type StoreResult struct {
	MemberFio string `json:"memberFio" db:"memberFio" params:"memberFio"`
	Book_name string `json:"book_Name" db:"book_Name" params:"book_Name"`
}

func (a *DBConnection) GetAuthorListIdBook(authorId int) ([]Result, error) {
	result := make([]Result, 0, 5)
	rows, err := a.db.NamedQuery(`SELECT authors.authorfio as author, books.bookname as books_name FROM authors
	JOIN books ON authors.authorid = books.author_id where books.author_id =:authorId
	group by authors.authorfio, books.bookname;`, map[string]interface{}{
		"authorId": authorId,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		res := Result{}
		err := rows.Scan(&res.Author, &res.Books_name)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}

func (a *DBConnection) GetMemberListIDBook(memberId int) ([]StoreResult, error) {
	storeResult := make([]StoreResult, 0, 5)
	rows, err := a.db.NamedQuery(`SELECT members.memberfio as memberfio, books. bookname as book_name FROM members
	JOIN books on members.memberid = books.memberid WHERE books.memberid=:memberId
	group by members.memberfio, books. bookname;`, map[string]interface{}{
		"memberId": memberId,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		res := StoreResult{}
		err := rows.Scan(&res.MemberFio, &res.Book_name)
		if err != nil {
			return nil, err
		}
		storeResult = append(storeResult, res)
	}
	return storeResult, nil
}
