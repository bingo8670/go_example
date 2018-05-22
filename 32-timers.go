// 我们经常希望在将来某个时候执行Go代码，或者在某个时间间隔内重复执行Go代码。 Go的内置计时器和自动收报机功能使这两项任务变得简单。我们先看看定时器，然后看看代码。
package main
import "time"
import "fmt"
func main() {

    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C
    fmt.Println("Timer 1 expired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
// $ go run 32-timers.go
// Timer 1 expired
// Timer 2 stopped
