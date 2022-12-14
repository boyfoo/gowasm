package core

import (
	"fmt"
	"golang.org/x/net/html"
	"mywasm/pkg/utils"
	"strings"
	"syscall/js"
)

type Method func(this js.Value, args []js.Value) interface{}
type Methods map[string]Method

var NotFoundMethod = fmt.Errorf("No such method")

func (ms Methods) GetMethod(key string) (js.Func, error) {
	if getf, ok := ms[key]; ok {
		return js.FuncOf(getf), nil
	}
	return js.Func{}, NotFoundMethod
}

// 组件原型
type Component struct {
	Template string
	Methods  Methods
}

func (c *Component) GetTemplate() string {

	return c.Template
}

// 这个就是 前面几课时
func (c *Component) MountTo(id string) {
	//做一些 初始化判断
	if c.Methods == nil {
		c.Methods = map[string]Method{}
	}

	jue, _ := html.Parse(strings.NewReader(c.GetTemplate()))
	app := NewElementByValue(GetByID(id))
	var f func(n *html.Node, parent *Element)
	f = func(n *html.Node, parent *Element) {
		var p *Element
		//是否是我们要处理的可用节点
		if utils.IsActiveNode(n) {
			p = NewElement(n).Create(c.Methods)
			if parent == nil || parent.JsValue.IsUndefined() {
				parent = app
			}
			p.AppendTo(parent) // 把当前元素 加入到 父元素(DOM)中

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, p)
		}
	}

	f(jue, app) // 递归解析html 元素
}
