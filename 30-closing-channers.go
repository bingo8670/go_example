// 关闭频道表示不会再发送任何值。这对于将完成情况传达给频道的接收器很有用。
package main
import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}


// $ go run 30-closing-channels.go  (输出信息顺序不固定，高并发)
// sent job 1
// received job 1
// sent job 2
// received job 2
// sent job 3
// received job 3
// sent all jobs
// received all jobs
