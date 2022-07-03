// Package perm
/*
Copyright © 2022 隔壁写代码的王叔叔
*/
package perm

// Perm 全排列输入的字符串，无需按照特定顺序
func Perm(array []rune, f func([]rune)) {
	if len(array) <= 0 {
		return
	}
	perm(array, f, 0)
}

// perm
func perm(array []rune, f func([]rune), i int) {
	if i == len(array) {
		f(array)
	}
	for j := i; j < len(array); j++ {
		array[i], array[j] = array[j], array[i]
		perm(array, f, i+1)
		array[i], array[j] = array[j], array[i]
	}
}
