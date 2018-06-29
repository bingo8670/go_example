// 退出
// 使用 os.Exit 来立即进行带给定状态的退出。

package main
import "fmt"
import "os"
func main() {
// 当使用 os.Exit 时 defer 将不会 执行，所以这里的 fmt.Println将永远不会被调用。

    defer fmt.Println("!")
// 退出并且退出状态为 3。

    os.Exit(3)
}
// 注意，不像例如 C 语言，Go 不使用在 main 中返回一个整数来指明退出状态。如果你想以非零状态退出，那么你就要使用 os.Exit。
//
// 如果你使用 go run 来运行 exit.go，那么退出状态将会被 go捕获并打印。

// $ go run 65-exit.go
// exit status 3
// 使用编译并执行一个二进制文件的方式，你可以在终端中查看退出状态。
//
// $ go build 65-exit.go
// $ ./65-exit
// $ echo $?
// 3
