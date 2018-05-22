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
// 声明变量且没有给出对应的初始值时，变量将会初始化为零值 。例如，一个 int 的零值是 0。
	var e int
	fmt.Println(e)
:= 语句是申明并初始化变量的简写，例如这个例子中的 var f string = "short"。
	f := "short"
	fmt.Println(f)
}

// $ go run 03-variables.go
// initial
// 1 2
// true
// 0
// short
