// Maps是Go的内置关联数据类型（有时称为其他语言的散列或字典）。
package main
import "fmt"
func main() {

    m := make(map[string]int)
    // 类似与ruby 里到hash[key=>value]

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    // _, 表示空标识符，避免了必须为返回值声明所有变量。
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

// go run 10-maps.go
// map: map[k1:7 k2:13]
// v1:  7
// len: 2
// map: map[k1:7]
// prs: false
// map: map[foo:1 bar:2]
