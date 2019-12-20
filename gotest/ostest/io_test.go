package ostest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"
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

	// Copy
	buf.resetRead()
	w := newBuffer(30, "")
	written, err := io.Copy(w, buf) // copy from read to writer until EOF or error
	fmt.Println("Copy:", err, written)
	fmt.Println()

	// CopyN
	buf.resetRead()
	w.reset()
	written, err = io.CopyN(w, buf, 20) // copy from read to writer until enough bytes or error
	fmt.Println("CopyN:", err, written)
	fmt.Println()

	// CopyBuffer
	buf.resetRead()
	w.reset()
	p = make([]byte, 10)
	written, err = io.CopyBuffer(w, buf, p) // like io.Copy but use the provided buffer other than allocating by go
	fmt.Println("CopyBuffer:", err, written)
	fmt.Println()

	// TeeReader
	buf.resetRead()
	w.reset()
	tee := io.TeeReader(buf, w) // when read from r, also write to w
	p = make([]byte, 12)
	for i := 0; i < 6; i++ {
		n, err := tee.Read(p)
		// when w full, return writer error.
		// however the read can still continue, until reader return EOF.
		fmt.Println("tee read:", err, n, string(p[:n]))
		fmt.Println("w buf:", string(w.buf[:w.woff]))
		fmt.Println("r.roff:", buf.roff)
	}
	fmt.Println()

	// MultiReader, MultiWriter
	r1, r2 := newBuffer(10, "abcdefghij"), newBuffer(10, "klmnopqrst")
	w1, w2, w3 := newBuffer(30, ""), newBuffer(20, ""), newBuffer(20, "")
	mr := io.MultiReader(r1, r2)
	mw := io.MultiWriter(w1, w2, w2)
	p = make([]byte, 24)
	n, err = mr.Read(p)
	fmt.Println("MultiReader:", err, n, string(p))
	n, err = mr.Read(p[n:])
	fmt.Println("MultiReader:", err, n, string(p)) // read readers serially until all reader EOF

	n, err = mw.Write(p)                              // write to all writers, if one error, the others won't continue
	fmt.Println("MultiWriter:", err, n)               // w2 error, return w2's n and err
	fmt.Println("Writer1:", string(w1.buf[:w1.woff])) // w1 success
	fmt.Println("Writer2:", string(w2.buf[:w2.woff])) // w2 return err
	fmt.Println("Writer3:", string(w3.buf[:w3.woff])) // won't continue to write w3
	fmt.Println()
}

func TestIo1(t *testing.T) {
	// Discard
	r := newBuffer(20, "1234567890")
	n, err := io.Copy(ioutil.Discard, r) // read from r and discard the content, only the read offset changes
	fmt.Println("Discard:", err, n)

	// Pipe
	pr, pw := io.Pipe()
	// linux syscall pipe() is used for interprocess communication that the data written to the write end
	// of the pipe is buffered by the kernel until it is read from the read end of the pipe.
	// golang io.Pipe() is like the syscall pipe() but for communication for goroutines.
	// there can be multiple readers and writers concurrently, reader will block when nothing to read,
	// writer will block until all written data are read. io.Pipe() do not buffer the data.
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		p := make([]byte, 8)
		n, err := pr.Read(p)
		fmt.Println("pr1 read:", err, n, string(p[:n]))
		wg.Done()
	}()
	go func() {
		p := make([]byte, 6)
		n, err := pr.Read(p)
		fmt.Println("pr2 read:", err, n, string(p[:n]))
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		p := make([]byte, 8)
		n, err := pr.Read(p)
		fmt.Println("pr3 read:", err, n, string(p[:n]))
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		n, err := pw.Write([]byte("1234567890qwertyuiop"))
		fmt.Println("pw write:", err, n)
		wg.Done()
	}()
	wg.Wait()
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

func TestCopyReadFromWriteTo(t *testing.T) {
	w := &wt{os.Stderr}
	r := &rd{bytes.NewReader([]byte("bbbbbbb\n"))}
	fmt.Printf("w: %p, r: %p\n", w, r)

	// w ReadFrom reader (bytes.Reader) and write to w.w (os.Stderr)
	n, err := w.ReadFrom(bytes.NewReader([]byte("aaaaaaa\n")))
	fmt.Println(n, err)

	// r read r.r (bytes.Reader) and WriteTo writer (os.Stderr)
	n, err = r.WriteTo(os.Stdout)
	fmt.Println(n, err)
	fmt.Println()

	w = &wt{os.Stderr}
	r = &rd{bytes.NewReader([]byte("ccccccc\n"))}
	n, err = io.Copy(w, r) // call r.WriteTo(w) first
	fmt.Println(n, err)
	fmt.Println()

	f, _ := os.Open("/proc/sys/kernel/pid_max")
	defer f.Close()
	n, err = io.Copy(w, f) // no r.WriteTo, call w.ReadFrom(r) instead
	fmt.Println(n, err)
	_, _ = f.Seek(0, 0)
	fmt.Println()

	n, err = io.Copy(os.Stdout, f) // just call Read and Write
	fmt.Println(n, err)
}

type wt struct {
	w io.Writer
}

type rd struct {
	r io.Reader
}

// the semantic of io.ReaderFrom interface:
// read from r (then process data implicitly, e.g. write the data to somewhere). return number read and error.
func (w *wt) ReadFrom(r io.Reader) (n int64, err error) {
	fmt.Printf("%p call ReadFrom\n", w)
	buf := make([]byte, 30)
	nn, err := r.Read(buf)
	_, _ = w.w.Write(buf[:nn])
	return int64(nn), err
}

func (w *wt) Write(b []byte) (n int, err error) {
	return w.w.Write(b)
}

// the semantic of io.WriterTo interface:
// (get some data implicitly, e.g. read data from somewhere) and write to w. return the number written and error.
func (r *rd) WriteTo(w io.Writer) (n int64, err error) {
	fmt.Printf("%p call WriteTo\n", r)
	buf := make([]byte, 30)
	nn, err := r.r.Read(buf)
	_, _ = w.Write(buf[:nn])
	return int64(nn), err
}

func (r *rd) Read(b []byte) (n int, err error) {
	return r.r.Read(b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
