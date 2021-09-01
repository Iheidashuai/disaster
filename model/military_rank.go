package model

type MilitaryRank struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	AttackPower int64  `json:"attack_power"`
	Defense     int64  `json:"defense"`
	Life        int64  `json:"life"`
}
