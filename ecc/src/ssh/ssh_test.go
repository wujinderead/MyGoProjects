package ssh

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"io/ioutil"
	"math/big"
	"strings"
	"testing"
	"util"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
)

func TestParsePrivateKey(t *testing.T) {
	byter, err := ioutil.ReadFile("/home/lgq/.ssh/id_ed25519")
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
		"/home/lgq/.ssh/id_rsa",
		"/home/lgq/.ssh/id_ed25519",
		"/home/lgq/.ssh/id_ecdsa"} {
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
		switch privKey.(type) {
		case *rsa.PrivateKey:
			rsaPrivKey := privKey.(*rsa.PrivateKey)
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
		case *ed25519.PrivateKey:
			sshKey := privKey.(*ed25519.PrivateKey)
			fmt.Println("priv: ", hex.EncodeToString([]byte(*sshKey)))
			fmt.Println("pub : ", hex.EncodeToString([]byte(sshKey.Public().(ed25519.PublicKey))))
			// the private key of ed25519 is actually a seed
			seed := []byte(*sshKey)[:ed25519.SeedSize]
			dgst := sha512.Sum512(seed)
			var mtpub, multiplier [32]byte
			copy(multiplier[:], dgst[:])
			curve25519.ScalarBaseMult(&mtpub, &multiplier)
			reverse := func(arr []byte) []byte {
				for i, j:= 0, len(arr)-1; i<j; i, j = i+1, j-1 {
					arr[i], arr[j] = arr[j], arr[i]
				}
				return arr
			}
			edpub := util.Curve25519XToEd25519Y(new(big.Int).SetBytes(reverse(mtpub[:])));
			fmt.Printf("edy :  %x\n", reverse(edpub.Bytes()))
		case *ecdsa.PrivateKey:
			ecKey := privKey.(*ecdsa.PrivateKey)
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

func TestParsePub(t *testing.T) {
	for _, file := range []string{
		"/home/lgq/.ssh/id_ed25519.pub",
		"/home/lgq/.ssh/id_ecdsa.pub",
		"/home/lgq/.ssh/id_rsa.pub"} {
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

		// fingerprint is to sha256 the public key marshal, then base64 (RawStdEncoding) the hash
		fingerprint := ssh.FingerprintSHA256(pubKey)
		fmt.Println("finger printer: ", fingerprint)
	}
}

func TestParsePubBase64(t *testing.T) {
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

func TestParsePubBase641(t *testing.T) {
	str := "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBN19Wr/teJ1PArr+VeLmhusv+gEaE9jcEDYlptIR/+XE5joPoKlmlrLO66iazPBk5RTzYJpZPWNprldmq/RueGw="
	byter, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("base64 decode err: ", err.Error())
	}
	pub, err := ssh.ParsePublicKey(byter)
	if err != nil {
		fmt.Println("parse public key err: ", err.Error())
	}
	fmt.Println(pub.Type())
	fmt.Println(hex.EncodeToString(pub.Marshal()))
	fp := ssh.FingerprintSHA256(pub)
	fmt.Println(fp)
}

func TestParseAuthorizedKey(t *testing.T) {
	str := "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIORgUkd8gFVXaJLIEryDOaZlva3q7h0Zn9Yr6Xm46h0x xzy@node0"
	pub, comment, opts, rest, err := ssh.ParseAuthorizedKey([]byte(str))
	if err != nil {
		fmt.Println("base64 decode err: ", err.Error())
		t.FailNow()
	}
	fmt.Println("pub: ", pub.Type())
	fmt.Println("comment: ", comment)
	fmt.Println("opts: ", opts)
	fmt.Println("rest: ", rest)
}

func TestSignAndVerify25519(t *testing.T) {
	msg := "A quick brown fox jumps over the lazy dog."
	pub1, priv1, _ := ed25519.GenerateKey(rand.Reader)
	pub2, priv2, _ := ed25519.GenerateKey(rand.Reader)
	sshPub1, _ := ssh.NewPublicKey(pub1)
	sshPub2, _ := ssh.NewPublicKey(pub2)
	signer1, _ := ssh.NewSignerFromKey(priv1)
	signer2, _ := ssh.NewSignerFromKey(priv2)
	signature1, _ := signer1.Sign(rand.Reader, []byte(msg))
	signature2, _ := signer2.Sign(rand.Reader, []byte(msg))
	fmt.Printf("signture1 format: %s, blob: %s\n", signature1.Format, hex.EncodeToString(signature1.Blob))
	fmt.Printf("signture2 format: %s, blob: %s\n", signature2.Format, hex.EncodeToString(signature2.Blob))
	fmt.Println(sshPub1.Verify([]byte(msg), signature1))
	fmt.Println(sshPub1.Verify([]byte(msg), signature2))
	fmt.Println(sshPub2.Verify([]byte(msg), signature1))
	fmt.Println(sshPub2.Verify([]byte(msg), signature2))
}

func TestHashHostname(t *testing.T) {
	// hashed hostname example:
	// |1|JlTvbaou69qWu4ApI7cJRTIb+Ro=|OWq2xirrJRIVcI11tkg0a0ZC2BY=
	// format: base64-ed salt, based64-ed hmac-sha1 of hostname using salt as ramdom key
	h135 := knownhosts.HashHostname("xzy@10.19.138.135")
	h181 := knownhosts.HashHostname("xh@10.19.138.181")
	h186 := knownhosts.HashHostname("taodd@10.19.138.186")
	fmt.Println("h135: ", h135)
	fmt.Println("h181: ", h181)
	fmt.Println("h186: ", h186)
	salt, _ := base64.StdEncoding.DecodeString("JlTvbaou69qWu4ApI7cJRTIb+Ro=")
	fmt.Println("salt: ", hex.EncodeToString(salt))
	hmacer := hmac.New(sha1.New, salt)
	hmacer.Write([]byte("xzy@10.19.138.135"))
	sum := hmacer.Sum(nil)
	fmt.Println("hashed name: ", hex.EncodeToString(sum))
	decoded, _ := base64.StdEncoding.DecodeString("OWq2xirrJRIVcI11tkg0a0ZC2BY=")
	fmt.Println("decode hash: ", hex.EncodeToString(decoded))
}