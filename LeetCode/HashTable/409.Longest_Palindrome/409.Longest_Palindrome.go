/*
* @Author: pzqu
* @Date:   2020-03-19 00:19
*/
package _09_Longest_Palindrome

func longestPalindrome(s string) int {
	var byteCount ['z'-'A' + 1]int
	for _,v :=  range s{
		byteCount[v-'A']++
	}
	var count,last int
	for _,v :=range byteCount{
		count += v/2*2
		if v%2==1{
			last  = 1
		}
	}
	count += last
	return count
}
