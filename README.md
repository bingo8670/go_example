## Go by Example
### 命令集
- go run xxx.go 执行xxx脚本；
- go build xxx.go 将xxx脚本转换为二进制文件；输入文件路径即可执行；
- gofmt -w \*.go 会格式化并重写所有 Go 源文件；

### 注释方法
- 单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。
- 多行注释也叫块注释，均已以 /* 开头，并以 \*/ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。


### 备忘
- for是Go唯一的循环构造。
- go脚本输出信息顺序不固定，高并发；
