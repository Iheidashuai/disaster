package dao

import (
	"context"
	"database/sql"
	"game/model"
	"github.com/golang/glog"
)

const (
	queryRoleById          = "select id,attack_power,defense,life,hat,clothing,shoes,weapon,pants,level,military_rank from role where id = ?"
	updateRoleBaseAttr     = "update role set attack = ?,defense = ?,life = ? where id = ?"
	updateRoleEquipment    = "update role set hat = ?,clothing = ?,shoes = ?,weapon = ?,pants = ?,role = ? where id = ?"
	updateRoleLevel        = "update role set level = ? where id = ?"
	updateRoleMilitaryRank = "update role set military_rank = ? where id = ?"
	addRole                = "insert into role(attack_power,defense,life) values(?,?,?)"
)

type RoleDao interface {
	QueryRoleById(ctx context.Context,id int64) (*model.Role,error)
	UpdateRoleBaseAttr(ctx context.Context,role *model.Role,id int64) error
	UpdateRoleEquipment(ctx context.Context,role *model.Role,id int64) error
	UpdateRoleLevel(ctx context.Context,level int64,id int64) error
	UpdateRoleMilitaryRank(ctx context.Context,level int64,id int64) error
	AddRole(ctx context.Context,role *model.Role) (int64,error)
}

func (d *dao) QueryRoleById(ctx context.Context,id int64) (*model.Role,error) {
	var (
		row *sql.Row
		err error
	)
	role := &model.Role{}
	if row = d.db.QueryRow(queryRoleById,id);row.Err() == nil {
		if err = row.Scan(&role.Id,&role.AttackPower,&role.Defense,&role.Life,&role.Hat,
			&role.Clothing,&role.Shoes,&role.Weapon,&role.Pants,&role.Level,&role.MilitaryRank);err != nil {
			glog.Errorf("QueryRoleById error(%s)",err)
			return nil, err
		}
	} else {
		glog.Errorf("QueryRoleById error(%s)",row.Err())
		return nil,row.Err()
	}
	return role, nil
}

func (d *dao) UpdateRoleBaseAttr(ctx context.Context,role *model.Role,id int64) error {
	_,err := d.db.Exec(updateRoleBaseAttr,&role.AttackPower,&role.Defense,&role.Life,id)
	if err != nil {
		glog.Errorf("UpdateRoleBaseAttr error(%s)",err)
		return err
	}
	return nil
}

func (d *dao) UpdateRoleEquipment(ctx context.Context,role *model.Role,id int64) error {
	_,err := d.db.Exec(updateRoleEquipment,&role.Hat,&role.Clothing,&role.Shoes,&role.Weapon,&role.Pants,id)
	if err != nil {
		glog.Errorf("UpdateRoleEquipment error(%s)",err)
		return err
	}
	return nil
}

func (d *dao) UpdateRoleLevel(ctx context.Context,level int64,id int64) error {
	_,err := d.db.Exec(updateRoleLevel,level,id)
	if err != nil {
		glog.Errorf("UpdateRoleLevel error(%s)",err)
		return err
	}
	return nil
}

func (d *dao) UpdateRoleMilitaryRank(ctx context.Context,level int64,id int64) error {
	_,err := d.db.Exec(updateRoleMilitaryRank,level,id)
	if err != nil {
		glog.Errorf("UpdateRoleMilitaryRank error(%s)",err)
		return err
	}
	return nil
}

func (d *dao) AddRole(ctx context.Context,role *model.Role) (int64,error) {
	r,err := d.db.Exec(addRole,&role.AttackPower,&role.Defense,&role.Life)
	if err != nil {
		glog.Errorf("UpdateRoleMilitaryRank error(%s)",err)
		return 0,err
	}
	return r.LastInsertId()
}