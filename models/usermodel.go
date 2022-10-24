package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("select id, nama_lengkap, email, username, password, role from users where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Username, &user.Password, &user.Role)
	}

	return nil
}

func (u UserModel) Create(user entities.User) bool {

	result, err := u.db.Exec("insert into users (nama_lengkap, email, username, password, role) values(?,?,?,?,?)",
		user.NamaLengkap, user.Email, user.Username, user.Password, user.Role)

		if err != nil {
			fmt.Println(err)
			return false
		}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0

}

func (u *UserModel) FindAll() ([]entities.User, error) {

	rows, err := u.db.Query("select id, nama_lengkap, email, username, role from users")
	if err != nil {
		return []entities.User{}, err
	}
	defer rows.Close()

	var dataUser []entities.User
	for rows.Next() {
		var user entities.User
		rows.Scan(&user.Id,
			&user.NamaLengkap,
			&user.Email,
			&user.Username,
			&user.Role)
		dataUser = append(dataUser, user)
	}
	return dataUser, nil
	
}

func (p *UserModel) Find(id int64, user *entities.User) error {

	return p.db.QueryRow("select * from users where id = ?", id).Scan(
		&user.Id,
		&user.NamaLengkap,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role)
}

func (p *UserModel) Update(user entities.User) error {

	_, err := p.db.Exec(
		"update users set nama_lengkap = ?, email = ?, username = ?, password = ?, role = ? where id = ?",
		user.NamaLengkap, user.Email, user.Username, user.Password, user.Role, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *UserModel) Delete(id int64) {
	p.db.Exec("delete from users where id = ?", id)
}


