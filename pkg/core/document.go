package core

import (
	"syscall/js"
)

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
//这个文件包含的是 html\js相关的简单封装
var Doc = js.Global().Get("document")
var Window = js.Global().Get("window")
var Console = js.Global().Get("console")

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func Log(message string) {
	Console.Call("log", message)
}
func Alert(message string) {
	Window.Call("alert", message)
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func AddEventListener(n js.Value, eventType string, fn js.Func) {

	n.Call("addEventListener", eventType, fn)
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func GetByID(id string) js.Value {
	ret := Doc.Call("getElementById", id)
	return ret
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
// 针对TextNode的处理
//  <span>xxxoo</span>
func CreateTextNode(data string) js.Value {
	ret := Doc.Call("createTextNode", data)
	return ret
}

// 本课程来 自 程序员在囧途(www.jtthink.com) 咨询群：98514334
// <input
func CreateElement(name string) js.Value {
	ret := Doc.Call("createElement", name)
	return ret
}

// 本课程来自 程 序员在囧途(www.jtthink.com) 咨询群：98514334
func AppendChild(p js.Value, child js.Value) {
	p.Call("appendChild", child)
}
