package main

import (
	"strconv"
)

// fは型パラメータを持つ関数
// Tは型パラメータ
// インタフェースStringerは、Tに対する型制約として使われている
func f[T Intger](xs []T) []string {
	var result []string
	for _, x := range xs {
		// xは型制約StringerによりString()メソッドが使える
		result = append(result, x.String())
	}
	return result
}

type Intger interface {
	String() string
}

type MyInt int

// MyIntはStringerを実装する
func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}
