// range 遍历，迭代各种数据结构中的元素。
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

// range 在 map 中迭代键值对。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

// range 在字符串中迭代 unicode 编码。第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
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
