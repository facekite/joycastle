// Package perm
/*
Copyright © 2022 隔壁写代码的王叔叔
*/
package main

import (
	"fmt"
	"github.com/facekite/joycastle/internal/perm"
)

func main() {
	perm.Perm([]rune("ABC"), func(a []rune) {
		fmt.Println(string(a))
	})
}
