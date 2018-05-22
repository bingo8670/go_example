// 定时器适用于当你想在将来做一些事情时使用.当你想要定期重复做某些事时，定时器就是用来做的。这是一个定时滴答滴答的例子，直到我们停止它。

package main
import "time"
import "fmt"
func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    //  1600 / 500 = 3
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
// $ go run 33-tickers.go
// Tick at 2018-05-21 20:39:19.058321156 +0800 CST m=+0.504074190
// Tick at 2018-05-21 20:39:19.557643324 +0800 CST m=+1.003381380
// Tick at 2018-05-21 20:39:20.056886638 +0800 CST m=+1.502609717
// Ticker stopped
