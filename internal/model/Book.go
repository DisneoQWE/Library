package model

import (
	"database/sql"
)

//swagger:model
type Book struct {
	//BookId id of the book
	//required: true
	BookId int `json:"book_id" db:"bookid" params:"book_id"`
	//BookName
	//required true
	BookName string `json:"book_name" db:"bookname" params:"book_name"`
	//BookGenre
	//required: false
	BookGenre string `json:"book_genre" db:"bookgenre" params:"book_genre"`
	//IsbnCode
	//required: false
	IsbnCode int `json:"isbn_code" db:"bookisbn" params:"isbn_code"`
	//required:true
	AuthorId int `json:"author_id" db:"author_id" params:"author_id"`
	//required: false
	MemberID sql.NullInt64 `json:"member_id" db:"memberid" params:"member_id"`
}

type BookResult struct {
	BookId    int    `json:"book_id" db:"bookid" params:"book_id"`
	BookName  string `json:"book_name" db:"bookname" params:"book_name"`
	BookGenre string `json:"book_genre" db:"bookgenre" params:"book_genre"`
	IsbnCode  int    `json:"isbn_code" db:"bookisbn" params:"isbn_code"`
	AuthorId  int    `json:"author_id" db:"author_id" params:"author_id"`
	MemberID  int    `json:"member_id" db:"memberid" params:"member_id"`
}

func ConvertNullInt(index int) sql.NullInt64 {
	var sql sql.NullInt64
	sql.Int64 = int64(index)
	if index == 0 {
		sql.Valid = false
	} else {
		sql.Valid = true
	}
	return sql
}

func ConvertFromUllIntToInt(nullInt64 sql.NullInt64) int {
	index := int(nullInt64.Int64)
	return index
}
