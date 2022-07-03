// Package jigsaw
/*
Copyright © 2022 隔壁写代码的王叔叔
*/
package jigsaw

// Box 一个Box可以摆放 width * height 块拼图，当fillCount为 width*height 表示拼图已满
type Box struct {
	width     uint
	height    uint
	table     [][]*Block
	fillCount uint
}

// Position 拼图配置，主要抽象位置
type Position struct {
	X uint
	Y uint
}

// IsFinished 是否完成，填充的区块数量等于最多可以填充的块数
func (box *Box) IsFinished() bool {
	return box.fillCount == box.width*box.height
}

// AddBlock 添加拼图到指定位置，如果指定位置已经存在返回已存在，如果无法匹配返回错误
func (box *Box) AddBlock(block *Block, pos Position) error {
	if block == nil {
		return BlockIsNil
	}
	if box.table[pos.X][pos.Y] != nil {
		return AlreadyExists
	}
	// 默认四周为平
	var left, right, top, bottom = FlatBlock, FlatBlock, FlatBlock, FlatBlock
	if pos.X > 0 {
		left = box.table[pos.X-1][pos.Y]
	}
	if pos.X < box.width-1 {
		right = box.table[pos.X+1][pos.Y]
	}
	if pos.Y > 0 {
		top = box.table[pos.X][pos.Y-1]
	}
	if pos.Y < box.height-1 {
		bottom = box.table[pos.X][pos.Y+1]
	}

	if left != nil && left.Right+block.Left != Flat {
		return CannotMatch
	}

	if top != nil && top.Bottom+block.Top != Flat {
		return CannotMatch
	}

	if right != nil && right.Left+block.Right != Flat {
		return CannotMatch
	}

	if bottom != nil && bottom.Top+block.Bottom != Flat {
		return CannotMatch
	}
	box.table[pos.X][pos.Y] = block
	box.fillCount += 1
	return nil
}

// NewBox 创建一个拼图BOX,初始化每个位置的拼图为空
func NewBox(width, height uint) *Box {
	return &Box{
		width:     width,
		height:    height,
		fillCount: width * height,
		table:     make([][]*Block, width*height, width*height),
	}
}
