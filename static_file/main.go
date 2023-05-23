package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 获取可执行文件的路径
	exePath, err := os.Executable()
	if err != nil {
		// 处理获取可执行文件路径失败的错误
		os.Exit(1)
	}

	// 获取可执行文件所在的目录
	root := filepath.Dir(exePath)
	fmt.Println("DIR=" + root)
	// 创建文件服务器处理器
	fileServer := http.FileServer(http.Dir(root))

	// 注册文件服务器处理器，并将文件路径前缀设置为空
	http.Handle("/", fileServer)

	// 启动服务器并监听指定的端口
	port := ":8000"
	fmt.Println(fmt.Sprintf("%s%s", "http://127.0.0.1", port))
	err = http.ListenAndServe(port, nil)
	if err != nil {
		// 处理启动服务器失败的错误
		os.Exit(1)
	}
}
