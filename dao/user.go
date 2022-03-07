package dao

import (
	"context"
	"database/sql"
	"fmt"
	"game/common"
	"game/model"
	"github.com/golang/glog"
)

const (
	queryUserByUid          = "select id, uid, username, password, salt, email, phone, age, sex, face, experience, vip, country, level from user where uid = ?"
	addUser = "insert into user(uid,username,password,salt,email,phone,age,sex,face) values(?,?,?,?,?,?,?,?,?) "
)

type RoleDao interface {
	QueryUserByUid(ctx context.Context,uid int64) (*model.User,error)
	Login(ctx context.Context,user *model.User) error
	Register(ctx context.Context,user *model.User) error
}

func (d *dao) QueryUserByUid(ctx context.Context,uid int64) (*model.User,error) {
	var (
		row *sql.Row
		err error
	)
	role := &model.User{}
	if row = d.db.QueryRow(queryUserByUid,uid);row.Err() == nil {
		if err = row.Scan(&role.Id,&role.Uid,&role.Username,&role.Password,&role.Salt,
			&role.Email,&role.Phone,&role.Age,&role.Sex,&role.Face,&role.Experience,&role.Vip,&role.Country,&role.Level);err != nil {
			glog.Errorf("QueryRoleById error(%s)",err)
			return nil, err
		}
	} else {
		glog.Errorf("QueryRoleById error(%s)",row.Err())
		return nil,row.Err()
	}
	return role, nil
}

func (d *dao) Login(ctx context.Context,user *model.User) error {
	if user.Uid <= 0 {
		return common.Ecode(401,"")
	}

	u,err := d.QueryUserByUid(ctx,user.Uid)
	if err != nil {
		return err
	}

	if d.md5.encrypt(fmt.Sprintf("%s%s",user.Password,u.Salt)) != u.Password {
		return common.Ecode(401,"")
	}

	return nil
}

func (d *dao) Register(ctx context.Context,user *model.User) error {
	if user.Uid <= 0 || user.Password == "" {
		return common.Ecode(401,"")
	}

	salt := d.md5.randomString(10)
	_,err := d.db.Exec(addUser,user.Uid,user.Username,d.md5.encrypt(fmt.Sprintf("%s%s",user.Password,salt)),salt,user.Email,user.Phone,user.Age,user.Sex,user.Face)
	if err != nil {
		return err
	}
	return nil
}

