// Package jigsaw
/*
Copyright © 2022 隔壁写代码的王叔叔
*/
package jigsaw

import "errors"

var (
	AlreadyExists = errors.New("already exists")
	CannotMatch   = errors.New("cannot match")
	BlockIsNil    = errors.New("block is nil")
)
