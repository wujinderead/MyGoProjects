package test

import (
	"testing"
)

func TestShaFromFile(t *testing.T) {
	file := "/home/xzy/概要设计说明V1.6.pdf"

	sha1sum, _ := ShaSumFromFile(file, 1)
	t.Logf("sha1sum : %x", sha1sum)
	sha1str, _ := ShaSumFromCmd(file, 1)
	t.Log("sha1Cmd: ", sha1str)

	sha224sum, _ := ShaSumFromFile(file, 224)
	t.Logf("sha224sum : %x", sha224sum)
	sha224str, _ := ShaSumFromCmd(file, 224)
	t.Log("sha224Cmd: ", sha224str)

	sha256sum, _ := ShaSumFromFile(file, 256)
	t.Logf("sha256sum : %x", sha256sum)
	sha256str, _ := ShaSumFromCmd(file, 256)
	t.Log("sha224Cmd: ", sha256str)

	sha384sum, _ := ShaSumFromFile(file, 384)
	t.Logf("sha384sum : %x", sha384sum)
	sha384str, _ := ShaSumFromCmd(file, 384)
	t.Log("sha224Cmd: ", sha384str)

	sha512sum, _ := ShaSumFromFile(file, 512)
	t.Logf("sha512sum : %x", sha512sum)
	sha512str, _ := ShaSumFromCmd(file, 512)
	t.Log("sha224Cmd: ", sha512str)
}

func TestMd5FromFile(t *testing.T) {
	file := "/home/xzy/概要设计说明V1.6.pdf"

	md5sum, _ := Md5SumFromFile(file)
	t.Logf("md5sum : %x", md5sum)
	md5str, _ := Md5SumFromCmd(file)
	t.Log("md5str: ", md5str)
}
