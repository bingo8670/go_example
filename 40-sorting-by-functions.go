// 有时候我们会想通过非自然顺序排序集合。例如，假设我们想按字符串的长度对字符串进行排序，而不是按字母顺序排列。以下是Go中的自定义排序示例。
package main
import "sort"
import "fmt"

// 为了在 Go 中使用自定义函数进行排序，我们需要一个对应的类型。这里我们创建一个为内置 []string 类型的别名的ByLength 类型，
type byLength []string

// 我们在类型中实现了 sort.Interface 的 Len，Less和 Swap 方法，这样我们就可以使用 sort 包的通用Sort 方法了，Len 和 Swap 通常在各个类型中都差不多，Less 将控制实际的自定义排序逻辑。在我们的例子中，我们想按字符串长度增加的顺序来排序，所以这里使用了 len(s[i]) 和 len(s[j])。
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// 一切都准备好了，我们现在可以通过将原始的 fruits 切片转型成 ByLength 来实现我们的自定排序了。然后对这个转型的切片使用 sort.Sort 方法。
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    // 通过调用sort.Sort自定义的排序方法
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
