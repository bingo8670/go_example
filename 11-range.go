// range迭代各种数据结构中的元素。
package main
import "fmt"
func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
    // range字符串迭代Unicode代码点。第一个值是第一个值rune和第二个rune自身的起始字节索引。
}


// $ go run 11-range.go
// sum: 9
// index: 1
// a -> apple
// b -> banana
// key: a
// key: b
// 0 103
// 1 111
