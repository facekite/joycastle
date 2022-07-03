// Package rand
/*
Copyright © 2022 隔壁写代码的王叔叔
*/
package rand

import (
	"time"
)

const BaseNumber = 13

// Rand13 获取 1-13中间的随机数
func Rand13() int {
	current := int(time.Now().UnixNano())
	return current%BaseNumber + 1
}

//GCD 求解两数 最大公约数
func GCD(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return GCD(y, tmp)
	} else {
		return y
	}
}

// LCM 求解两数最小公倍数
func LCM(x, y int) int {
	return x * y / GCD(x, y)
}

// RandN 利用最小公倍数求的最少需要求取随机数的次数;最少一次
func RandN(n int) int {
	if n < 2 {
		return n
	}
	result := 0
	lcm := LCM(n, BaseNumber)
	for i := 0; i < lcm/BaseNumber; i++ {
		rand := Rand13()
		result += rand - 1
	}
	result = result * n / lcm
	return result + 1
}

// Rand5 直接采用RandN 传入 5 求解
func Rand5() int {
	return RandN(5)
}
