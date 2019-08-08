package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"syscall"
	"time"
	"unsafe"
)

const flockRetryTimeout = 50 * time.Millisecond
const maxMapSize = 50 * time.Millisecond

func main() {
	//testMmapReadonly()
	//testMmapWritable()
	//testMmapReadonlyWrite()
	testMmapNewFile()
}

func testMmapReadonly() {
	filename := "/home/xzy/t_bz_ga_fz_zfba_xryglaj.csv"
	// open the file
	file, err := os.OpenFile(filename, os.O_RDONLY, os.FileMode(0644))
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("file close err:", err)
		}
	}()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	fmt.Println("file path:", file.Name(), "file descriptor:", file.Fd())

	// acquire file lock, exclusive=true if rdwr, exclusive=false if readonly
	if err := flock(file, false, 100*time.Millisecond); err != nil {
		return
	}
	defer func() {
		err := funlock(file)
		if err != nil {
			fmt.Println("funlock err:", err)
		}
	}()

	// get file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("stat err:", err)
		return
	}
	size := stat.Size()
	sz := size / 3

	// mmap the file
	b, err := mmap(file, int(sz), syscall.PROT_READ) // mmap 1/3 of the file
	if err != nil {
		fmt.Println("mmap err:", err)
		return
	}
	fmt.Println("file size:", size, "mmap size:", sz, "mmap bytes:", len(b))

	// read the file content
	br := bytes.NewReader(b)
	buf := bufio.NewReader(br)
	for {
		str, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("line before eof:", str)
			} else {
				fmt.Println("read err:", err)
			}
			break
		}
		fmt.Println(str)
	}

	// munmap file
	err = munmap(b)
	if err != nil {
		fmt.Println("munmap err:", err)
		return
	}

	// unlock file and close file are in defer stack
}

func testMmapWritable() {
	filename := "/home/xzy/t_bz_ga_fz_zfba_xryglaj.csv"
	// open the file
	file, err := os.OpenFile(filename, os.O_RDWR, os.FileMode(0644))
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("file close err:", err)
		}
	}()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	fmt.Println("file path:", file.Name(), "file descriptor:", file.Fd())

	// acquire file lock, exclusive=true if rdwr, exclusive=false if readonly
	if err := flock(file, true, 100*time.Millisecond); err != nil {
		return
	}
	defer func() {
		err := funlock(file)
		if err != nil {
			fmt.Println("funlock err:", err)
		}
	}()

	// get file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("stat err:", err)
		return
	}
	size := stat.Size()
	sz := size / 3

	// mmap the file
	b, err := mmap(file, int(sz), syscall.PROT_READ|syscall.PROT_WRITE) // mmap 1/3 of the file
	if err != nil {
		fmt.Println("mmap err:", err)
		return
	}
	fmt.Println("file size:", size, "mmap size:", sz, "mmap bytes:", len(b))

	pagesize := os.Getpagesize()
	for i := 0; i < len(b)/pagesize; i++ {
		ind := i * pagesize
		byter := b[ind : ind+pagesize]
		for {
			oc := bytes.Index(byter, []byte("合同诈骗案"))
			if oc == -1 {
				break
			}
			copy(b[ind+oc:], []byte("辣真的牛皮"))
		}
		err := fdatasync(file)
		if err != nil {
			fmt.Println("data sync err:", err)
		}
		time.Sleep(10 * time.Millisecond)
	}

	// munmap file
	err = munmap(b)
	if err != nil {
		fmt.Println("munmap err:", err)
		return
	}

	// unlock file and close file are in defer stack
}

func testMmapReadonlyWrite() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	filename := "/home/xzy/t_bz_ga_fz_zfba_xryglaj.csv"
	// open the file
	file, err := os.OpenFile(filename, os.O_RDWR, os.FileMode(0644))
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("file close err:", err)
		}
	}()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	fmt.Println("file path:", file.Name(), "file descriptor:", file.Fd())

	// acquire file lock, exclusive=true if rdwr, exclusive=false if readonly
	if err := flock(file, true, 100*time.Millisecond); err != nil {
		return
	}
	defer func() {
		err := funlock(file)
		if err != nil {
			fmt.Println("funlock err:", err)
		}
	}()

	// get file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("stat err:", err)
		return
	}
	size := stat.Size()
	sz := size / 3

	// mmap the file, can't modify mmap'ed bytes in PROT_READ mode, will throw SIGSEGV
	// fatal error: fault
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x7f7780b2d433 pc=0x4537fe]
	b, err := mmap(file, int(sz), syscall.PROT_READ) // mmap 1/3 of the file
	if err != nil {
		fmt.Println("mmap err:", err)
		return
	}
	fmt.Println("file size:", size, "mmap size:", sz, "mmap bytes:", len(b))

	pagesize := os.Getpagesize()
	for i := 0; i < len(b)/pagesize; i++ {
		ind := i * pagesize
		byter := b[ind : ind+pagesize]
		for {
			oc := bytes.Index(byter, []byte("辣真的牛皮"))
			if oc == -1 {
				break
			}
			copy(b[ind+oc:], []byte("图片及原创"))
		}
		err := fdatasync(file)
		if err != nil {
			fmt.Println("data sync err:", err)
		}
		time.Sleep(10 * time.Millisecond)
	}

	// munmap file
	err = munmap(b)
	if err != nil {
		fmt.Println("munmap err:", err)
		return
	}

	// unlock file and close file are in defer stack
}

func testMmapNewFile() {
	filename := "/tmp/testmmap"
	// open the file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.FileMode(0666))
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("file close err:", err)
		}
	}()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	fmt.Println("file path:", file.Name(), "file descriptor:", file.Fd())

	// acquire file lock, exclusive=true if rdwr, exclusive=false if readonly
	if err := flock(file, true, 100*time.Millisecond); err != nil {
		return
	}
	defer func() {
		err := funlock(file)
		if err != nil {
			fmt.Println("funlock err:", err)
		}
	}()

	// expand file size
	numpage := 3
	pagesize := os.Getpagesize()
	err = file.Truncate(int64(pagesize * numpage))
	if err != nil {
		fmt.Println("truncate err:", err)
	}

	// mmap the file
	b, err := mmap(file, numpage*pagesize, syscall.PROT_WRITE)
	if err != nil {
		fmt.Println("mmap err:", err)
		return
	}
	fmt.Println("mmap blen:", len(b))

	// write some number to mmap bytes
	for i := 0; i < numpage; i++ {
		nrand := rand.Int63n(int64(pagesize)/int64(unsafe.Sizeof(int64(0))) - 1)
		fmt.Println("page", i, nrand)
		isize := int(unsafe.Sizeof(int64(0)))
		pagebase := i * pagesize
		binary.PutVarint(b[pagebase:pagebase+isize], int64(nrand))
		bi := pagebase + isize
		for j := int64(0); j < nrand; j++ {
			n := rand.Int63n(10000000000)
			binary.PutVarint(b[bi:bi+isize], n)
			bi += isize
		}
	}
	// print mmap'ed bytes
	for i := 0; i < len(b); i += 32 {
		fmt.Println(b[i : i+32])
	}

	// munmap file
	err = munmap(b) // file is synced when munmap
	if err != nil {
		fmt.Println("munmap err:", err)
		return
	}
	b = nil

	// unlock file and close file are in defer stack
}

func fdatasync(file *os.File) error {
	return syscall.Fdatasync(int(file.Fd()))
}

func flock(file *os.File, exclusive bool, timeout time.Duration) error {
	var t time.Time
	if timeout != 0 {
		t = time.Now()
	}
	fd := file.Fd()
	flag := syscall.LOCK_NB
	if exclusive {
		flag |= syscall.LOCK_EX
	} else {
		flag |= syscall.LOCK_SH
	}
	for {
		// Attempt to obtain an exclusive lock.
		err := syscall.Flock(int(fd), flag)
		if err == nil {
			return nil
		} else if err != syscall.EWOULDBLOCK {
			return err
		}

		// If we timed out then return an error.
		if timeout != 0 && time.Since(t) > timeout-flockRetryTimeout {
			return errors.New("lock file timeout")
		}

		// Wait for a bit and try again.
		time.Sleep(flockRetryTimeout)
	}
}

// funlock releases an advisory lock on a file descriptor.
func funlock(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}

// mmap memory maps a DB's data file.
func mmap(file *os.File, sz int, prot int) ([]byte, error) {
	// Map the data file to memory.
	b, err := syscall.Mmap(int(file.Fd()), 0, sz, prot, syscall.MAP_SHARED)
	if err != nil {
		return nil, err
	}

	// Advise the kernel that the mmap is accessed randomly.
	err = madvise(b, syscall.MADV_RANDOM)
	if err != nil && err != syscall.ENOSYS {
		// Ignore not implemented error in kernel because it still works.
		return nil, fmt.Errorf("madvise: %s", err)
	}

	// Save the original byte slice and convert to a byte array pointer.
	return b, nil
}

// munmap unmaps a DB's data file from memory.
func munmap(b []byte) error {
	// Ignore the unmap if we have no mapped data.
	if b == nil {
		return nil
	}

	// Unmap using the original byte slice.
	err := syscall.Munmap(b)
	b = nil
	return err
}

// NOTE: This function is copied from stdlib because it is not available on darwin.
func madvise(b []byte, advice int) (err error) {
	// Advise the kernel that the mmap is accessed randomly.
	advice = syscall.MADV_RANDOM
	_, _, e1 := syscall.Syscall(syscall.SYS_MADVISE, uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), uintptr(advice))
	if e1 != 0 {
		err = e1
	}
	return
}
