package kyu8

import "strings"

// https://www.codewars.com/kata/57a1fd2ce298a731b20006a4/train/go

func IsPalindrome(str string) bool {
	lower := strings.ToLower(str)
	chars := []rune(lower)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return lower == string(chars)
}
