// 接口被命名为方法签名的集合。
package main
import "fmt"
import "math"

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
// 如果变量具有接口类型，那么我们可以调用指定接口中的方法。这是一个通用的度量函数，利用它可以处理任何几何图形。

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}

// $ go run 20-interfaces.go
// {3 4}
// 12
// 14
// {5}
// 78.53981633974483
// 31.41592653589793
