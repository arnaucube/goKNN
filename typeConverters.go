package main

func arrayIntToFloat64(arr []int) []float64 {
	var r []float64
	for i := 0; i < len(arr); i++ {
		r = append(r, float64(arr[i]))
	}
	return r

}
