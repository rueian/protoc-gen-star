package main

import (
	"github.com/rueian/protoc-gen-star"
	"github.com/rueian/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		ASTPrinter(),
		JSONify(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
