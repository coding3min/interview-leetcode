/*
* @Author: pzqu
* @Date:   2020-03-18 01:23
*/
package _160_Find_Words_That_Can_Be_Formed_by_Characters

import (
	"encoding/json"
)

//优化前
func countCharacters(words []string, chars string) int {
	cmap := make(map[uint8]int)

	for i := 0; i < len(chars); i++ {
		cmap[chars[i]] ++
	}

	count := 0
	tmpMap := make(map[uint8]int)

	for i := 0; i < len(words); i++ {
		ExtractInto(cmap, &tmpMap)
		j := 0
		for ; j < len(words[i]); j++ {
			if tmpMap[words[i][j]] <= 0 {
				break
			}
			tmpMap[words[i][j]] --
		}
		if j == len(words[i]) {
			count += len(words[i])
		}
	}

	return count
}

/*
	把interface类型转换成我们想要的struct类型
	这个通用方法可以转换成任意一个想要的类型
 */
func ExtractInto(source interface{}, to interface{}) error {
	b, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, to)
	return err
}

//优化后
func countCharacters2(words []string, chars string) int {
	var byteCount [26]int
	for _, char := range chars {
		byteCount[char-'a']++
	}

	ret := 0
	for _, word := range words {
		bc, match := byteCount, true
		for _, char := range word {
			if bc[char-'a'] <= 0 {
				match = false
				break
			}
			bc[char-'a']--
		}
		if match {
			ret += len(word)
		}
	}
	return ret
}
