package store

import (
	"RestApiLibrary/internal/model"
	"log"
)

func (a *DBConnection) GetMembers() ([]model.Member, error) {
	members := make([]model.Member, 0, 5) //make create a heap in memory for a slice maps and initializes and puts zero or empty
	rows, err := a.db.NamedQuery("SELECT * FROM members", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		member := model.Member{}
		err := rows.Scan(&member.MemberId, &member.MemberFio)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func (a *DBConnection) PostNewMember(member *model.Member) error {
	a.CreateZeroMember()
	_, err := a.db.NamedExec(`INSERT INTO members (memberid, memberfio)
	VALUES (DEFAULT,:memberfio)`, map[string]interface{}{
		"memberfio": member.MemberFio})
	if err != nil {
		return err
	}
	return nil
}

func (a *DBConnection) CreateZeroMember() error {
	members, err := a.GetMembers()
	flag := false
	for index := range members {
		if members[index].MemberId != 0 {
			flag = true
		}
	}
	if flag {
		_, err = a.db.NamedQuery(`INSERT INTO members(memberid, memberfio) values (:memberid, :memberfio)`,
			map[string]interface{}{
				"memberid":  0,
				"memberfio": "Никто"})
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *DBConnection) PatchNewMember(member *model.Member, memberid int) error {
	a.CreateZeroMember()
	if memberid != 0 {
		member.MemberId = memberid
		_, err := a.db.NamedQuery(`UPDATE members SET memberfio=:memberfio WHERE memberid=:memberid`, map[string]interface{}{
			"memberid":  member.MemberId,
			"memberfio": member.MemberFio})
		if err != nil {
			log.Println("DB ERROR:", err.Error())
			return err
		}
	}
	return nil
}

func (a *DBConnection) DeleteMember(memberid int) error {
	a.CreateZeroMember()
	if memberid != 0 {
		_, err := a.db.NamedQuery(`UPDATE books SET memberid=null WHERE memberid=:memberid`, map[string]interface{}{
			"memberid": memberid,
		})
		if err != nil {
			return err
		}
		_, err = a.db.NamedQuery(`DELETE FROM members WHERE memberid=:memberid`, map[string]interface{}{
			"memberid": memberid,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *DBConnection) GetMemberById(memberId int) (*model.Member, error) {
	member := new(model.Member)
	rows, err := a.db.NamedQuery(`SELECT * FROM members WHERE memberid=:memberid`, map[string]interface{}{
		"memberid": memberId})
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&member.MemberId, &member.MemberFio)
	}
	if err != nil {
		return nil, err
	}
	return member, nil
}
