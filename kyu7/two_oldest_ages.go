package kyu7

// https://www.codewars.com/kata/511f11d355fe575d2c000001/train/go

func TwoOldestAges(ages []int) [2]int {
	oldestAges := [2]int{0, 0}
	for _, val := range ages {
		if val > oldestAges[1] {
			oldestAges[0], oldestAges[1] = oldestAges[1], val
		} else if val > oldestAges[0] {
			oldestAges[0] = val
		}
	}
	return oldestAges
}
