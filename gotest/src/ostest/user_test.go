package ostest

import (
	"testing"
	"os/user"
	"fmt"
)

func TestUser(t *testing.T) {
	myself, err := user.Current()
	if err != nil {
		fmt.Println("get self err: ", err.Error())
	}
	fmt.Println("name: ", myself.Name)
	fmt.Println("gid: ", myself.Gid)
	fmt.Println("home: ", myself.HomeDir)
	fmt.Println("uid: ", myself.Uid)
	fmt.Println("uname: ", myself.Uid)
	gids, err := myself.GroupIds()
	if err != nil {
		fmt.Println("get gids err: ", err.Error())
	}
	fmt.Println("gids: ", gids)

	group, err := user.LookupGroup("sshd")
	if err != nil {
		fmt.Println("get group err:", err)
	}
	fmt.Println("gid:", group.Gid)
	fmt.Println("name:", group.Name)
}
