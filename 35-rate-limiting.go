// 限速是控制资源利用率和维护服务质量的重要机制。 Go优雅地支持goroutines，渠道和代码的速率限制。
package main
import "time"
import "fmt"
func main() {

    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3)

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

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
