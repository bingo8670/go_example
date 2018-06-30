## [Go by Example](https://books.studygolang.com/gobyexample/)
### 命令集
- go run xxx.go 执行xxx脚本；
- go build xxx.go 将xxx脚本转换为二进制文件；输入文件路径即可执行；
- gofmt -w \*.go 会格式化并重写所有 Go 源文件；

### 术语
- goroutine：协程


### 备忘
- for是Go唯一的循环构造。
- go脚本输出信息顺序不固定，高并发；
- 在使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...] 的格式显示
- 要创建一个空 map，需要使用内建的 make:make(map[key-type]val-type).
- := 语句是申明并初始化变量的简写，例如var f string = "short" 等效于f := "short"
- \_, 表示空标识符，避免了必须为返回值声明所有变量，用于返回特定变量值；

#### 注释方法
- 单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。
- 多行注释也叫块注释，均已以 /* 开头，并以 \*/ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。
