package model

type Role struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Level        uint64  `json:"level"`
	MilitaryRank uint64  `json:"military_rank"`

	AttackPower uint64 `json:"attack_power"`
	Defense     uint64 `json:"defense"`
	Life        uint64 `json:"life"`

	Hat      uint64 `json:"hat"`
	Clothing uint64 `json:"clothing"`
	Pants    uint64 `json:"pants"`
	Shoes    uint64 `json:"shoes"`
	Weapon   uint64 `json:"weapon"`

	Ctime uint8 `json:"ctime"`
	Mtime uint8 `json:"mtime"`
}

func ParseEquipment(equipment uint64) (uint64, uint64) {
	// 装备高 8 位代表装备 id，低 8 位代表强化等级
	equipmentId := equipment >> 8
	equipmentLevel := equipment & 0xff
	return equipmentId, equipmentLevel
}

func FormatEquipment(equipmentId, equipmentLevel uint64) uint64 {
	return equipmentId<<8 + equipmentLevel
}

// Strengthen 强化装备
func Strengthen(equipment uint64, success bool) uint64 {
	_, equipmentLevel := ParseEquipment(equipment)
	if success {
		return strengthenSuccess(equipmentLevel)
	}
	return strengthenFailed(equipmentLevel)
}

func strengthenSuccess(equipmentLevel uint64) uint64 {
	return Add(equipmentLevel, 1)
}

func strengthenFailed(equipmentLevel uint64) uint64 {
	return Sub(equipmentLevel, 1)
}