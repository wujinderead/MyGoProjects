package util

import (
	"sync"
	"os"
	"fmt"
)

var filename = "/dev/urandom"
var lock sync.Mutex
var reservoir = make([]byte, 32)
var pos = int64(len(reservoir))
var file *os.File = nil

// read from reservoir and put to byter
func NextBytes(bytes []byte) {
	lock.Lock()
	defer lock.Unlock()
	//lenBytes := int64(len(bytes))
	off := int64(0)
	for off < int64(len(bytes)) {
		fillInReservoir()
		toFill := min(int64(len(bytes)) - off, int64(len(reservoir)) - pos)
		fmt.Printf("tofill: %d, len: %d, off: %d, lenr: %d, pos :%d, minus1 :%d, minus2: %d\n",
			toFill, int64(len(bytes)), off, int64(len(reservoir)), pos, int64(len(bytes)) - off, int64(len(reservoir)) - pos)
		copy(bytes[off:], reservoir[pos:pos+toFill])
		off += toFill
		pos += toFill
		fmt.Printf("bytes: %x\n", bytes)
		fmt.Printf("tofill: %d, len: %d, off: %d, lenr: %d, pos :%d, minus1 :%d, minus2: %d\n",
			toFill, int64(len(bytes)), off, int64(len(reservoir)), pos, int64(len(bytes)) - off, int64(len(reservoir)) - pos)
	}
}

// set path for random device
func SetRandomDevPath(path string) {
	lock.Lock()
	defer lock.Unlock()
	filename = path
	Close()
}

// reset the reservoir and close the file
func Close() {
	lock.Lock()
	defer lock.Unlock()
	if file != nil {
		file.Close()
	}
	file = nil
}

func fillInReservoir() {
	if pos == int64(len(reservoir)) {
		if file == nil {
			file, err := os.Open(filename)
			if err != nil {
				panic("error open file")
			}
			_, err = file.Read(reservoir)
			if err != nil {
				panic("error read file")
			}
			fmt.Printf("reservoir: %x\n", reservoir)
		}
		pos = 0
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}