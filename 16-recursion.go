// Go 支持 递归。这里是一个经典的阶乘示例。
package main
import "fmt"

// fact 函数在到达 fact(0) 前一直调用自身。
func fact(n int) int {
  if n == 0 {
      return 1
  }
  return n * fact(n-1)
}
func main() {
  fmt.Println(fact(7))
  // 7!   7的阶乘
}

// $ go run 16-recursion.go
// 5040
