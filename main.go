package main

import (
	"fmt"
	"generics_sample/test_code"
)

func main() {
	int_slice := []int{1, 2, 3, 4, 5}
	int_av := test_code.Average(int_slice)
	fmt.Printf("int_cal=%v\n", int_av)
	float_slice := []float64{1.2, 2.3, 3.0, 4.9, 5.1}
	float_av := test_code.Average(float_slice)
	fmt.Printf("float_cal=%v\n", float_av)

}
