/*
* @Author: pzqu
* @Date:   2020-03-19 00:19
*/
package _09_Longest_Palindrome

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tables := []struct {
		x string
		y int
	}{
		{"abccccdd", 7},
		{"z", 1},
	}

	for _, table := range tables {
		y := longestPalindrome(table.x)
		if y != table.y {
			t.Fatalf("error : %v %v,but %v", table.x, table.y, y)
		}
	}
}
