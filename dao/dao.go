package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type dao struct {
	db *sql.DB
	md5 *MD5
}

func NewDao() *dao{
	db,err := sql.Open("mysql", "root:123456@tcp(101.43.3.32:4513)/disaster")
	if err != nil {
		panic(err)
	}
	return &dao{db: db,md5: &MD5{}}
}