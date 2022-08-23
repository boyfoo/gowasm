# go wasm 学习笔记

## wasm 形式 1

#### js文件

下载`go`官方提供的调用`js` `https://github.com/golang/go/blob/go1.15.15/misc/wasm/wasm_exec.js`

#### 编译

`.\build.bat`

#### 运行

浏览`htmls`下的`index.html`

## 支持 wasi 形式 2.0

`go`要满足版本要求使用`1.18`

`go`原本编译的`wasm`只能在浏览器环境中运行，要让所有语言在任何环境运行使用此方案 `https://github.com/wasmerio/wasmer/releases/tag/2.2.1`

`windows`系统安装`$v="2.2.1"; iwr https://win.wasmer.io -useb | iex`

安装`tinygo`编译器，编译`wasi` `https://github.com/tinygo-org/tinygo/releases/tag/v0.23.0`

设置解压目录地址为环境变量`TINYGOROOT`,其中`bin`目录加入`PATH`中

还需要一个工具链库`https://github.com/WebAssembly/binaryen/releases/tag/version_108`

解压后将`bin`目录配置至`PATH`上

#### 编译

使用`tinygo`编译，在`go`代码中使用的包`https://tinygo.org/docs/reference/lang-support/stdlib/` 需要查看文档是否支持编译，两个√才满足生产环境使用要求，否则不要使用

编译命令`tinygo build -wasm-abi=generic -target=wasi -o main.wasm main.go`

运行`wasmer --dir . main.wasm`

## go 调用 wasi 3.0

需要使用`cgo`所以要用`linux`系统

修改了`main.go`文件，重新在`windows`编译：`tinygo build -scheduler=none -target=wasi -o main.wasm main.go`

将文件同步到`linux`系统中，运行`go run call.go`

## go 修改html 4.0

修改`main.go`h和`html`代码，使用`go`代码修改元素

原本可以使用`go`直接遍历，但是`go`编译后的包太大，所以依旧使用`tinygo`，但缺点是有的功能不支持

`go`之前运行要使用`wasm_exec.js`，`tinygo`编译的也是一样，但是他有他自己的`js`文件

下载`js`文件， `https://github.com/tinygo-org/tinygo/blob/v0.23.0/targets/wasm_exec.js`

## go 仿照 vue 渲染html 5-6

## go 仿照 vue 组件化 7.0

`tinygobuild.bat` 编译后，打开html

## go 仿照 vue 调用函数 8.0 (简单方式)

`components/test.go`内新增`onclick="jtthink()"`新增点击事件函数

在`main.go`定义函数

## go 仿照 vue 调用函数 8.1 (导出函数)

修改`index.html`与`main.go` 重新编译后，点击页面新按钮