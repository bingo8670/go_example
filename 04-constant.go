// Go支持字符，字符串，布尔值和数值的常量。
package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
  fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}


// $ go run 04-constant.go
// constant
// 6e+11
// 600000000000
// -0.28470407323754404
