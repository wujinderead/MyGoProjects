package ostest

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
	"syscall"
	"testing"
	"time"
)

func TestLookPath(t *testing.T) {
	for _, command := range []string{"ls", "python", "python3", "java"} {
		path, err := exec.LookPath(command)
		if err != nil {
			fmt.Printf("cmd: %s, err: %s\n", command, err.Error())
		}
		fmt.Printf("%s: %s\n", command, path)
	}
}

func TestRun(t *testing.T) {
	cmd := exec.Command("bash", "-c", `echo "不仅当前进程会收到信号";sleep 2; echo "Μια έκδοσηστηνισπανική"`)
	// this will block until finish
	outpipe, err := cmd.StdoutPipe()
	runerr := cmd.Start()
	if runerr != nil || err != nil {
		fmt.Println("start err: ", runerr)
		fmt.Println("outpipe err: ", err)
		return
	}
	buf := make([]byte, 20)
	for {
		n, err := outpipe.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read end.")
				break
			} else {
				fmt.Println("read err: ", err.Error())
				return
			}
		}
		fmt.Println(string(buf[:n]))
	}
}

func TestOutput(t *testing.T) {
	cmd := exec.Command("bash", "-c", "sleep 3;ls -lh; echo ${name}")
	// it will block until finish
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("get output err: ", err.Error())
		return
	}
	fmt.Println(string(output))
}

func TestStartWait(t *testing.T) {
	cmd := exec.Command("bash", "-c", "echo ${name};sleep 3;ls -lh")
	cmd.Dir = "/home/xzy"
	cmd.Env = []string{"name=lgq"}
	if false {
		cmd.Stdout = os.Stdout // direct the output to os.Stdout (i.e. console)
		cmd.Stderr = os.Stderr // direct the error to os.Stderr (i.e. console)
	}
	if false {
		file, err := os.OpenFile("/home/xzy/outerr", os.O_CREATE | os.O_WRONLY | os.O_APPEND, os.FileMode(0644))
		if err != nil {
			fmt.Println("open file err: ", err.Error())
			t.FailNow()
		}
		cmd.Stdout = file // direct the output to file
		cmd.Stderr = file
	}
	var outbuf *bytes.Buffer = nil
	var errbuf *bytes.Buffer = nil
	if true {
		outbuf = bytes.NewBuffer(nil)
		errbuf = bytes.NewBuffer(nil)
		cmd.Stdout = outbuf // direct the output to a buf
		cmd.Stderr = errbuf
	}

	// won't block
	err := cmd.Start()
	fmt.Println("after start: ", time.Now())
	process := cmd.Process
	fmt.Println("pid: ", process.Pid)
	if false {  // kill the process
		time.Sleep(time.Second)
		err := process.Kill()
		if err != nil {
			fmt.Println("kill err: ", err.Error())
		}
		if outbuf != nil {
			fmt.Println("out: ", outbuf.String())
		}
		if errbuf != nil {
			fmt.Println("err: ", errbuf.String())
		}
		state := cmd.ProcessState  //
		fmt.Println("state: ", state)
		return
	}
	if true {  // signal the process
		time.Sleep(time.Second)
		err := process.Signal(syscall.SIGTERM)
		if err != nil {
			fmt.Println("signal err: ", err.Error())
		}
		if outbuf != nil {
			fmt.Println("out: ", outbuf.String())
		}
		if errbuf != nil {
			fmt.Println("err: ", errbuf.String())
		}
		state := cmd.ProcessState  // state is available after calling Run or Wait, i.e. after the process is finished
		if state != nil {
			fmt.Println("\nstate pid: ", state.Pid())
			fmt.Println("state exited: ", state.Exited())
			fmt.Println("state success: ", state.Success())
			fmt.Println("state string: ", state.String())
			fmt.Println("state sys time: ", state.SystemTime())
			fmt.Println("state user time: ", state.UserTime())
			fmt.Println("state sys type: ", reflect.TypeOf(state.Sys()))
			fmt.Println("state sys usage type: ", reflect.TypeOf(state.SysUsage()))
		}
	}
	fmt.Println()

	// block until finish
	werr := cmd.Wait()
	fmt.Println("wait end: ", time.Now())
	if outbuf != nil {
		fmt.Println("out: ", outbuf.String())
	}
	if errbuf != nil {
		fmt.Println("err: ", errbuf.String())
	}

	state := cmd.ProcessState  // state is available after calling Run or Wait, i.e. after the process is finished
	fmt.Println("\nstate pid: ", state.Pid())
	fmt.Println("state exited: ", state.Exited())
	fmt.Println("state success: ", state.Success())
	fmt.Println("state string: ", state.String())
	fmt.Println("state sys time: ", state.SystemTime())
	fmt.Println("state user time: ", state.UserTime())
	fmt.Println("state sys type: ", reflect.TypeOf(state.Sys()))
	fmt.Println("state sys usage type: ", reflect.TypeOf(state.SysUsage()))
	if err != nil {
		fmt.Println("run err: ", err.Error())
		return
	}
	if werr != nil {
		fmt.Println("wait err: ", werr.Error())
		return
	}
}
