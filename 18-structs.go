// Go的结构是字段类型的集合。它们用于将数据分组在一起以形成记录。
package main
import "fmt"

// 这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
    name string
    age  int
}
func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

// 省略的字段将被初始化为零值。
    fmt.Println(person{name: "Fred"})

// & 前缀生成一个结构体指针。
    fmt.Println(&person{name: "Ann", age: 40})

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}

// $ go run 18-structs.go
// {Bob 20}
// {Alice 30}
// {Fred 0}
// &{Ann 40}
// Sean
// 50
// 51
