package code

import "strings"

// 最后一个单词的长度
// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
func lengthOfLastWord(s string) int {
	end := len(s) - 1
	// 去掉末尾的空格
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return 0
	}
	// 计算最后一个单词的长度
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return end - start
}

// 反转字符串
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
func reverseString(str string) {
	// 双指针
	s := []byte(str)
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	// 横向扫描
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

// 两个字符串的最长公共前缀
func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

// 反转字符串中的单词
// 给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
// 说明：
// 输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
// 翻转后单词间应当仅用一个空格分隔。
// 翻转后的字符串中不应包含额外的空格。
func reverseWords(s string) string {
	// 双指针
	// 去掉首尾空格
	s = strings.Trim(s, " ")
	res := make([]byte, 0)
	left, right := len(s)-1, len(s)-1
	for left >= 0 {
		for left >= 0 && s[left] != ' ' {
			left--
		}
		res = append(res, []byte(s[left+1:right+1])...)
		res = append(res, ' ')
		// 去掉单词间多余的空格
		for left >= 0 && s[left] == ' ' {
			left--
		}
		// 下一个单词的尾部
		right = left
	}
	return string(res[:len(res)-1])
}

// N字形变换
// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 N 字形排列
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如：
// "LEETCODEISHIRING" 行数为 3 时，排列如下：
// L   C   I   R
// E T O E S I I G
// E   D   H   N
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如：
// "LCIRETOESIIGEDHN"
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, min(numRows, len(s)))
	curRow := 0
	goingDown := false
	for _, c := range s {
		rows[curRow] += string(c)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}
	res := ""
	for _, row := range rows {
		res += row
	}
	return res
}
