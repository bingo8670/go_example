// 有时候我们会想通过非自然顺序排序集合。例如，假设我们想按字符串的长度对字符串进行排序，而不是按字母顺序排列。以下是Go中的自定义排序示例。
package main
import "sort"
import "fmt"

type byLength []string

func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    // 通过调用sort.Sort自定义的排序方法
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
