package kyu8

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/5a2fd38b55519ed98f0000ce/train/go

func MultiTable(number int) string {
	var sb strings.Builder
	for i := 1; i <= 10; i++ {
		str := fmt.Sprintf("%d * %d = %d\n", i, number, i*number)
		sb.WriteString(str)
	}

	return strings.TrimRight(sb.String(), "\n")
}
