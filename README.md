# go wasm 学习笔记

## wasm 形式

#### js文件

下载`go`官方提供的调用`js` `https://github.com/golang/go/blob/go1.15.15/misc/wasm/wasm_exec.js`

#### 编译

`.\build.bat`

#### 运行

浏览`htmls`下的`index.html`

## 支持 wasi 形式

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