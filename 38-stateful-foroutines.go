// 在前面的例子中，我们使用显式锁定互斥体来跨多个goroutines同步对共享状态的访问。另一个选择是使用goroutines和通道的内置同步功能来实现相同的结果。这种基于频道的方法与Go的共享内存的想法保持一致，即通过沟通并让每条数据拥有1个goroutine。
package main
import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}
func main() {
// 跟踪我们做了多少次读写操作。
    var readOps uint64
    var writeOps uint64

    reads := make(chan *readOp)
    writes := make(chan *writeOp)
// 该goroutine在读取和写入通道上重复选择，在请求到达时对其进行响应。通过首先执行所请求的操作并且然后在响应通道上发送值resp以指示成功（以及在读取的情况下期望的值）来执行响应。
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()
// 每次读取都需要构建一个readOp，通过读取通道发送，并通过提供的resp通道接收结果。
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    // 启动10个goroutine来模拟写入，使用与读取操作相同的模式。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}

// $ go run 38-stateful-goroutines.go
// readOps: 71708
// writeOps: 7177
