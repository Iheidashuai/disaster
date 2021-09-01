package dao

import (
	"context"
	"database/sql"
	"game/model"
	"github.com/golang/glog"
)

const (
	queryRoleById          = "select id,role_id,attack_power,defense,life,hat,clothing,shoes,weapon,pants,level,military_rank from role_attr where role_id = ?"
	updateRoleBaseAttr     = "update role_attr set attack = ?,defense = ?,life = ? where role_id = ?"
	updateRoleEquipment    = "update role_attr set hat = ?,clothing = ?,shoes = ?,weapon = ?,pants = ?,role = ? where role_id = ?"
	updateRoleLevel        = "update role_attr set level = ? where role_id = ?"
	updateRoleMilitaryRank = "update role_attr set military_rank = ? where role_id = ?"
	addRoleAttr            = "insert role_attr role(role_id,attack_power,defense,life) values(?,?,?)"
)

type RoleAttrDao interface {
	QueryRoleAttrById(ctx context.Context, id int64) (*model.RoleAttr, error)
	UpdateRoleAttrBaseAttr(ctx context.Context, role *model.RoleAttr, id int64) error
	UpdateRoleAttrEquipment(ctx context.Context, role *model.RoleAttr, id int64) error
	UpdateRoleAttrLevel(ctx context.Context, level int64, id int64) error
	UpdateRoleAttrMilitaryRank(ctx context.Context, level int64, id int64) error
	AddRoleAttr(ctx context.Context, role *model.RoleAttr) (int64, error)
}

func (d *dao) QueryRoleAttrById(ctx context.Context, id int64) (*model.RoleAttr, error) {
	var (
		row *sql.Row
		err error
	)
	role := &model.RoleAttr{}
	if row = d.db.QueryRow(queryRoleById, id); row.Err() == nil {
		if err = row.Scan(&role.Id, &role.RoleId, &role.AttackPower, &role.Defense, &role.Life, &role.Hat,
			&role.Clothing, &role.Shoes, &role.Weapon, &role.Pants, &role.Level, &role.MilitaryRank); err != nil {
			glog.Errorf("QueryRoleById error(%s)", err)
			return nil, err
		}
	} else {
		glog.Errorf("QueryRoleById error(%s)", row.Err())
		return nil, row.Err()
	}
	return role, nil
}

func (d *dao) UpdateRoleAttrBaseAttr(ctx context.Context, role *model.RoleAttr, id int64) error {
	_, err := d.db.Exec(updateRoleBaseAttr, &role.AttackPower, &role.Defense, &role.Life, id)
	if err != nil {
		glog.Errorf("UpdateRoleBaseAttr error(%s)", err)
		return err
	}
	return nil
}

func (d *dao) UpdateRoleAttrEquipment(ctx context.Context, role *model.RoleAttr, id int64) error {
	_, err := d.db.Exec(updateRoleEquipment, &role.Hat, &role.Clothing, &role.Shoes, &role.Weapon, &role.Pants, id)
	if err != nil {
		glog.Errorf("UpdateRoleEquipment error(%s)", err)
		return err
	}
	return nil
}

func (d *dao) UpdateRoleAttrLevel(ctx context.Context, level int64, id int64) error {
	_, err := d.db.Exec(updateRoleLevel, level, id)
	if err != nil {
		glog.Errorf("UpdateRoleLevel error(%s)", err)
		return err
	}
	return nil
}

func (d *dao) UpdateRoleAttrMilitaryRank(ctx context.Context, level int64, id int64) error {
	_, err := d.db.Exec(updateRoleMilitaryRank, level, id)
	if err != nil {
		glog.Errorf("UpdateRoleMilitaryRank error(%s)", err)
		return err
	}
	return nil
}

func (d *dao) AddRoleAttr(ctx context.Context, role *model.RoleAttr) (int64, error) {
	r, err := d.db.Exec(addRoleAttr, &role.AttackPower, &role.Defense, &role.Life)
	if err != nil {
		glog.Errorf("UpdateRoleMilitaryRank error(%s)", err)
		return 0, err
	}
	return r.LastInsertId()
}
