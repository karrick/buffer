package buffer

import "io"

// Credit to: http://dave.cheney.net/2015/06/05/friday-pop-quiz-the-smallest-buffer

// A Buffer is a variable-sized buffer of bytes with Read and Write
// methods. The zero value for Buffer is an empty buffer ready to use.
type Buffer []byte

// Write writes len(p) bytes from p to the Buffer.
func (b *Buffer) Write(p []byte) (int, error) {
	*b = append(*b, p...)
	return len(p), nil
}

// Read reads up to len(p) bytes into p from the Buffer.
func (b *Buffer) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(*b) == 0 {
		return 0, io.EOF
	}
	n := copy(p, *b)
	*b = (*b)[n:]
	return n, nil
}
