package main

func Clamp(num, min, max float32) float32 {
	if num < min {
		return min
	} else if num > max {
		return max
	} else {
		return num
	}
}
