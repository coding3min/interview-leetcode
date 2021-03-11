/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 *
 * https://leetcode-cn.com/problems/bulls-and-cows/description/
 *
 * algorithms
 * Medium (49.44%)
 * Likes:    116
 * Dislikes: 0
 * Total Accepted:    23.6K
 * Total Submissions: 47.7K
 * Testcase Example:  '"1807"\n"7810"'
 *
 * 你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：
 * 
 * 
 * 你写出一个秘密数字，并请朋友猜这个数字是多少。
 * 朋友每猜测一次，你就会给他一个提示，告诉他的猜测数字中有多少位属于数字和确切位置都猜对了（称为“Bulls”,
 * 公牛），有多少位属于数字猜对了但是位置不对（称为“Cows”, 奶牛）。
 * 朋友根据提示继续猜，直到猜出秘密数字。
 * 
 * 
 * 请写出一个根据秘密数字和朋友的猜测数返回提示的函数，返回字符串的格式为 xAyB ，x 和 y 都是数字，A 表示公牛，用 B 表示奶牛。
 * 
 * 
 * xA 表示有 x 位数字出现在秘密数字中，且位置都与秘密数字一致。
 * yB 表示有 y 位数字出现在秘密数字中，但位置与秘密数字不一致。
 * 
 * 
 * 请注意秘密数字和朋友的猜测数都可能含有重复数字，每位数字只能统计一次。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: secret = "1807", guess = "7810"
 * 输出: "1A3B"
 * 解释: 1 公牛和 3 奶牛。公牛是 8，奶牛是 0, 1 和 7。
 * 
 * 示例 2:
 * 
 * 输入: secret = "1123", guess = "0111"
 * 输出: "1A1B"
 * 解释: 朋友猜测数中的第一个 1 是公牛，第二个或第三个 1 可被视为奶牛。
 * 
 * 
 * 
 * 说明: 你可以假设秘密数字和朋友的猜测数都只包含数字，并且它们的长度永远相等。
 * 
 */

// @lc code=start
func getHint(secret string, guess string) string {
	secretArray := [10]int{}
	cowArray := [10]int{}
	bullCnt,cowCnt := 0,0
	for k,v := range strings.Split(secret,""){
		if v == string(guess[k]){
			bullCnt++
		}else{
			// 这里已经排除了公牛，统计了剩余字条和guess字符的次数
			secretArray[secret[k]-'0']++
			cowArray[guess[k] - '0']++
		}
	}
	for i:=0;i<len(cowArray);i++{
		// 这里为什么要取最小值，是因为只取重叠的部分
		// "111000" "000111" 应该输出6B
		if cowArray[i]<secretArray[i]{
			cowCnt += cowArray[i]
		}else{
			cowCnt += secretArray[i]
		}
	}
	return strconv.Itoa(bullCnt) + "A" + strconv.Itoa(cowCnt) +"B"
}

func getHint(secret string, guess string) string {
	secretArray := [10]int{}
	cowArray := [10]int{}
	bullCnt,cowCnt := 0,0
	for k,v := range strings.Split(secret,""){
		if v == string(guess[k]){
			bullCnt++
		}else{
			// 这里已经排除了公牛，统计了剩余字条和guess字符的次数
			gInt,_ := strconv.Atoi(string(guess[k]))
			secretArray[secret[k]-'0']++
			cowArray[guess[k] - '']++
		}
	}
	for i:=0;i<len(cowArray);i++{
		// 这里为什么要取最小值，是因为只取重叠的部分
		// "111000" "000111" 应该输出6B
		if cowArray[i]<secretArray[i]{
			cowCnt += cowArray[i]
		}else{
			cowCnt += secretArray[i]
		}
	}
	return strconv.Itoa(bullCnt) + "A" + strconv.Itoa(cowCnt) +"B"
}
// @lc code=end

