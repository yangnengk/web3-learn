package main

import (
	"fmt"
	"github.com/learn/init_compile/foo"
)

/*
在 go build -tags "tag1" 时，会编译出 func_tag1.go 文件。
在 go build -tags "tag2" 时，会编译出 func_tag2.go 文件。
在 go build -tags "tag3" 时，会编译出 func_tag3.go 文件。
在 go build 时，会编译出 func_default.go 文件。
------------------------------------------------
PS G:\go_home\go_workspace\src\web3-learn\web3-learn-01\编译成可执行文件\init_compile> go build -tags "tag1 tag2"
PS G:\go_home\go_workspace\src\web3-learn\web3-learn-01\编译成可执行文件\init_compile> .\init_compile.exe
Tag:  [Default Tag1 Tag2]
PS G:\go_home\go_workspace\src\web3-learn\web3-learn-01\编译成可执行文件\init_compile> go build -tags "tag1 tag2 tag3"
PS G:\go_home\go_workspace\src\web3-learn\web3-learn-01\编译成可执行文件\init_compile> .\init_compile.exe
Tag:  [Default Tag1 Tag2]
PS G:\go_home\go_workspace\src\web3-learn\web3-learn-01\编译成可执行文件\init_compile> go build -tags "tag1 tag3"
PS G:\go_home\go_workspace\src\web3-learn\web3-learn-01\编译成可执行文件\init_compile> .\init_compile.exe
Tag:  [Default Tag1 Tag3]
*/
func main() {
	//foo.PlatformSpecificFunction()
	fmt.Println("Tag: ", foo.TAG)
}
