package nonrepeating

/**
目标：寻找最长不含有重复字符的子串
算法：
对于每一个字母x
* lastOccurred[x]不存在，或者<start->无需操作
* lastOccurred[x]>=start->更新start
* 更新lastOccurred[x],更新maxLength

注：习题来自于leetcode
https://leetcode.com/problems/longest-substring-without-repeating-characters/description
*/

func LengthOfNonRepeatingSubStr(s string) int {
	//代码优化思路：空间换时间
	lastOccurred := make([]int, 0xffff) //65535
	//lastOccurred[0x65] = 1
	//lastOccurred[0x8bFE] = 6
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	//lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	//string转换为byte切片
	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			//if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
