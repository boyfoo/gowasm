package dom

import "syscall/js"

var Doc = js.Global().Get("document")

func GetByID(id string) js.Value {
	return Doc.Call("getElementById", id)
}

func CreateTextNode(data string) js.Value {
	return Doc.Call("createTextNode", data)
}

func CreateElement(name string) js.Value {
	return Doc.Call("createElement", name)
}

func AppendChild(p js.Value, child js.Value) {
	p.Call("appendChild", child)
}
