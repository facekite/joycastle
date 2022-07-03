// Package rand
/*
Copyright © 2022 隔壁写代码的王叔叔
*/
package rand

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestRand5(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := RandN(5)
		t.Logf("rand number %v", v)
		assert.Equal(t, v <= 5 && v >= 1, true)
	}
}
