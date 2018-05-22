// 可变参数函数，可以使用任意数量的结尾参数来调用变量函数。例如，fmt.Println是一个常用的可变参数函数。
package main
import "fmt"

// 这个函数使用任意数目的 int 作为参数。
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
func main() {

    sum(1, 2)
    sum(1, 2, 3)

// 如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}

// $ go run 14-variadic-functions.go
// [1 2] 3
// [1 2 3] 6
// [1 2 3 4] 10
