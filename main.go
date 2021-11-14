package main

import (
	"codewars/kyu7"
	// "codewars/kyu8"
	"fmt"
)

func main() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	fmt.Println(kyu7.MxDifLg(a1, a2))
	s1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	s2 := []string{"bbbaaayddqbbrrrv"}
	fmt.Println(kyu7.MxDifLg(s1, s2))
}
