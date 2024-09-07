package asm

import (
	_ "embed"
	"fmt"
	. "strconv"
	. "strings"
	"unicode"
)

var (
	dataCounter int = 0
	//go:embed assembly/data.s
	data string
	//go:embed assembly/init.s
	Init string
	//go:embed assembly/exit.s
	exit string
	//go:embed assembly/write.s
	write string
)

func arg(n int) string {
	return fmt.Sprintf("${%d}", n)
}

func GenWordData(word string) string {
	buf := data
	wordids := []rune(word)
	for i, char := range wordids {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			continue
		}
		wordids[i] = '_'
	}
	wordid := string(wordids[:]) + Itoa(dataCounter)
	wordidl := wordid
	wordidl += "l"
	buf = ReplaceAll(buf, arg(0), wordid)
	buf = ReplaceAll(buf, arg(1), word)
	if ContainsAny(word, "\n") {
		buf = ReplaceAll(buf, "\n", "")
		buf = ReplaceAll(buf, arg(2), ", 0xA")
	} else {
		buf = ReplaceAll(buf, arg(2), "")
	}
	buf = ReplaceAll(buf, arg(3), wordidl)
	buf = ReplaceAll(buf, arg(4), wordid)
	dataCounter++
	return buf
}

func GenExit(status int) string {
	buf := ReplaceAll(exit, arg(0), Itoa(status))
	return buf
}

func GenWrite(word string) string {
	buf := write
	buf = ReplaceAll(buf, arg(0), "1")
	buf = ReplaceAll(buf, arg(1), "lisp0")
	buf = ReplaceAll(buf, arg(2), "lisp0l")
	return buf
}
