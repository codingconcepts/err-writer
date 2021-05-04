package writer

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	buf := &bytes.Buffer{}
	w := Writer{w: buf}

	w.Write([]byte("a"))
	w.Write([]byte("b"))
	w.err = errors.New("something went wrong")
	w.Write([]byte("c"))

	assert.Equal(t, "ab", buf.String())
}

func TestWriteln(t *testing.T) {
	buf := &bytes.Buffer{}
	w := Writer{w: buf}

	w.Writeln([]byte("a"))
	w.Writeln([]byte("b"))
	w.err = errors.New("something went wrong")
	w.Writeln([]byte("c"))

	assert.Equal(t, "a\nb\n", buf.String())
}

func TestWritef(t *testing.T) {
	buf := &bytes.Buffer{}
	w := Writer{w: buf}

	w.Writef("%s", "a")
	w.Writef("%s", "b")
	w.err = errors.New("something went wrong")
	w.Writef("%s", "c")

	assert.Equal(t, "a b", buf.String())
}
