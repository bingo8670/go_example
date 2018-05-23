// 接口 是方法特征的命名集合。
package main
import "fmt"
import "math"

// 这里是一个几何体的基本接口。
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

// 如果一个变量的是接口类型，那么我们可以调用这个被命名的接口中的方法。这里有一个一通用的 measure 函数，利用这个特性，它可以用在任何 geometry 上。
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

// 结构体类型 circle 和 rect 都实现了 geometry接口，所以我们可以使用它们的实例作为 measure 的参数。
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
