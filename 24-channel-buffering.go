// 默认情况下，通道是无缓冲的，这意味着如果有相应的接收（< - chan）准备好接收发送的值，它们将只接受发送（chan < - ）。缓冲的频道接受数量有限的数值，而这些数值没有相应的接收器。
package main
import "fmt"
func main() {

  messages := make(chan string, 2)

  messages <- "buffered"
  messages <- "channel"

  fmt.Println(<-messages)
  fmt.Println(<-messages)
}


// $ go run 24-channel-buffering.go
// buffered
// channel
