package dao

import (
	"context"
	"errors"
	"fmt"
	"game/model"
	"game/utils"
	"github.com/golang/glog"
)

const (
	queryRole      = "select id,account,username,password from role where role_id = ?"
	updatePassword = "update role set password = ? where role_id = ?"
	registerRole   = "insert into role(username,password) values(?,?,?)"
)

type RoleDao interface {
	GetRole(ctx context.Context, roleId int64) (*model.Role, error)
	Login(ctx context.Context, roleId, password int64) error
	Register(ctx context.Context, role *model.Role) (int64, error)
	UpdatePassword(ctx context.Context, roleId int64, oldPassword, newPassword int64) error
}

func (d *dao) GetRole(ctx context.Context, roleId int64) (*model.Role, error) {
	r := &model.Role{}
	row := d.db.QueryRow(queryRole, roleId)
	if row.Err() != nil {
		if err := row.Scan(&r.Id, &r.Account, &r.Username, &r.Password); err != nil {
			glog.Errorf("GetRole QueryRow error(%s)", err)
			return nil, err
		}
	}
	return r, nil
}

func (d *dao) Login(ctx context.Context, roleId int64, password string) error {
	r, err := d.GetRole(ctx, roleId)
	if err != nil {
		glog.Errorf("Login error(%s)", err)
		return err
	}

	salt := r.Salt
	if utils.Encrypt(password+salt) == r.Password {
		return nil
	}
	return errors.New(fmt.Sprintf("role (%d) password error", roleId))
}

func (d *dao) Register(ctx context.Context, role *model.Role) (int64, error) {
	d.db.Exec(registerRole, role.Username, role.Password)
	return 0, nil
}

func (d *dao) UpdatePassword(ctx context.Context, roleId int64, oldPassword, newPassword int64) error {
	panic("implement me")
}
