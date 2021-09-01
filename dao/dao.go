package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Dao interface {
	RoleDao
	RoleAttrDao
	RoleMapDao
}

type dao struct {
	db *sql.DB
}

func NewDao() *dao {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.106:3306)/disaster")
	if err != nil {
		panic(err)
	}
	return &dao{db: db}
}
