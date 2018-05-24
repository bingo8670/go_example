// 超时 对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的。得益于通道和 select，在 Go中实现超时操作是简洁而优雅的。
package main
import "time"
import "fmt"
func main() {

// 假如我们执行一个外部调用，并在 2 秒后通过通道 c1 返回它的执行结果。
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

// 这里是使用 select 实现一个超时操作。res := <- c1 等待结果，<-Time.After 等待超时时间 1 秒后发送的值。由于 select 默认处理第一个已准备好的接收操作，如果这个操作超过了允许的 1 秒的话，将会执行超时 case。
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

// 如果我允许一个长一点的超时时间 3 秒，将会成功的从 c2接收到值，并且打印出结果。
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}

// $ go run 28-timeouts.go
// timeout 1
// result 2
