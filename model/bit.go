package model

func Add(num1, num2 uint64) uint64 {
	if num2 == 0 {
		return num1
	}
	sum := num1 ^ num2
	carry := (num1 & num2) << 1
	return Add(sum, carry)
}

func Sub(num1, num2 uint64) uint64 {
	subtract := Add(^num2, 1)     // 先求减数的补码（取反加一）
	result := Add(num1, subtract) // add()即上述加法运算
	return result
}
