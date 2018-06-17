// 行过滤器
// 一个行过滤器 在读取标准输入流的输入，处理该输入，然后将得到一些的结果输出到标准输出的程序中是常见的一个功能。grep 和 sed 是常见的行过滤器。

// 这里是一个使用 Go 编写的行过滤器示例，它将所有的输入文字转化为大写的版本。你可以使用这个模式来写一个你自己的 Go行过滤器。
package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
func main() {
// 对 os.Stdin 使用一个带缓冲的 scanner，让我们可以直接使用方便的 Scan 方法来直接读取一行，每次调用该方法可以让 scanner 读取下一行。

    scanner := bufio.NewScanner(os.Stdin)
// Text 返回当前的 token，现在是输入的下一行。

    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
// 写出大写的行。

        fmt.Println(ucl)
    }
// 检查 Scan 的错误。文件结束符是可以接受的，并且不会被 Scan 当作一个错误。

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
// 试一下我们的行过滤器，首先创建多一个有小写行的文件。

// $ echo 'hello'   > /tmp/lines
// $ echo 'filter' >> /tmp/lines
// 然后使用行过滤器来得到大些的行。
//
// $ cat /tmp/lines | go run line-filters.go
// HELLO
// FILTER
