package kyu6

// https://www.codewars.com/kata/585d7d5adb20cf33cb000235/train/go

func FindUniq(arr []float32) float32 {
	if arr[0] == arr[1] {
		for i := 2; i < len(arr); i++ {
			if arr[i] != arr[0] {
				return arr[i]
			}
		}
	} else if arr[0] == arr[2] {
		return arr[1]
	}
	return arr[0]
}
