package main

import (
	"github.com/bytecodealliance/wasmtime-go"
)

func main() {
	// Almost all operations in wasmtime require a contextual `store`
	// argument to share, so create that first
	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModuleFromFile(store.Engine, "./python/bin/python.wasm")
	check(err)

	config := wasmtime.NewWasiConfig()
	config.InheritStderr()
	config.InheritStdout()
	config.InheritStdin()
	config.SetArgv([]string{""})
	check(config.PreopenDir("python/lib", "lib"))
	config.SetEnv([]string{
		"PYTHONPATH"}, []string{""})

	wasi, err := wasmtime.NewWasiInstance(store, config, "wasi_unstable")
	check(err)

	linker := wasmtime.NewLinker(store)
	err = linker.DefineWasi(wasi)
	if err != nil {
		panic(err)
	}

	// Next up we instantiate a module which is where we link in all our
	// imports. We've got one import so we pass that in here.
	instance, err := linker.Instantiate(module)
	check(err)
	// After we've instantiated we can lookup our `run` function and call
	// it.was
	// for _, export := range instance.Exports() {
	// 	spew.Dump(export.Type().FuncType().)
	// }
	run := instance.GetExport("_start").Func()
	_, err = run.Call()
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
