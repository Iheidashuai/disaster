package dao

import (
	"context"
	"game/model"
	"github.com/golang/glog"
)

const (
	queryRoleMap         = "select id,role_id,city_id,can_go from role_map"
	queryRoleMapByRoleId = "select id,role_id,map_id,can_go from role_map where role_id = ?"

	updateCityStateByRoleId = "update role_map set can_go = ? where role_id = ? and map_id = ?"
)

type RoleMapDao interface {
	QueryRoleMap(ctx context.Context) ([]*model.RoleMap, error)
	QueryRoleMapByRoleId(ctx context.Context, id int64) (*model.RoleMap, error)
	UpdateCityStateByRoleId(ctx context.Context, roleId int64, mapId int64, isOpen bool) error
}

func (d *dao) QueryRoleMap(ctx context.Context) ([]*model.RoleMap, error) {
	result := make([]*model.RoleMap, 0)
	rows, err := d.db.Query(queryRoleMap)
	if err != nil {
		glog.Errorf("QueryRoleMap Query error(%s)", err)
	}
	for rows.Next() {
		r := &model.RoleMap{}
		if err = rows.Scan(&r.Id, &r.RoleId, &r.CityId, &r.CanGo); err != nil {
			glog.Errorf("QueryRoleMap Scan error(%v)", err)
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

func (d *dao) QueryRoleMapByRoleId(ctx context.Context, id int64) (*model.RoleMap, error) {
	r := &model.RoleMap{}
	row := d.db.QueryRow(queryRoleMapByRoleId, id)
	if row.Err() != nil {
		glog.Errorf("QueryRoleMapByRoleId QueryRow error(%s)", row.Err())
		return nil, row.Err()
	}
	if err := row.Scan(&r.Id, &r.RoleId, &r.CityId, &r.CanGo); err != nil {
		glog.Errorf("QueryRoleMapByRoleId Scan error(%s)", row.Err())
		return nil, err
	}
	return r, nil
}

func (d *dao) UpdateCityStateByRoleId(ctx context.Context, roleId int64, mapId int64, isOpen bool) error {
	var canGo int64
	if isOpen {
		canGo = 1
	}
	_, err := d.db.Exec(updateCityStateByRoleId, canGo, roleId, mapId)
	if err != nil {
		return err
	}
	return nil
}
