// Go有各种值类型，包括字符串，整数，浮点数，布尔值等。
package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// $ go run 02-values.go
// golang
// 1+1 = 2
// 7.0/3.0 = 2.3333333333333335
// false
// true
// false
