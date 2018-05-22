// 在Go中，数组是具有特定长度的编号的元素序列。
package main
import "fmt"
func main() {

    var a [5]int
    // int 默认为0
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])


    fmt.Println("len:", len(a))


    b := [5]int{1, 2, 3, 4, 5}
    // 数组声明句法
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

// go run 08-arrays.go
// emp: [0 0 0 0 0]
// set: [0 0 0 0 100]
// get: 100
// len: 5
// dcl: [1 2 3 4 5]
// 2d:  [[0 1 2] [1 2 3]]
