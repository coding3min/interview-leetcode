//题目链接：https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/?envType=study-plan&id=lcof
package main

//这里我想先解释一下为什么负数范围比正数范围大 1 .
//因为对有符号数，在计算机中表达时，最高位约定为符号位，符号位为 0 时代表正数，符号位为 1 时代表负数。
//对 int32，真正表达数值的部分就是剩下的 31 位，正数范围很容易理解为 1~2^31-1（2147483647）
//有一个特殊的值，就是 0 值，对符号位 1 和 0 的时候，都是 0，用两个形式表达同一个数无疑是浪费的，
//所以，我们就规定，符号位为 1 时的全 0，表达 -2147483648，这也是负数表示的范围比正数多1的原因。
//综上，对于任意位的，无论是8位，16位，32位甚至64位的整数类型表示范围的计算公式为：
//如总位数为n位，那么有符号数的范围为：-2^(n-1) ~ 2^(n-1)-1，无符号数的表示范围为：0~2^n-1


//此题解参考自链接：https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/submissions/

//首部空格： 删除之即可；
//符号位： 三种情况，即 ''+'' , ''-'' , ''无符号" ；新建一个变量保存符号位，返回前判断正负即可。
//非数字字符： 遇到首个非数字的字符时，应立即返回。
//数字字符：
//字符转数字： “此数字的 ASCII 码” 与 “ 00 的 ASCII 码” 相减即可；
//数字拼接： 若从左向右遍历数字，设当前位字符为 c ，当前位数字为 x ，数字结果为 res，则数字拼接公式为：res = 10 * res + x
//x = ascii(c) - ascii('0')
//
//接下来是比较关键的数字越界处理：
//题目要求返回的数值范围应在 [-2^{31}, 2^{31} - 1]，因此需要考虑数字越界问题。
//而由于题目指出 环境只能存储 32 位大小的有符号整数 ，因此判断数字越界时，要始终保持 res 在 int 类型的取值范围内。
//
//在每轮数字拼接前，判断 res 在此轮拼接后是否超过 2147483647，若超过则加上符号位直接返回。
//设数字拼接边界 boundary = 2147483647 // 10 = 214748364，则以下两种情况越界：
//
//1. res > boundry， 拼接后越界
//2. res = boundry && x > 7，拼接后越界
//
//这种处理方式非常巧妙，> ‘7’ 这个操作看似只考虑到了最大值，其实也考虑了最小值，只要大于，直接返回即可。
const (
	MAX_INT = 1 << 31 - 1
	MIN_INT = -1 * (1 << 31)
)

func strToInt(str string) int {
	// index 为 str 遍历到的位置
	index := 0
	// 为防止读取空格出现越界panic，需要提前处理str为空字符或全空格的情况
	if str == ""{
		return 0
	}
	for i:=0;i<len(str);i++{
		if str[i] != ' '{
			break
		}
		if i == len(str) - 1{
			return 0
		}
	}
	// 读取前置空格
	for str[index] == ' '{
		index ++
	}
	// 符号位初始化为 0
	sign := 1
	if str[index] == '+'{
		index ++
	} else if str[index] == '-' {
		sign = -1
		index ++
	}
	boundary := MAX_INT/10
	res := 0
	for _,s := range str[index:]{
		// 读取到非数字字符，直接跳出
		if s < '0' || s > '9'{
			break
		}
		// 预计出现越界情况，根据符号位返回最大或最小值
		if res > boundary || res == boundary && s > '7'{
			if sign == 1{
				return MAX_INT
			} else {
				return MIN_INT
			}
		}
		res = res * 10 + int(s - '0')

	}
	return sign * res
}