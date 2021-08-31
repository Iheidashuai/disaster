package model

type MilitaryRank struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	AttackPower int64  `json:"attack_power"`
	Defense     int64  `json:"defense"`
	Life        int64  `json:"life"`

	Soldiers              []int64 `json:"soldiers"`
	JuniorSergeants       []int64 `json:"junior_sergeants"`
	IntermediateSergeants []int64 `json:"intermediate_sergeants"`
	AdvancedSergeants     []int64 `json:"advanced_sergeants"`

	JuniorOfficers       []int64 `json:"junior_officers"`
	IntermediateOfficers []int64 `json:"intermediate_officers"`
	AdvancedOfficers     []int64 `json:"advanced_officers"`
}
