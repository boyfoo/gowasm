package core

import (
	"golang.org/x/net/html"
	"syscall/js"
)

//这个文件 针对 html.Node做封装
// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
type Element struct {
	Parent  *Element
	Node    *html.Node
	JsValue js.Value
}

// 创建元素， 必须要执行Create才会真的创建。链式调用 。返回自己。  同时设置属性
// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func (e *Element) Create() *Element {
	if e.Node.Type != html.TextNode {
		n := CreateElement(e.Node.Data)
		for _, arr := range e.Node.Attr {
			n.Call("setAttribute", arr.Key, arr.Val)
		}
		e.JsValue = n
	} else {
		n := CreateTextNode(e.Node.Data)
		e.JsValue = n
	}

	return e
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func (e *Element) AppendTo(parent *Element) *Element {
	if e.JsValue.IsUndefined() {
		return e
	}
	AppendChild(parent.JsValue, e.JsValue)
	e.Parent = parent
	return e
}

func NewElement(node *html.Node) *Element {
	return &Element{Node: node, JsValue: js.Undefined()}
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
//一般用于 跟节点 ，如app
func NewElementByValue(v js.Value) *Element {
	return &Element{JsValue: v}
}
