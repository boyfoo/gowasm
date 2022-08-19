package main

import "fmt"

func main() {

}

//go:export test
func Test() {
	fmt.Println("我是你大哥")
}
