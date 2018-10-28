package ssh

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"math/big"
	"strings"
	"testing"
	)

func TestParsePrivateKey(t *testing.T) {
	byter, err := ioutil.ReadFile("/home/xzy/.ssh/id_rsa")
	if err != nil {
		fmt.Println("read file err: ", err.Error())
		t.FailNow()
	}
	block, _ := pem.Decode(byter)
	fmt.Println("header: ", block.Headers)
	fmt.Println("type: ", block.Type)
	_, err = ssh.ParsePrivateKey(byter)
	if err != nil {
		fmt.Println("parse private key err: ", err.Error())
		t.FailNow()
	}
}

func TestParseRawPrivateKey(t *testing.T) {
	for _, file := range []string{
		"/home/xzy/.ssh/id_rsa",
		"/home/xzy/.ssh/id_ed25519",
		"/home/xzy/.ssh/id_ecdsa"} {
		byter, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("read file err: ", err.Error())
			t.FailNow()
		}
		block, _ := pem.Decode(byter)
		fmt.Println("type: ", block.Type)
		privKey, err := ssh.ParseRawPrivateKey(byter)
		if err != nil {
			fmt.Println("parse private key err: ", err.Error())
			t.FailNow()
		}
		if rsaPrivKey, ok := privKey.(*rsa.PrivateKey); ok {
			fmt.Println("D: ", hex.EncodeToString(rsaPrivKey.D.Bytes()))
			fmt.Println("N: ", hex.EncodeToString(rsaPrivKey.N.Bytes()))
			fmt.Println("p: ", hex.EncodeToString(rsaPrivKey.Primes[0].Bytes()))
			fmt.Println("q: ", hex.EncodeToString(rsaPrivKey.Primes[1].Bytes()))
			fmt.Println("e: ", rsaPrivKey.E)
			n := new(big.Int).Mul(rsaPrivKey.Primes[0], rsaPrivKey.Primes[1])
			fmt.Println(n.Cmp(rsaPrivKey.N))
			one := new(big.Int).SetInt64(1)
			p_1 := new(big.Int).Sub(rsaPrivKey.Primes[0], one)
			q_1 := new(big.Int).Sub(rsaPrivKey.Primes[1], one)
			lambda := new(big.Int).Mul(p_1, q_1)
			lambda.ModInverse(new(big.Int).SetInt64(65537), lambda)
			fmt.Println(lambda.Cmp(rsaPrivKey.D))
		}
		if sshKey, ok := privKey.(*ed25519.PrivateKey); ok {
			fmt.Println("priv: ", hex.EncodeToString([]byte(*sshKey)))
			fmt.Println("pub : ", hex.EncodeToString([]byte(sshKey.Public().(ed25519.PublicKey))))
			sshKey.Public()
		}
		if ecKey, ok := privKey.(*ecdsa.PrivateKey); ok {
			fmt.Println("curve: ", ecKey.Curve.Params().Name)
			fmt.Println("D: ", hex.EncodeToString(ecKey.D.Bytes()))
			fmt.Println("X: ", hex.EncodeToString(ecKey.X.Bytes()))
			fmt.Println("Y: ", hex.EncodeToString(ecKey.Y.Bytes()))
			x, y := ecKey.Curve.ScalarBaseMult(ecKey.D.Bytes())
			fmt.Println(ecKey.X.Cmp(x), ecKey.Y.Cmp(y))
		}
		fmt.Println()
	}
}

func Test(t *testing.T) {
	for _, file := range []string{
		"/home/xzy/.ssh/id_ed25519.pub",
		"/home/xzy/.ssh/id_ecdsa.pub",
		"/home/xzy/.ssh/id_rsa.pub"} {
		byter, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("read file err: ", err.Error())
			t.FailNow()
		}
		str := string(byter)
		strs := strings.Split(str, " ")
		fmt.Println("type: ", strs[0])
		fmt.Println("host: ", strs[2][:len(strs[2])-1])
		fmt.Println("base64 pub key: ", strs[1])
		decoded, err := base64.StdEncoding.DecodeString(strs[1])
		if err != nil {
			fmt.Println("base64 decode err: ", err.Error())
			t.FailNow()
		}
		fmt.Println("base64 decoded: ", hex.EncodeToString(decoded))
		pubKey, err := ssh.ParsePublicKey(decoded)
		if err != nil {
			fmt.Println("parse public key err: ", err.Error())
			t.FailNow()
		}
		fmt.Println("type: ", pubKey.Type())
		fmt.Println("pubkey marshal: ", hex.EncodeToString(pubKey.Marshal()))
		fmt.Println()
	}
}

func Test1(t *testing.T) {
	str := "AAAAC3NzaC1lZDI1NTE5AAAAIORgUkd8gFVXaJLIEryDOaZlva3q7h0Zn9Yr6Xm46h0x"
	byter, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("base64 decode err: ", err.Error())
	}
	fmt.Println(hex.EncodeToString(byter))
	fmt.Println(hex.EncodeToString([]byte("ssh-ed25519")))
	_, err = ssh.ParsePublicKey(byter)
	if err != nil {
		fmt.Println("parse public key err: ", err.Error())
	}

	str = "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHpwL+Pj2PvPPVlxexI02d5D8kWwTzKMe2zqnQY2K3CfHIrY+eItfh0qyasuYVDUAQqu144TJqPwFNeaRbkoZNU="
	byter, err = base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("base64 decode err: ", err.Error())
	}
	fmt.Println(hex.EncodeToString(byter))
	fmt.Println(hex.EncodeToString([]byte("ecdsa-sha2-nistp256")))
	byterr, err := hex.DecodeString("6e69737470323536")
	fmt.Println(string(byterr))
}
