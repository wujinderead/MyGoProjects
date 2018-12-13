package ostest

import (
	"fmt"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	// get command that start current process
	execer, err := os.Executable()
	if err != nil {
		fmt.Println("get executable err: ", err.Error())
		return
	}
	fmt.Println("who exec current process: ", execer)

	// get current process info
	fmt.Println("egid:", os.Getegid())
	fmt.Println("euid:", os.Geteuid())
	fmt.Println("gid:", os.Getgid())
	fmt.Println("uid:", os.Getuid())
	fmt.Println("pid:", os.Getpid())
	fmt.Println("ppid", os.Getppid())
	fmt.Println("pagesize:", os.Getpagesize())
	host, _ := os.Hostname()
	fmt.Println("hostname:", host)
	wd, _ := os.Getwd()
	fmt.Println("wd:", wd)
	groups, _ := os.Getgroups()
	fmt.Println("groups:", groups)
	fmt.Println("env:", os.Environ())
	fmt.Println("expanded:", os.Expand("===${aa}===", func(s string) string {return "haha"}))

	// get and make link
	linker, err := os.Readlink("/usr/local/man")
	fmt.Println("link:", linker)
	if err != nil {
		fmt.Println("get link err: ", err.Error())
		return
	}
	err = os.Symlink("/usr/local/go/bin/go", "/home/xzy/goer")
	if err != nil {
		fmt.Println("make symbol link err: ", err.Error())
		return
	}

	// change directory
	err = os.Chdir("/home/xzy")
	if err != nil {
		fmt.Println("change dir err: ", err.Error())
		return
	}

	// create and close file
	file, err := os.Create("testfile")
	defer func() {
		if file != nil {
			if err != nil {
				fmt.Println("close file error: ", err.Error())
			}
		}
	}()
	if err != nil {
		fmt.Println("create err: ", err.Error())
		return
	}

	// get file info
	fd := file.Fd()
	name := file.Name()
	fmt.Println("file name: ", name)
	fmt.Println("file descriptor: ", fd)

	// change file mode
	err = os.Chmod(name, os.FileMode(0775))
	if err != nil {
		fmt.Println("chmod err: ", err.Error())
		return
	}

	// write command to file
	_, _ = file.WriteString("#!/bin/bash\n")
	_, _ = file.WriteString("echo $@;date;sleep 2;date;echo $name")
	err = file.Close()
	if err != nil {
		fmt.Println("close file error: ", err.Error())
		return
	}

	// run it
	var attr os.ProcAttr
	attr.Env = []string{"name=lgq"}
	attr.Dir = "/home/xzy"
	process, err := os.StartProcess("/home/xzy/testfile", []string{"aaa bbb"}, &attr)
	if err != nil {
		fmt.Println("start process err: ", err.Error())
		return
	}
	_,_ = process.Wait()

	// delete file
	err = os.Remove(name)
	if err != nil {
		fmt.Println("remove err: ", err.Error())
	}
}
