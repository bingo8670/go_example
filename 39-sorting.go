// Go的排序包实现了内置和用户定义类型的排序。先来看看为内置的排序。
package main
import "fmt"
import "sort"
func main() {

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}

// $ go run 39-sorting.go
// Strings: [a b c]
// Ints:    [2 4 7]
// Sorted:  true
