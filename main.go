package main

import (
	"fmt"
	"mywasm/pkg/components"
)

func main() {
	components.TestComp.MountTo("app")
}

//go:export test
func Test() {
	fmt.Println("test")
}
