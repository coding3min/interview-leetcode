/*
* @Author: pzqu
* @Date:   2020-03-16 23:30
*/
package String

import (
	"bytes"
	"strconv"
)

func compressString(S string) string {
	if len(S) <= 2 {
		return S
	}
	count := 1
	var res bytes.Buffer
	res.WriteString(string(S[0]))
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			count ++
		} else {
			res.WriteString(strconv.Itoa(count))
			res.WriteString(string(S[i]))
			count = 1
		}
	}

	res.WriteString(strconv.Itoa(count))
	resStr := res.String()
	if len(resStr) < len(S) {
		return resStr
	} else {
		return S
	}
}
