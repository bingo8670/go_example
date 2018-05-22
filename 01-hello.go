package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}

// $ go run 01-hello-world.go
// hello world
// $ go build hello-world.go
// $ ./hello-world
// hello world

// go build hello.go 将hello脚本转换为二进制文件；输入文件路径 ./hello-world 即可执行；
