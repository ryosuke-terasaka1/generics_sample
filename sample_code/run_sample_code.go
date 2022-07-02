package sample_code

import (
	"fmt"
)

func run_sumple_code() {
	fmt.Println(f([]MyInt{1, 2, 3, 4}))

	s := New[string]()
	s.Push("hello")
	s.Push("world")
	fmt.Println(s.Pop()) // world
	fmt.Println(s.Pop()) // hello

	stringValue := "111"
	intValue := 111
	boolValue := false

	fmt.Printf("sampleFuncGenerics1: %#v\n", sampleFuncGenerics1(stringValue))
	fmt.Printf("sampleFuncGenerics1: %#v\n", sampleFuncGenerics1(intValue))
	fmt.Printf("sampleFuncGenerics1: %#v\n", sampleFuncGenerics1(boolValue))
}
