// functions 函数, 是Go的核心
package main
import "fmt"

func plus(a int, b int) int {

// Go 需要明确的返回值，它不会自动返回最后一个表达式的值
  return a + b
}

func plusPlus(a, b, c int) int {
  return a + b + c
}
func main() {

// 通过 name(args) 来调用一个函数，
  res := plus(1, 2)
  fmt.Println("1+2 =", res)
  res = plusPlus(1, 2, 3)
  fmt.Println("1+2+3 =", res)
}

// $ go run 12-functions.go
// 1+2 = 3
// 1+2+3 = 6
