// 可以使用任意数量的结尾参数来调用变量函数。例如，fmt.Println是一个常用的可变参数函数。
package main
import "fmt"

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

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}

// $ go run 14-variadic-functions.go 
// [1 2] 3
// [1 2 3] 6
// [1 2 3 4] 10
