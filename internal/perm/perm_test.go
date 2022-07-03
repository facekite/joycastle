// Package perm
/*
Copyright © 2022 隔壁写代码的王叔叔
*/
package perm

import (
	"sort"
	"testing"
)
import "github.com/go-playground/assert/v2"

func TestPerm(t *testing.T) {
	t.Run("Test Empty Array", func(t *testing.T) {
		empty := []rune("")
		var got []string
		Perm(empty, func(array []rune) {
			got = append(got, string(array))
		})
		var want []string
		assert.Equal(t, got, want)
	})

	t.Run("Test ABC", func(t *testing.T) {
		input := []rune("ABC")
		var got []string
		Perm(input, func(array []rune) {
			got = append(got, string(array))
		})
		want := []string{
			"ABC",
			"ACB",
			"BAC",
			"BCA",
			"CAB",
			"CBA",
		}
		sort.Strings(got)
		sort.Strings(want)
		assert.Equal(t, got, want)
	})

	t.Run("Test ABC Any Order", func(t *testing.T) {
		inputs := []string{
			"ABC",
			"ACB",
			"BAC",
			"BCA",
			"CAB",
			"CBA",
		}
		for _, input := range inputs {
			var got []string
			Perm([]rune(input), func(array []rune) {
				got = append(got, string(array))
			})
			want := []string{
				"ABC",
				"ACB",
				"BAC",
				"BCA",
				"CAB",
				"CBA",
			}
			sort.Strings(got)
			sort.Strings(want)
			assert.Equal(t, got, want)
		}

	})

}
