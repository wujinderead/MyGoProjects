package stdlib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestIoUtil(t *testing.T) {
	// return file content
	content, err := ioutil.ReadFile("/home/xzy/.bashrc")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	lines := bytes.Split(content, []byte{'\n'})
	fmt.Println("line0:", string(lines[0]))
	fmt.Println("line1:", string(lines[1]))
	fmt.Println()

	// get all file infos in a directory
	fileinfos, err := ioutil.ReadDir("/home/xzy/")
	if err != nil {
		fmt.Println("read dir err:", err)
	}
	for i := range fileinfos {
		fmt.Printf("%15s, %10s, %8d, %v\n",
			fileinfos[i].Name(), fileinfos[i].Mode(), fileinfos[i].Size(), fileinfos[i].IsDir())
	}
	fmt.Println()

	// read add content from a io.Reader until error or EOF
	file, err := os.Open("/home/xzy/.bashrc")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer func() {
		err := file.Close()
		fmt.Println("close file, err:", err)
	}()
	content, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read all err:", err)
		return
	}
	lines = bytes.Split(content, []byte{'\n'})
	fmt.Println("line0:", string(lines[0]))
	fmt.Println("line1:", string(lines[1]))
	fmt.Println()
}

func TestTempFile(t *testing.T) {
	// create a temp dir, if dir is "", it use os.TempDir
	name, err := ioutil.TempDir("", "lalala-")
	if err != nil {
		fmt.Println("temp dir err:", err)
		return
	}
	fmt.Println("temp dir:", name)

	// create a temp file, if dir is "", it use os.TempDir
	file, err := ioutil.TempFile(name, "tester-*")
	if err != nil {
		fmt.Println("temp dir err:", err)
		return
	}
	fmt.Println("temp file:", file.Name())
	err = file.Close()
	if err != nil {
		fmt.Println("close err:", err)
		return
	}

	// write content to file
	err = ioutil.WriteFile(file.Name(), []byte("hahaha\n"), os.FileMode(644))
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
}
