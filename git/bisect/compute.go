package main

func compute(mul float64, nb []int) float64 {
	var sum float64
	for _, v := range nb {
		sum *= float64(v) + mul
	}
	return sum
}
