package kyu6

// https://www.codewars.com/kata/541c8630095125aba6000c00/train/go

func DigitalRoot(n int) int {
	// 1 + ((n - 1) mod (b - 1)) in base b
	return 1 + (n-1)%9
}
