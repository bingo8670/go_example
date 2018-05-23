// Go支持可以形成闭包的匿名函数。匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
package main
import "fmt"

// 这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。这个返回的函数使用闭包的方式 隐藏 变量 i。
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {

// 我们调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt。这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时都会更新 i 的值。
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
