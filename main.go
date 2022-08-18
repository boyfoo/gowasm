package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("myname", "shenyi")
	<-make(chan struct{})
}
