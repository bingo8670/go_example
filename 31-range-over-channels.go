// 在前面的例子中，我们看到for和range提供了对基本数据结构的迭代。我们也可以使用这种语法来迭代从通道接收到的值。
package main
import "fmt"
func main() {

  queue := make(chan string, 2)
  queue <- "one"
  queue <- "two"
  close(queue)

  for elem := range queue {
      fmt.Println(elem)
  }
}

// $ go run 31-range-over-channels.go
// one
// two
