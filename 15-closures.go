// Go支持可以形成闭包的匿名函数。当你想定义一个内联函数而不必命名它时，匿名函数很有用。
package main
import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}


// $ go run 15-closures.go
// 1
// 2
// 3
// 1
