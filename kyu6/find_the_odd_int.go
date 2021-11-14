package kyu6

// https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

func FindOdd(seq []int) int {
	// Could be solved using XOR
	var number int
	vals := map[int]int{}
	for i := 0; i < len(seq); i++ {
		if _, mapContainsKey := vals[seq[i]]; mapContainsKey {
			vals[seq[i]] += 1
		} else {
			vals[seq[i]] = 1
		}
	}

	for key, val := range vals {
		if val%2 != 0 {
			number = key
		}
	}
	return number
}
