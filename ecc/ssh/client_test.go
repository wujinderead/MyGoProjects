package ssh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"testing"
)

var closer = func(str string, closer io.Closer) {
	err := closer.Close()
	if err != nil {
		fmt.Printf("close '%s', err: %s\n", str, err.Error())
	} else {
		fmt.Println(str, "closed")
	}
}

func TestSshConnect(t *testing.T) {
	config := new(ssh.ClientConfig)
	config.User = "xzy"
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	config.Auth = []ssh.AuthMethod{ssh.Password("111")}
	client, err := ssh.Dial("tcp", "10.19.138.135:22", config)
	if err != nil {
		fmt.Println("dial err:", err.Error())
		t.FailNow()
	}
	fmt.Println("client user:", client.User())
	defer closer("client", client)

	session, err := client.NewSession()
	if err != nil {
		fmt.Println("create session err:", err.Error())
		t.FailNow()
	}
	defer closer("session", session)

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ifconfig"); err != nil {
		fmt.Println("session run err:", err.Error())
	}
	fmt.Println("ifconfig:\n", b.String())
}
