// 通道上的基本发送和接收处于阻塞状态。但是，我们可以使用带有默认子句的select来实现非阻塞式发送，接收，甚至是非阻塞式多路选择。
package main
import "fmt"
func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}

// $ go run 29-non-blocking-channel-operations.go
// no message received
// no message sent
// no activity
