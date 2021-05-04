package writer

import (
	"bytes"
	"fmt"
	"io"
)

// Writer keeps track of errors and turns write operations into no-ops in the
// event that an error occurs, prevening you from having to check for errors
// on every write.
type Writer struct {
	w   io.Writer
	err error
}

// New returns a pointer to a new instance of Writer.
func New(w io.Writer) *Writer {
	return &Writer{
		w: w,
	}
}

// Write writes a byte slice to the underlying buffer, ignoring the write
// operation if an error occurred in previous writes.
func (w *Writer) Write(buf []byte) {
	if w.err != nil {
		return
	}

	_, w.err = w.w.Write(buf)
}

// Writeln writes a byte slice to the underlying buffer as a new line, ignoring the write
// operation if an error occurred in previous writes.
func (w *Writer) Writeln(buf []byte) {
	if w.err != nil {
		return
	}

	if !bytes.HasSuffix(buf, []byte("\n")) {
		buf = append(buf, byte('\n'))
	}

	_, w.err = w.w.Write(buf)
}

// Writef writes a formatted string to the underlying buffer, ignoring the write
// operation if an error occurred in previous writes.
func (w *Writer) Writef(format string, a ...interface{}) {
	if w.err != nil {
		return
	}

	b := []byte(fmt.Sprintf(format, a...))
	_, w.err = w.w.Write(b)
}
