// Go支持在结构类型上定义的方法。
package main
import "fmt"
type rect struct {
    width, height int
}

// 这里的 area 方法有一个接收器类型 rect。
func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

// Go 自动处理方法调用时的值和指针之间的转化。你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}

// $ go run 19-methods.go
// area:  50
// perim: 30
// area:  50
// perim: 30
