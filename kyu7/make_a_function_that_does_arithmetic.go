package kyu7

// https://www.codewars.com/kata/583f158ea20cfcbeb400000a/train/go

func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	default:
		return 0
	}
}
