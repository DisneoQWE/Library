package model

// swagger:model
type Author struct {
	//Author ID of author
	//required: true
	AuthorId int `json:"author_id" db:"authorid" params:"author_id" binding:""`
	//Author Fio
	//required: true
	AuthorFio string `json:"author_fio" db:"authorfio" params:"author_id"`
	//Author AuthorPseudonym
	//required: false
	AuthorPseudonym string `json:"author_pseudonym" db:"authorpseudonym" params:"author_pseudonym"`
	//Author AuthorSpecialization
	//required: false
	AuthorSpecialization string `json:"author_specialization" db:"authorspecialization" params:"author_specialization"`
}
