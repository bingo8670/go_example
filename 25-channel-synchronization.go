// 我们可以使用通道来跨goroutines同步执行。以下是使用阻塞接收等待goroutine完成的示例。
package main
import "fmt"
import "time"

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}
func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}

// $ go run 25-channel-synchronization.go
// working...done
