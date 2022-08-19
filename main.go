package main

import (
	"fmt"
	"golang.org/x/net/html"
	"gowasm/pkg/dom"
	"strings"
)

var forgetElement = []string{"html", "head", "body"}

func isForget(data string) bool {
	for _, s := range forgetElement {
		if s == data {
			return true
		}
	}
	return false
}

const htmlStr = `
<div>
    <span class="teststyle">jtthink.com</span>
</div>
`

func main() {
	jue, _ := html.Parse(strings.NewReader(htmlStr))
	app := dom.NewElementByValue(dom.GetByID("app"))
	var f func(n *html.Node, parent *dom.Element)
	f = func(n *html.Node, parent *dom.Element) {
		var p *dom.Element

		if n.Type == html.TextNode {
			p = dom.NewElement(n).Create()
		}

		if (n.Type == html.ElementNode || n.Type == html.TextNode) && strings.Trim(n.Data, " ") != "" && !isForget(n.Data) {

			p = dom.NewElement(n).Create()

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

//go:export test
func Test() {
	fmt.Println("test")
}
