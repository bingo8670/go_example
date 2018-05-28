// 在前面的例子中，我们用互斥锁进行了明确的锁定来让共享的state 跨多个 Go 协程同步访问。另一个选择是使用内置的 Go协程和通道的的同步特性来达到同样的效果。这个基于通道的方法和 Go 通过通信以及 每个 Go 协程间通过通讯来共享内存，确保每块数据有单独的 Go 协程所有的思路是一致的。
package main
import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// 在这个例子中，state 将被一个单独的 Go 协程拥有。这就能够保证数据在并行读取时不会混乱。为了对 state 进行读取或者写入，其他的 Go 协程将发送一条数据到拥有的 Go协程中，然后接收对应的回复。结构体 readOp 和 writeOp封装这些请求，并且是拥有 Go 协程响应的一个方式。
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
// 和前面一样，我们将计算我们执行操作的次数。
    var readOps uint64
    var writeOps uint64

// reads 和 writes 通道分别将被其他 Go 协程用来发布读和写请求。
    reads := make(chan *readOp)
    writes := make(chan *writeOp)

// 这个就是拥有 state 的那个 Go 协程，和前面例子中的map一样，不过这里是被这个状态协程私有的。这个 Go 协程反复响应到达的请求。先响应到达的请求，然后返回一个值到响应通道 resp 来表示操作成功（或者是 reads 中请求的值）
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
    
// 启动 100 个 Go 协程通过 reads 通道发起对 state 所有者Go 协程的读取请求。每个读取请求需要构造一个 readOp，发送它到 reads 通道中，并通过给定的 resp 通道接收结果。
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
