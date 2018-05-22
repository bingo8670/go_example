// Go内置了对多返回值的支持。这个特性经常用在惯用的Go中，例如从函数返回结果和错误值。
package main
import "fmt"

// (int, int) 在这个函数中标志着这个函数返回 2 个 int。
func vals() (int, int) {
  return 3, 7
}
func main() {

  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)
  
// 如果你仅仅想返回值的一部分的话，你可以使用空白定义符 _。
  _, c := vals()
  fmt.Println(c)
}


// $ go run 13-multiple-return-values.go
// 3
// 7
// 7
