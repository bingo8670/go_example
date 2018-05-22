// 当使用频道作为功能参数时，您可以指定频道是否仅用于发送或接收值。这种特异性增加了程序的类型安全性。
package main
import "fmt"

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}


// $ go run 26-channel-directions.go
// passed message
