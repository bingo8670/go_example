// 在前面的例子中，我们看到了如何用原子操作来管理简单的计数器状态。对于更复杂的状态，我们可以使用互斥锁跨多个goroutines安全地访问数据。
package main
import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)
func main() {

    var state = make(map[int]int)

    var mutex = &sync.Mutex{}
// 跟踪我们做了多少次读写操作。
    var readOps uint64
    var writeOps uint64
// 启动100个goroutines来执行对该状态的重复读取，每个goroutine每毫秒执行一次。
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
