// for是Go中唯一的循环构造。
package main
import "fmt"
func main() {
    i := 1

// for是Go唯一的循环构造。

    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
        // break 方法用于停止该无限循环；
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
            // continue 方法用于跳过符合条件的情况；
        }
        fmt.Println(n)
    }
}


// $ go run 05-for.go
// 1
// 2
// 3
// 7
// 8
// 9
// loop
// 1
// 3
// 5
