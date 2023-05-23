package mrpc

import (
	"fmt"
	"package/mcom"
)

// mrpc 包

type RPC struct {
	Api mcom.CApi
}

func (r *RPC) Start() {
	// 实现 RPC 版本的函数逻辑
	fmt.Println("rpc start")
	r.GetRpcVersion()
	r.Api.GetApiVersion()
}
func (r *RPC) GetRpcVersion() {
	fmt.Println("rpc version 2.0")
}
