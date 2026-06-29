package utils

import (
	"math/rand"
)

// Возвращает случайное число от 0 до n, не паникует
func RandomFrom0ToN(n int) int {
	if n <= 0 {
		return 0
	}

	return rand.Intn(n)
}

// Возвращает случайное число от 1 до n, не паникует
func RandomFrom1ToN(n uint) int {
	if n <= 1 {
		return 1
	}

	return rand.Intn(int(n)) + 1
}

// Возвращает одно случайное число из набора предложенных
func ChooseRandom(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[rand.Intn(len(nums))]
}
