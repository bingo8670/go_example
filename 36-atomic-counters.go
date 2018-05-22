// 在Go中管理状态的主要机制是通过渠道进行沟通。我们在工作池中看到了这一点。还有一些管理状态的选项。这里我们来看看使用sync / atomic包来处理由多个goroutine访问的原子计数器。
package main
import "fmt"
import "time"
import "sync/atomic"
func main() {

    var ops uint64

    for i := 0; i < 50; i++ {
        go func() {
            for {

                atomic.AddUint64(&ops, 1)

                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}

// $ go run atomic-counters.go
// ops: 41419  计数器数值不固定
