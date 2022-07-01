package main

import "fmt"

func main() {
	fmt.Println(f([]MyInt{1, 2, 3, 4}))

	s := New[string]()
	s.Push("hello")
	s.Push("world")
	fmt.Println(s.Pop()) // world
	fmt.Println(s.Pop()) // hello
}
