/*
* @Author: pzqu
* @Date:   2020-03-16 23:30
*/
package String

import (
	"testing"
)

//Test
func TestCompressString(t *testing.T) {

	tables := []struct {
		x string
		y string
	}{
		{
			"aabcccccaaa", "a2b1c5a3",
		},
		{
			"aabcccccaa", "a2b1c5a2",
		},
		{"", "",},
		{"abbccd", "abbccd",},
	}

	for _, table := range tables {
		y := compressString(table.x)
		if y != table.y {
			t.Fatalf("error x:%s,y:%s", table.x, table.y)
		}
	}

}
