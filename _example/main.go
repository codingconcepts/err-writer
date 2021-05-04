package main

import (
	"bytes"
	"fmt"

	writer "github.com/codingconcepts/err-writer"
)

func main() {
	buf := &bytes.Buffer{}

	w := writer.New(buf)
	w.Write([]byte("a\n"))
	w.Writef("%s\n", "b")
	w.Writeln([]byte("c"))

	fmt.Println(buf.String())
}
