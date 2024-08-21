package main

import (
	"strings"
)

// FullJustify 文本左右对齐
func FullJustify(words []string, maxWidth int) []string {
	res := []string{words[0]}
	n := len(words)
	if n == 1 {
		return res
	}
	for i := 1; i < n; i++ {
		word := words[i]
		//fmt.Println(word)
		if len(res[len(res)-1])+len(word) <= maxWidth {
			res[len(res)-1] += " " + word
		} else {
			// 重新排版
			//res[pos] = redo(res[pos], maxWidth)
			if len(word) == maxWidth {
				res = append(res, word)
			} else {
				res = append(res, word+" ")
			}
		}
	}
	return res
}

// 重新排版
func redo(str string, maxWidth int) string {
	ss := strings.Split(str, " ")
	with := 0
	slen := len(ss) - 1 // 单词间的间隔数
	if slen == 0 {
		return ss[0] + strings.Repeat("", maxWidth-len(ss[0]))
	}
	for _, v := range ss {
		with += len(v)
	}
	// 剩余
	remainder := maxWidth - with
	a := remainder / slen
	b := remainder % slen
	res := ss[0] + strings.Repeat("", a+b)

	for i := 1; i < slen-1; i++ {
		res += ss[i] + strings.Repeat("", a+b)
	}
	res += ss[slen] + strings.Repeat("", a)
	return res
}
