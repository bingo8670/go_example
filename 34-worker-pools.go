// 在这个例子中，我们将看到如何使用 Go 协程和通道实现一个工作池 。
package main
import "fmt"
import "time"

// 这是我们将要在多个并发实例中支持的任务了。这些执行者将从 jobs 通道接收任务，并且通过 results 发送对应的结果。我们将让每个任务间隔 1s 来模仿一个耗时的任务。
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

// 为了使用 worker 工作池并且收集他们的结果，我们需要2 个通道。
func main() {

    jobs := make(chan int, 100)
    results := make(chan int, 100)
    // 需要向他们发送工作并收集他们的结果。我们为此制作了2个频道。
// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务。
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 5; a++ {
        <-results
    }
}

// $ time go run 34-worker-pools.go
// worker 1 started  job 1
// worker 2 started  job 2
// worker 3 started  job 3
// worker 1 finished job 1
// worker 1 started  job 4
// worker 2 finished job 2
// worker 2 started  job 5
// worker 3 finished job 3
// worker 1 finished job 4
// worker 2 finished job 5
// real	0m2.358s
