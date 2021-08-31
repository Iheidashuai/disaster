package model

type BattleShip struct {
	Id    int64 `json:"id"`
	Level int64 `json:"level"`

	AttackPower int64 `json:"attack_power"`
	Defense     int64 `json:"defense"`
	Life        int64 `json:"life"`
}
