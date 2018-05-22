// Go内置了对多个返回值的支持。这个特性经常用在惯用的Go中，例如从函数返回结果和错误值。
package main
import "fmt"

func vals() (int, int) {
  return 3, 7
}
func main() {

  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c := vals()
  fmt.Println(c)
}


// $ go run 13-multiple-return-values.go
// 3
// 7
// 7
