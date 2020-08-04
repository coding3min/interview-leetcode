/*
* @Author: pzqu
* @Date:   2020-03-18 01:23
*/
package _160_Find_Words_That_Can_Be_Formed_by_Characters

import (
	"strings"
	"testing"
)

var tables []struct {
	x []string
	y string
	z int
}

func init() {
	tables = []struct {
		x []string
		y string
		z int
	}{
		{[]string{"cat", "bt", "hat", "tree"}, "atach", 6},
		{[]string{"hello", "world", "leetcode"}, "welldonehoneyr", 10},
	}
}

//Test
func TestCountCharacters(t *testing.T) {
	for _, table := range tables {
		z := countCharacters(table.x, table.y)
		if z != table.z {
			t.Fatalf("error : %v %v %v,but len:%v", strings.Join(table.x, ","), table.y, table.z, z)
		}
	}
}

//Test
func TestCountCharacters2(t *testing.T) {
	for _, table := range tables {
		z := countCharacters2(table.x, table.y)
		if z != table.z {
			t.Fatalf("error : %v %v %d,but len:%d", strings.Join(table.x, ","), table.y, table.z, z)
		}
	}
}
