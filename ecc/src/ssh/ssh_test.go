package ssh

import (
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"math/big"
	"strings"
	"testing"
	"util"
)

var (
	rsaPriv = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzKbYNIc067KkDFHYxF2Tg3qUlB6/lV5o086PIRpRNCbt0c3H
UJIFq84+GmkvgmiuLScF9rRiOVlJDLMYnHS/ugzHGlJNwz9aU6aoT24kki/Stjai
tN/CWGJMVvNy4UY9QRt8laETcYV9r0XU7DnfOAdVM2b0gMHp6rgro6zA3mGK3wOR
I4Cq9D6oZT4bL3T6rUzC/A7HkS5vNuE2PALA7i+JLBnOVncoyOL7O0PS7NVtOZAx
A/3N4+ajb0x/m3eR5j/hOHsClT1USlgpM98YkPtWY5KwP2hI6Ynp4LqdlX7laooD
pRlOROPOchz4SpAEool0iSGPIQYt6EdwEl6ygwIDAQABAoIBAB6y/IXMrnSY5KDw
eiriuqkjbzxU7HpUojb7ql7V2s6O9GffjYGZlf4yvwApPTY7y7z8OJnMb1uY+CtO
hmeZ39Th69AX/pBGZZ9cxay8ogHH6LzqrzegxT+K51a3yEjgx3mHzQWJFyiVVMhB
GnKhL8nw16gRTqYt8JAENo+j736sMEkbfoNRrxm8TpBl1RaDkw9FmNMWvPKpkFf4
HYtWKvzlX87hxAVZ2X0a0VrnyN63zv9f1Hvi6g0wAWrG+7bT6zPZEil2oAKC36II
yuvHHD4Wc4405CwO73PtxOTAJvK4MSqhxOnwLcHxTCfw7+S/Orsd22VzFmOpfziE
u1YfIBECgYEA+GVleyLOTKnNNfah7bFAMBcAm2/Glh1iwFElHBYoMW9finuFbS1w
YzHYGVivW9wlxS72rPxUpPrH6yVWmtdEPi+RYdVpcxEHj0q2kyE6xNqdlyk/Oa/x
JW/4FfaGG1zd0WjqlRiShDE2TJf38hPY1yE/LuDRq8mHfYegDeu0TfsCgYEA0uqj
Hx8h6cAiOawh6CpQWlUTdnUkzcKvCfjCE/bv4P/HkMtBeG/ihzmnobQYQoDStrkm
msUyNdmSlR1KltVYLETac/evUzrpBpHjIVipsNOXVLTv5uy2TGnV7V5XuQNLpMSI
0Jjx+m4dD3Osq9zmbZSmDM3Svl8TVCtTMxm2LxkCgYEA+AJtd6vB5YOotFejSCsx
FpLw9UF+O0Xt4m1iqw9oZCt6bk90YhT7YN9Uj8IfnI1LXPzOKNvsO6l1UNBAD2wd
5CUkeFVX6x62uJh1gKOuBPzuWg5B5XxJPwLz5iH1tn5br4mcpu8Y40ormAAn/RlZ
6Tp11n18e5RFZs2yvhN4PF8CgYB2LOGY3mix/+UtSzT0UEEVW/W7uYcVgq9wduDH
LuTYvHekuT3FrWrPOY6jG7U8DdICb1sh/LtVUMLAqdjRCliM9UcxEuY5TBikhbkt
RfBOE0AHRhnk2VyLFAG5LdMY5q/LchL2TbvHBUtjDP0CjpLNcyxWoDwkTkEWN/A2
AYICgQKBgGzgI9Jier+/xMNn+VF7Ud/ldQMXIGtyNm2dhP7dowIfFaU3OpyRmnFy
6UWGFNzPJid98sK2DeUzbQKMS5z/ir4siXWr+E9XmjYR9YJzcJAlmfF2r2mx6JY3
8Gzr2HFxkGtkaRhd9Dj3GBslIV+AzeuXptj4zL6Y2ZrhaYjU11x9
-----END RSA PRIVATE KEY-----`
	ecdsaPriv = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACDkYFJHfIBVV2iSyBK8gzmmZb2t6u4dGZ/WK+l5uOodMQAAAJDquhTa6roU
2gAAAAtzc2gtZWQyNTUxOQAAACDkYFJHfIBVV2iSyBK8gzmmZb2t6u4dGZ/WK+l5uOodMQ
AAAEAAN96983j4YQEmgeD1RG+Nh4IfPkw18ysqdt3oRmKx4ORgUkd8gFVXaJLIEryDOaZl
va3q7h0Zn9Yr6Xm46h0xAAAACXh6eUBub2RlMAECAwQ=
-----END OPENSSH PRIVATE KEY-----`
	ed25519Priv = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIcFXlQ3DW8SbZgZLCTJuSY/ItXmOQnWROFQc6pvv+msoAoGCCqGSM49
AwEHoUQDQgAEenAv4+PY+889WXF7EjTZ3kPyRbBPMox7bOqdBjYrcJ8citj54i1+
HSrJqy5hUNQBCq7XjhMmo/AU15pFuShk1Q==
-----END EC PRIVATE KEY-----`
	rsaPub     = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDMptg0hzTrsqQMUdjEXZODepSUHr+VXmjTzo8hGlE0Ju3RzcdQkgWrzj4aaS+CaK4tJwX2tGI5WUkMsxicdL+6DMcaUk3DP1pTpqhPbiSSL9K2NqK038JYYkxW83LhRj1BG3yVoRNxhX2vRdTsOd84B1UzZvSAwenquCujrMDeYYrfA5EjgKr0PqhlPhsvdPqtTML8DseRLm824TY8AsDuL4ksGc5WdyjI4vs7Q9Ls1W05kDED/c3j5qNvTH+bd5HmP+E4ewKVPVRKWCkz3xiQ+1ZjkrA/aEjpiengup2VfuVqigOlGU5E485yHPhKkASiiXSJIY8hBi3oR3ASXrKD xzy@node0"
	ecdsaPub   = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHpwL+Pj2PvPPVlxexI02d5D8kWwTzKMe2zqnQY2K3CfHIrY+eItfh0qyasuYVDUAQqu144TJqPwFNeaRbkoZNU= xzy@node0"
	ed25519Pub = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIORgUkd8gFVXaJLIEryDOaZlva3q7h0Zn9Yr6Xm46h0x xzy@node0"
)

func TestParseRawPrivateKey(t *testing.T) {
	for _, privFile := range []string{rsaPriv, ecdsaPriv, ed25519Priv} {
		block, _ := pem.Decode([]byte(privFile))
		fmt.Println("type: ", block.Type)
		privKey, err := ssh.ParseRawPrivateKey([]byte(privFile))
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
			// change from little-endian to big-endian
			reverse := func(arr []byte) []byte {
				for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
					arr[i], arr[j] = arr[j], arr[i]
				}
				return arr
			}
			edpub := util.Curve25519XToEd25519Y(new(big.Int).SetBytes(reverse(mtpub[:])))
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
	for _, pubkey := range []string{rsaPub, ecdsaPub, ed25519Pub} {
		strs := strings.Split(pubkey, " ")
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
		fmt.Println("finger printer:", fingerprint)
		hashed := sha256.Sum256(pubKey.Marshal())
		base64ed := base64.RawStdEncoding.EncodeToString(hashed[:])
		fmt.Println("my generated  :       ", base64ed)
		fmt.Println()
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
	// hashed hostname (`xzy@10.19.138.135`) example:
	// |1|JlTvbaou69qWu4ApI7cJRTIb+Ro=|OWq2xirrJRIVcI11tkg0a0ZC2BY=
	// format: base64-ed random salt, based64-ed hmac-sha1 of hostname using salt as ramdom key
	h135 := knownhosts.HashHostname("xzy@10.19.138.135")
	h181 := knownhosts.HashHostname("xh@10.19.138.181")
	h186 := knownhosts.HashHostname("taodd@10.19.138.186")
	fmt.Println("h135: ", h135)
	fmt.Println("h181: ", h181)
	fmt.Println("h186: ", h186)

	hostname := "xzy@10.19.138.135"
	hashedHost := knownhosts.HashHostname(hostname)
	base64ed_salt := strings.Split(hashedHost, "|")[2]
	base64ed_hmac := strings.Split(hashedHost, "|")[3]
	fmt.Printf("host: %s, hashedHost: %s\n", hostname, hashedHost)
	salt, _ := base64.StdEncoding.DecodeString(base64ed_salt)
	fmt.Println("salt: ", hex.EncodeToString(salt))
	hmacer := hmac.New(sha1.New, salt)
	hmacer.Write([]byte("xzy@10.19.138.135"))
	sum := hmacer.Sum(nil)
	fmt.Println("hashed name: ", hex.EncodeToString(sum))
	decoded, _ := base64.StdEncoding.DecodeString(base64ed_hmac)
	fmt.Println("decode hash: ", hex.EncodeToString(decoded))
}
