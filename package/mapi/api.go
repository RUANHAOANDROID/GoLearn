package mapi

import (
	"fmt"
	"package/mcom"
)

type API struct {
	// ...
	Rpc mcom.CRpc
}

func (a *API) Start() {
	fmt.Println("mapi start")
	a.GetApiVersion()
	a.Rpc.GetRpcVersion()
}
func (a *API) GetApiVersion() {
	fmt.Println("mapi version = 1.0")
}
