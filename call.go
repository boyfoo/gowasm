package main

import (
	"github.com/wasmerio/wasmer-go/wasmer"
	"io/ioutil"
	"log"
)

func main() {
	wasmBytes, err := ioutil.ReadFile("main.wasm")
	if err != nil {
		log.Fatalln(err)
	}
	store := wasmer.NewStore(wasmer.NewEngine())

	// 编译wasm字节模块
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		log.Fatalln(err)
	}
	// 创建运行环境
	wasiEnv, err := wasmer.NewWasiStateBuilder("wasi-program").Finalize()
	if err != nil {
		log.Fatalln(err)
	}
	object, err := wasiEnv.GenerateImportObject(store, module)
	if err != nil {
		log.Fatalln(err)
	}
	instance, err := wasmer.NewInstance(module, object)
	if err != nil {
		log.Fatalln(err)
	}
	// 调用
	test, err := instance.Exports.GetFunction("test")
	if err != nil {
		log.Fatalln(err)
	}
	test()
}
