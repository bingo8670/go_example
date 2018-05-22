// goroutine是一个轻量级的执行线程。
package main
import "fmt"
func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
func main() {

    f("direct")

    go f("goroutine")


    go func(msg string) {
        fmt.Println(msg)
    }("going")

    fmt.Scanln()
    fmt.Println("done")
}

// $ go run 22-goroutines.go
// direct : 0
// direct : 1
// direct : 2
// goroutine : 0
// goroutine : 1
// goroutine : 2
// going  （位置不固定）
// <enter>
// done
