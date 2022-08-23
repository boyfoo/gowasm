package utils

import (
	"golang.org/x/net/html"
	"strings"
)

var forgetElement = []string{"html", "head", "body"}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func isForget(data string) bool {
	for _, s := range forgetElement {
		if s == data {
			return true
		}
	}
	return false
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func IsActiveNode(n *html.Node) bool {
	if (n.Type == html.ElementNode || n.Type == html.TextNode) && strings.Trim(n.Data, " ") != "" && !isForget(n.Data) {
		return true
	}
	return false
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
