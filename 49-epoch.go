// 时间戳
// 一般程序会有获取 Unix 时间的秒数，毫秒数，或者微秒数的需要。来看看如何用 Go 来实现。

package main
import "fmt"
import "time"
func main() {
// 分别使用带 Unix 或者 UnixNano 的 time.Now来获取从自协调世界时起到现在的秒数或者纳秒数。

    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)
// 注意 UnixMillis 是不存在的，所以要得到毫秒数的话，你要自己手动的从纳秒转化一下。

    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)
// 你也可以将协调世界时起的整数秒或者纳秒转化到相应的时间。

    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}

// $ go run 49-epoch.go
// 2012-10-31 16:13:58.292387 +0000 UTC
// 1351700038
// 1351700038292
// 1351700038292387000
// 2012-10-31 16:13:58 +0000 UTC
// 2012-10-31 16:13:58.292387 +0000 UTC
