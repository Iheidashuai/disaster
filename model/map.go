package model

import (
	"fmt"
)

type RoleMap struct {
	Id     int64 `json:"id"`
	RoleId int64 `json:"role_id"`
	CityId int64 `json:"city_id"`
	CanGo  int64 `json:"can_go"`
}

var CityMap = map[int64]string{
	1:  "沙朵镇",
	2:  "白银港",
	3:  "黑暗森林",
	4:  "无尽之海",
	5:  "海歌城",
	6:  "陨星城",
	7:  "云翼村",
	8:  "枫林城",
	9:  "荆棘镇",
	10: "北境草原",
}

type City struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	IsOpen bool   `json:"is_open"`
}

type Map struct {
	Cities []*City
	Edges  map[City][]*City
}

func (c *City) Open() {
	c.IsOpen = true
}

func (c *City) String() string {
	return fmt.Sprintf("%v", c.Name)
}

func (m *Map) Go(srcCity, desCity *City) bool {
	for _, nearCity := range m.Edges[*srcCity] {
		if nearCity.String() == desCity.String() && desCity.IsOpen {
			return true
		}
	}
	return false
}

func (m *Map) Init(globalMap map[City][]*City) {
	m.Edges = globalMap
	for city := range globalMap {
		m.Cities = append(m.Cities, &city)
	}
}
