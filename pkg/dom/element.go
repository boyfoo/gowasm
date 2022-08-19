package dom

import (
	"golang.org/x/net/html"
	"syscall/js"
)

type Element struct {
	Parent  *Element
	Node    *html.Node
	JsValue js.Value
}

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

func NewElementByValue(v js.Value) *Element {
	return &Element{JsValue: v}
}
