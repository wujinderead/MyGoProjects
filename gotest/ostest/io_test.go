package ostest

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

const maxread = 12

var (
	closedErr     = errors.New("already closed")
	bufferFullErr = errors.New("buffer fulled")
	noContentErr  = errors.New("no content")
	seekErr       = errors.New("seek out of content")
)

func TestIo(t *testing.T) {
	// Read
	buf := newBuffer(50, "赵客缦胡缨，吴钩霜雪明.")
	p := make([]byte, 6)
	for i := 0; i < 8; i++ {
		if i == 7 {
			_ = buf.Close()
		}
		n, err := buf.Read(p)
		fmt.Println("Read:", err, n, string(p[:n]))
	}
	fmt.Println()

	// Write
	buf.resetRead()
	p = []byte("abcdef")
	for i := 0; i < 4; i++ {
		n, err := buf.Write(p)
		fmt.Println("Write:", err, n, string(p[:n]))
	}
	fmt.Println()

	// ReadAll
	buf.resetRead()
	all, err := ioutil.ReadAll(buf)
	fmt.Println("ReadAll:", err, string(all), len(all)) // read buf until EOF or err
	fmt.Println()

	// ReadFull
	buf.resetRead()
	p = make([]byte, 30)
	n, err := io.ReadFull(buf, p)
	fmt.Println("ReadFull:", err, n, string(p[:n])) // read until p is full
	p = make([]byte, 60)
	n, err = io.ReadFull(buf, p)
	fmt.Println("ReadFull:", err, n, string(p[:n])) // it will read continuously until full or encounter error

	// ReadAtLeast
	buf.resetRead()
	p = make([]byte, 60)
	n, err = io.ReadAtLeast(buf, p, 60)
	fmt.Println("ReadAtLeast:", err, n, string(p[:n])) // it will read continuously until read enough or encouter error
	fmt.Println()

	// WriteString, StringWriter
	buf.reset()
	n, err = io.WriteString(buf, "银鞍照白马，飒沓如流星。")
	fmt.Println("WriteString:", err, n)
	n, err = buf.WriteString("十步杀一人，千里不留行。")
	fmt.Println("StringWriter:", err, n)
	fmt.Println()

	// Seek and Read
	p = make([]byte, 12)
	n, err = buf.Read(p) // read from offset 0
	fmt.Println("Read:", err, n, string(p[:n]))
	s, err := buf.Seek(42, io.SeekStart)
	fmt.Println("Seek:", err, s)
	n, err = buf.Read(p) // seek to offset 42 and reads
	fmt.Println("Read:", err, n, string(p[:n]))
	fmt.Println()

	// todo io.copy

}

func newBuffer(size int, init string) *buffer {
	buf := new(buffer)
	buf.closed = false
	buf.roff = 0
	buf.size = size
	buf.buf = make([]byte, size)
	copy(buf.buf, []byte(init))
	buf.woff = min(len(init), size)
	return buf
}

type buffer struct {
	closed bool
	roff   int
	woff   int
	size   int
	buf    []byte
}

func (b *buffer) Close() error {
	b.closed = true
	return nil
}

func (b *buffer) reset() {
	b.closed = false
	b.roff = 0
	b.woff = 0
}

func (b *buffer) resetRead() {
	b.closed = false
	b.roff = 0
}

func (b *buffer) Read(p []byte) (int, error) {
	if b.closed || b.roff == b.size { // closed or read to end, return EOF
		return 0, io.EOF
	}
	if b.roff >= b.woff { // no content, return nil err
		return 0, noContentErr
	}
	canRead := min(min(b.woff-b.roff, maxread), len(p))
	copy(p, b.buf[b.roff:b.roff+canRead])
	b.roff += canRead
	return canRead, nil
}

func (b *buffer) Write(p []byte) (canWrite int, err error) {
	if b.closed {
		return 0, closedErr
	}
	canWrite = min(b.size-b.woff, len(p))
	copy(b.buf[b.woff:], p)
	b.woff += canWrite
	if canWrite < len(p) {
		err = bufferFullErr
	}
	return
}

// seek read pointers
func (b *buffer) Seek(offset int64, whence int) (int64, error) {
	// ignore whence, always seek from start
	if offset >= int64(b.woff) {
		return 0, seekErr
	}
	b.roff = int(offset)
	return offset, nil
}

func (b *buffer) WriteString(str string) (int, error) {
	if b.closed {
		return 0, closedErr
	}
	return b.Write([]byte(str))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
