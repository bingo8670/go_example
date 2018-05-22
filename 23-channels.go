// 通道是连接并发混合程序的管道。您可以从一个goroutine将值发送到通道中，并将这些值接收到另一个goroutine中。
package main
import "fmt"
func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}

// $ go run 23-channels.go
// ping
