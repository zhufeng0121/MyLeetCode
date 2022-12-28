package SwordToOffer

import "strings"

// 只考虑字母和数字字符，可以忽略大小写

// 思路应该是双指针 left,right 当 left == right 时退出
func isPalindromeV(s string) bool {

	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		if !isLetterOrNum(s[left]) {
			left++
			continue
		}
		if !isLetterOrNum(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--

	}

	return true

}

func isLetterOrNum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '1' && c <= '9') || (c >= 'A' && c <= 'Z')

}
