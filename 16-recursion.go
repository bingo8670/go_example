// Go支持递归函数。这是一个经典的因子示例。
package main
import "fmt"

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
