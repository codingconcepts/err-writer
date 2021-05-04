# err-writer
A simple monad writer based on Rob Pike's https://blog.golang.org/errors-are-values.

## Installation

```
$ go get -u github.com/codingconcepts/err-writer
```

## Usage

``` go
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
    
    // Outputs:
    // a
    // b
    // c
}
```