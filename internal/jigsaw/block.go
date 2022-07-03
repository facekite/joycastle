// Package jigsaw
/*
Copyright © 2022 隔壁写代码的王叔叔
*/
package jigsaw

const (
	Concave = -1 // 凹
	Convex  = 1  // 凸
	Flat    = 0  // 平
)

// FlatBlock 一个四周都平的拼图块
var FlatBlock = &Block{
	Left:   Flat,
	Right:  Flat,
	Top:    Flat,
	Bottom: Flat,
}

// Block 定义一个拼图块,设定左右上下四边的凹凸平状态
type Block struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}
