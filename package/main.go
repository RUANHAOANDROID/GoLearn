package main

import (
	"fmt"
	"package/mapi"
	"package/mrpc"
)

func main() {
	//使用new 关键字动态分配指定类型的nil，并返回该类型的指针
	api := new(mapi.API)
	rpc := new(mrpc.RPC)
	//赋值
	api.Rpc = rpc
	rpc.Api = api
	//调用api
	api.Start()
	fmt.Println("------------")
	//调用rpc
	rpc.Start()
}
