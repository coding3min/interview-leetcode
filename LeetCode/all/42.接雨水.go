/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (52.59%)
 * Likes:    1984
 * Dislikes: 0
 * Total Accepted:    190.4K
 * Total Submissions: 351.9K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == height.length
 * 0 
 * 0 
 * 
 * 
 */

// @lc code=start
// 双指针，空间复杂度O1
func trap(height []int) int {
    length := len(height)
    if length <3{
        return 0
	}
	// 双指针的题目,初始化左右指针和左右最大值
    left,right,leftMax,rightMax := 0,length-1,height[0],height[length-1] 
	res := 0
	
	// 只要一边比另一边大，小的那边进一位，只要比原来的小，能装的水就是和这一边最大值的差值
 	// 所以要一直维护一个左最大值和右最大值，并且要保证左最大在新值左边，右最大在新值右边
	// 不用担心接不住雨水，因为另一边一定能兜住，这是站在全局思考的方式
    for left<right{
		   // 如果左最大<右最大,左++ 反之 右--，得到当前新值
        if leftMax<rightMax{
            left = left+1
            leftNew := height[left]
			   // 因为左最大<右最大，左起到低洼的作用，新值小于低洼得到当前水
			   // 否则更新最大值
            if leftNew < leftMax{
                res += leftMax - leftNew 
            }else{
                leftMax = leftNew
            }
        }else{
            right = right - 1
            rightNew := height[right]
            if rightNew < rightMax{
                res += rightMax - rightNew 
            }else{
                rightMax = rightNew
            }
        }
    } 
    return res
}

// 用栈的方式，空间复杂度高
// func trap(height []int) int {
// 	if len(height)<3{
// 		return 0
// 	}
// 	stack := []int{}
// 	res := 0
// 	for k,v := range height{
// 		// 只要新值大于前一个值就进入循环，说明有可能出来低洼
// 		for len(stack)>0 && height[stack[len(stack)-1]] < v{
//             topKey := stack[len(stack)-1]
//             stack = stack[:len(stack)-1]
//             if len(stack) == 0{
//                 break
//             }
//             secondKey := stack[len(stack)-1]
// 			// 左右两边下标的差值，就是低洼的长度，用栈的方法是按层累计的
// 			// 比如 [5 3] 4；5和4之间 下标 2-0-1 = 1，中间是一格
// 			water := k - secondKey - 1
// 			// 低洼是由左右两边，最小的那一边决定的 比如 [5,3],4 是由4决定的
// 			// 所以水高度是4-3
//             minValue := min(height[secondKey],v)
// 			// 因为是按层，假如原始数组是 [5,3,2,4],那么应该是如下顺序
// 			// stack[5,3,2],4 -> stack[5,3],4 res+=(3-2)*(3-1-1)
// 			// stack[5,3],4 -> stack[5],4 res+=(4-3)*(3-0-1)
//             res += (minValue - height[topKey]) * water
// 		}	
//         if len(stack) == 0 || height[stack[len(stack)-1]]>=v{
// 			stack = append(stack,k)
// 		}
// 	}
// 	return res
// }

// func min(x,y int) int {
//     if x<y{
//         return x
//     }
//     return y
// }

// @lc code=end

