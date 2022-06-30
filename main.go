package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(f([]MyStr{"1", "2", "3", "4"}))
}

func f[T Integers](xs []T) []int {
	var result []int
	for _, x := range xs {
		result = append(result, x.Integer())
	}
	return result
}

type Integers interface {
	Integer() int
}

type MyStr string

func (i MyStr) Integer() (int, error) {
	return strconv.Atoi(string(i))
}
