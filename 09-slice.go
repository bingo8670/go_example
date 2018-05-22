// Slices是Go中的一个关键数据类型，与数组相比，它提供了一个更强大的序列接口。
package main
import "fmt"
func main() {
// 不像数组，slice 的类型仅由它所包含的元素决定（不像数组中还需要元素的个数）。要创建一个长度非零的空slice，需要使用内建的方法 make。这里我们创建了一个长度为3的 string 类型 slice（初始化为零值）。
    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

// slice 支持比数组更多的操作。其中一个是内建的 append，它返回一个包含了一个或者多个新值的 slice。注意我们接受返回由 append返回的新的 slice 值。
    s = append(s, "d")
    s = append(s, "e", "f")
    // append 方法添加元素
    fmt.Println("apd:", s)

// Slice 也可以被 copy。这里我们创建一个空的和 s 有相同长度的 slice c，并且将 s 复制给 c。
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

// 序号从2到5个元素
    l := s[2:5]

    fmt.Println("sl1:", l)

// 序号从开头到6的元素
    l = s[:5]
    fmt.Println("sl2:", l)

// 序号从2到末尾到元素
    l = s[2:]
    fmt.Println("sl3:", l)

// 在一行代码中声明并初始化一个 slice 变量。
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

// go run 09-slices.go
// emp: [  ]
// set: [a b c]
// get: c
// len: 3
// apd: [a b c d e f]
// cpy: [a b c d e f]
// sl1: [c d e]
// sl2: [a b c d e]
// sl3: [c d e f]
// dcl: [g h i]
// 2d:  [[0] [1 2] [2 3 4]]
