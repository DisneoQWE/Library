package model

//swagger:model
type Member struct {
	//required: true
	MemberId int `json:"memberid" db:"memberid" params:"memberid"`
	//required: true
	MemberFio string `json:"memberfio" db:"memberfio" params:"memberfio"`
}
