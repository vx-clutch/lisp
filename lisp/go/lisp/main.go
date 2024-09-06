package main

import (
	"lisp/lisp/go/asm"
	"lisp/lisp/go/cli"
	"lisp/lisp/go/compiler
	"lisp/lisp/go/types"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		// exits
		cli.Fatal("No files found")
	}
	if strings.Contains(os.Args[1], ".lisp") == false || !strings.Contains(os.Args[1], ".cl") == false {
		cli.Fatal("Incorrect file extention")
	}
	test()
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(content)
	compiler.Compile(input)
}

func test() {
	var asm types.Assembly
	init := asm.GenStartAsm("test.lisp")
	write := asm.GenWriteAsm(1, "mayo", 4)
	asm.Asm = append(asm.Asm, init)
	asm.Asm = append(asm.Asm, write)
	cli.Debug("asm", asm.Reduce())
}
