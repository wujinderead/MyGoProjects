package util

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_pem(t *testing.T) {
	filename := "/home/xzy/ca.key"
	byter, _ := ioutil.ReadFile(filename)
	block, rest := pem.Decode(byter)
	fmt.Printf("rest: %x\n", rest)
	fmt.Println("header: ", block.Headers)
	fmt.Println("type: ", block.Type)
	fmt.Printf("bypes: %x\n", block.Bytes)

	file, _ := os.Open(filename)
	br := bufio.NewReader(file)
	buffer := bytes.NewBufferString("")
	for {
		str, err := br.ReadString('\n')
		if err != nil {
			break
		}
		if !strings.HasPrefix(str, "-----") {
			buffer.WriteString(str)
		}
	}
	b64d, _ := base64.StdEncoding.DecodeString(buffer.String())
	fmt.Printf("b64d : %x\n", b64d)
	fmt.Println(bytes.Equal(b64d, block.Bytes))
}

func Test_pem_rsa(t *testing.T) {
	filename := "/home/xzy/ca.key"
	byter, _ := ioutil.ReadFile(filename)
	block, rest := pem.Decode(byter)
	fmt.Printf("rest: %x\n", rest)
	fmt.Println("header: ", block.Headers)
	fmt.Println("type: ", block.Type)
	fmt.Printf("bytes: %x\n", block.Bytes)
}

func Test_pem_ec(t *testing.T) {
	filename := "/home/lgq/ec.key"
	byter, _ := ioutil.ReadFile(filename)
	block, rest := pem.Decode(byter)
	fmt.Printf("rest: %x\n", rest)
	fmt.Println("header: ", block.Headers)
	fmt.Println("type: ", block.Type)
	fmt.Printf("bytes: %x\n", block.Bytes)

	block, rest = pem.Decode(rest)
	fmt.Printf("rest: %x\n", rest)
	fmt.Println("header: ", block.Headers)
	fmt.Println("type: ", block.Type)
	fmt.Printf("bytes: %x\n", block.Bytes)
}

/*
openssl genrsa -out rsa_private_key.pem 512
openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

openssl ecparam -list_curves
openssl ecparam -genkey -out ec.key -name secp384r1
openssl ec -in ec.key -text
openssl ec -in ec.key -pubout -out ec.pub
openssl ec -in ec.pub -pubin -text
*/
