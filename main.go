package main

import (
	"fmt"
	"mywasm/pkg/components"
	"syscall/js"
)

func main() {
	js.Global().Set("jtthink", js.FuncOf(func(this js.Value, args []js.Value) any {
		js.Global().Get("window").Call("alert", "弹窗了")
		return nil
	}))
	components.TestComp.MountTo("app")

	<-make(chan struct{})
}

//go:export mytest
func Test() int {
	fmt.Println("test")
	return 123
}
