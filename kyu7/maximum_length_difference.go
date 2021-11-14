package kyu7

// https://www.codewars.com/kata/5663f5305102699bad000056/train/go

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	max := 0
	diff := 0
	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a2); j++ {
			diff = Abs(len(a1[i]) - len(a2[j]))
			if diff > max {
				max = diff
			}
		}
	}
	return max
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
