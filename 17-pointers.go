// Go支持指针，允许您在程序中传递对值和记录的引用。
package main
import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}
func main() {
    i := 1
    fmt.Println("initial:", i)
    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)
    // ＆i语法给出i的存储器地址，即指向i的指针。

    fmt.Println("pointer:", &i)
}

// $ go run 17-pointers.go
// initial: 1
// zeroval: 1
// zeroptr: 0
// pointer: 0x42131100
