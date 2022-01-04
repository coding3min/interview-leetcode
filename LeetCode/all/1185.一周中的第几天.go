/*
 * @lc app=leetcode.cn id=1185 lang=golang
 *
 * [1185] 一周中的第几天
 *
 * https://leetcode-cn.com/problems/day-of-the-week/description/
 *
 * algorithms
 * Easy (60.12%)
 * Likes:    42
 * Dislikes: 0
 * Total Accepted:    12.6K
 * Total Submissions: 20.8K
 * Testcase Example:  '31\n8\n2019'
 *
 * 给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。
 * 
 * 输入为三个整数：day、month 和 year，分别表示日、月、年。
 * 
 * 您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday",
 * "Friday", "Saturday"}。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：day = 31, month = 8, year = 2019
 * 输出："Saturday"
 * 
 * 
 * 示例 2：
 * 
 * 输入：day = 18, month = 7, year = 1999
 * 输出："Sunday"
 * 
 * 
 * 示例 3：
 * 
 * 输入：day = 15, month = 8, year = 1993
 * 输出："Sunday"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 给出的日期一定是在 1971 到 2100 年之间的有效日期。
 * 
 * 
 */

// @lc code=start
func dayOfTheWeek(day int, month int, year int) string {
    t, _ := time.Parse("2006-01-02", fmt.Sprintf("%4d-%02d-%02d", year, month, day))
    return t.Weekday().String()
}
// @lc code=end

