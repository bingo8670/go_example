// 速率限制(英) 是一个重要的控制服务资源利用和质量的途径。Go 通过 Go 协程、通道和打点器优美的支持了速率限制。
package main
import "time"
import "fmt"
func main() {

    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

// 这个 limiter 通道将每 200ms 接收一个值。这个是速率限制任务中的管理器。
    limiter := time.Tick(200 * time.Millisecond)

// 通过在每次请求前阻塞 limiter 通道的一个接收，我们限制自己每 200ms 执行一次请求。
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

// 有时候我们想临时进行速率限制，并且不影响整体的速率控制我们可以通过通道缓冲来实现。这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制。
    burstyLimiter := make(chan time.Time, 3)

// 想将通道填充需要临时改变3次的值，做好准备。
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

// 每 200 ms 我们将添加一个新的值到 burstyLimiter中，直到达到 3 个的限制。
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

// 现在模拟超过 5 个的接入请求。它们中刚开始的 3 个将由于受 burstyLimiter 的“脉冲”影响。
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}

// $ go run rate-limiting.go
// request 1 2018-05-22 10:41:19.033757555 +0800 CST m=+0.201206445
// request 2 2018-05-22 10:41:19.233071839 +0800 CST m=+0.400514750
// request 3 2018-05-22 10:41:19.433202143 +0800 CST m=+0.600639051
// request 4 2018-05-22 10:41:19.633190233 +0800 CST m=+0.800621141
// request 5 2018-05-22 10:41:19.83332706  +0800 CST m=+1.000751965
// 运行我们的程序，我们看到第一批请求按照需要每约200毫秒处理一次。

// request 1 2018-05-22 10:41:19.833469628 +0800 CST m=+1.000894529
// request 2 2018-05-22 10:41:19.833554623 +0800 CST m=+1.000979521
// request 3 2018-05-22 10:41:19.833575137 +0800 CST m=+1.001000035
// request 4 2018-05-22 10:41:20.034167436 +0800 CST m=+1.201586316
// request 5 2018-05-22 10:41:20.233651952 +0800 CST m=+1.401064848
// 对于第二批请求，由于可突发速率限制，我们立即为前3个服务，然后为剩余的2个服务，每个延迟约200毫秒。
