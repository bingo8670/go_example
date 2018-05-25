// 在前面的例子中，我们讲过 for 和 range为基本的数据结构提供了迭代的功能。我们也可以使用这个语法来遍历从通道中取得的值。
package main
import "fmt"
func main() {

// 我们将遍历在 queue 通道中的两个值。
  queue := make(chan string, 2)
  queue <- "one"
  queue <- "two"
  close(queue)

// 这个 range 迭代从 queue 中得到的每个值。因为我们在前面 close 了这个通道，这个迭代会在接收完 2 个值之后结束。如果我们没有 close 它，我们将在这个循环中继续阻塞执行，等待接收第三个值
  for elem := range queue {
      fmt.Println(elem)
  }
}

// $ go run 31-range-over-channels.go
// one
// two
