package model

type User struct {
	Id         int64  `json:"id,omitempty"`
	Uid        int64  `json:"uid,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Salt       string `json:"salt,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Age        int64  `json:"age,omitempty"`
	Sex        int64  `json:"sex,omitempty"`
	Face       string `json:"face,omitempty"`
	Experience int64  `json:"experience,omitempty"`
	Vip        int64  `json:"vip,omitempty"`
	Country    int64  `json:"country,omitempty"`
	Level      int64  `json:"level,omitempty"`

	Ctime int64 `json:"ctime"`
	Mtime int64 `json:"mtime"`
}

type UserAttr struct {
	Id            int64 `json:"id,omitempty"`
	Uid           int64 `json:"uid,omitempty"`
	AttackPower   int64 `json:"attack_power,omitempty"`
	Defense       int64 `json:"defense,omitempty"`
	Vitality      int64 `json:"vitality,omitempty"`
	Armor         int64 `json:"armor,omitempty"`
	ArmorPiercing int64 `json:"armor_piercing,omitempty"`
	Crit          int64 `json:"crit,omitempty"`

	Ctime int64 `json:"ctime"`
	Mtime int64 `json:"mtime"`
}

type Role struct {
	User     User     `json:"user"`
	UserAttr UserAttr `json:"user_attr"`
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
