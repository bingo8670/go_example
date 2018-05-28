// 在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。对于更加复杂的情况，我们可以使用一个互斥锁来在 Go 协程间安全的访问数据。
package main
import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)
func main() {

// 在我们的例子中，state 是一个 map。
    var state = make(map[int]int)

// 这里的 mutex 将同步对 state 的访问。
    var mutex = &sync.Mutex{}
// 跟踪我们做了多少次读写操作。
    var readOps uint64
    var writeOps uint64

// 这里我们运行 100 个 Go 协程来重复读取 state。
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
              // 对于每次读取，我们选择一个键来访问，锁定互斥体以确保对状态的独占访问，读取所选键的值，解锁互斥锁，以及增加readOps计数。
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)

                time.Sleep(time.Millisecond)
            }
        }()
    }
// 启动10个goroutine来模拟写入，使用与读取操作相同的模式。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

// 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放。这个释放一般是自动处理的，像例如每个通道操作后或者 time.Sleep 的阻塞调用后相似，但是在这个例子中我们需要手动的处理。
    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}

// $ go run 37-mutexes.go
// readOps: 83285
// writeOps: 8320
// state: map[1:97 4:53 0:33 2:15 3:2]
