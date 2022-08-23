package components

import "mywasm/pkg/core"

var TestComp = &core.Component{
	Template: `
<div>
	<span style="color:red">hello jtthink</span>
	<input type="button" value="点我" onclick="jtthink()" />
</div>
`,
}
