package main

import "syscall/js"

func main() {
	doc := js.Global().Get("document")
	app := doc.Call("getElementById", "app")
	if !app.IsNull() {
		app.Set("innerHTML", "hello zxzx")
	}
}
