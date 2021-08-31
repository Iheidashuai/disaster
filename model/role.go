package model

import "time"

type Role struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Level        int64  `json:"level"`
	MilitaryRank int64  `json:"military_rank"`

	AttackPower int64 `json:"attack_power"`
	Defense     int64 `json:"defense"`
	Life        int64 `json:"life"`

	Hat      int64 `json:"hat"`
	Clothing int64 `json:"clothing"`
	Pants    int64 `json:"pants"`
	Shoes    int64 `json:"shoes"`
	Weapon   int64 `json:"weapon"`

	Ctime time.Time `json:"ctime"`
	Mtime time.Time `json:"mtime"`
}
