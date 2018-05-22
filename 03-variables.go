// 在Go中，变量被显式声明并被编译器用于例如检查函数调用的类型正确性。
package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}

// $ go run 03-variables.go
// initial
// 1 2
// true
// 0
// short
