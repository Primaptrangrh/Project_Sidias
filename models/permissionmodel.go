package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
)

type PermissionModel struct {
	db *sql.DB
}

func NewPermissionModel() *PermissionModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &PermissionModel{
		db: conn,
	}
}

func (u PermissionModel) Where(permission *entities.Permission, fieldName, fieldValue string) error {

	row, err := u.db.Query("select id, nama_lengkap, email, departement, position, reason from permission where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&permission.Id, &permission.NamaLengkap, &permission.Email, &permission.Departemen, &permission.Position, &permission.Reason)
	}

	return nil
}

func (u *PermissionModel) FindAllPermission() ([]entities.Permission, error) {

	rows, err := u.db.Query("select id, nama_lengkap, email, departement, position, reason from permission")
	if err != nil {
		return []entities.Permission{}, err
	}
	defer rows.Close()

	var dataPermission []entities.Permission
	for rows.Next() {
		var permission entities.Permission
		rows.Scan(&permission.Id,
			&permission.NamaLengkap,
			&permission.Email,
			&permission.Departemen,
			&permission.Position,
			&permission.Reason)
		dataPermission = append(dataPermission, permission)
	}
	return dataPermission, nil
	
}

func (u PermissionModel) CreatePermission(permission entities.Permission) bool {

	result, err := u.db.Exec("insert into permission (nama_lengkap, email, departement, position, reason) values(?,?,?,?,?)",
	permission.NamaLengkap, permission.Email, permission.Departemen, permission.Position, permission.Reason)

		if err != nil {
			fmt.Println(err)
			return false
		}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0

}