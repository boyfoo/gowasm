package components

import (
	"mywasm/pkg/core"
	"syscall/js"
)

//业务 组件 ---模仿Vue 的。
// 只需要直接 在Go里写struct ，就能和Vue一样
var TestComp *core.Component = &core.Component{
	Template: `
	<div>
	   <span style="color:red">hello jtthink</span>
	   <input type="button" value="点我" @click="test" />
	</div>
`,
	Methods: core.Methods{
		"test": func(this js.Value, args []js.Value) interface{} {
			core.Alert("this os test (@click)")
			return nil
		},
	},
}
