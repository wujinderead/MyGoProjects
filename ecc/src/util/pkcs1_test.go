package util

import (
	"testing"
	"encoding/pem"
	"io/ioutil"
	"crypto/x509"
	"fmt"
	"encoding/hex"
	"golang.org/x/crypto/ssh"
	"encoding/base64"
)

const rsaKeyStr = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCjWKgqV5fSp8ThXVwvQc29988gXUNZ06fxCHXPjaO/TI7Dj8EhsaD3pt691d1oorDfpiKe1Cs+GqiI6bUb90eBljUgoEkYL5SroOsQOkBTqXTij0np4/piOt2ofQJNMDqCbY+D8GX3yGZRWtc0rfrU+t1TzmrWqsXyrbKHcAp4x6mGTjoFwHMb+bzoRcZyg6PwnV19MMJQj0BMEs7KJyzd7kiz3oE5yLIIkp1eCmBi2tRkuc83+rCNluqtfZhPTvnf8IpBY7GsCQwytyl2cdmGNO2bXxZGktYmLIAU3eOIHdEwWyBZGhAfh/70mbcIkQ+2M5VjJMPIQbodWKhExH4X"
const ecKeyStr = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCPByzkRCJe/NzvR5nbJ8OlZ4VFKpJgDLeIPuekiUfgpxLWNu/yZzqX5+IFIJ8TVYw0ILH5ZFxNkqcTrSrNQB0k="
const edKeyStr = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIH+J7IhCQyuIRRFHfTPZ3LdeGG/0lZoMAgLcySptrgvU"

func TestRsaKey(t *testing.T) {
	rsaKeyContent, _ := ioutil.ReadFile("/home/xzy/ssh_rsa_priv.key")
	block, _ := pem.Decode(rsaKeyContent)
	rsaPrivKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	fmt.Println("d: ", hex.EncodeToString(rsaPrivKey.D.Bytes()))
	fmt.Println("e: ", rsaPrivKey.E)
	fmt.Println("n: ", hex.EncodeToString(rsaPrivKey.N.Bytes()))

	ecKeyContent, _ := ioutil.ReadFile("/home/xzy/ssh_ecdsa_priv.key")
	block, _ = pem.Decode(ecKeyContent)
	ecPrivKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	fmt.Println("name:", ecPrivKey.Curve.Params().Name)
	fmt.Println("d:", hex.EncodeToString(ecPrivKey.D.Bytes()))

	func () {
		var err error
		ecKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(ecKeyStr))
		if err != nil {
			panic(err)
		}
		fmt.Println(ecKey.Type())
		fmt.Println(base64.StdEncoding.EncodeToString(ecKey.Marshal()))
		edKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(edKeyStr))
		if err != nil {
			panic(err)
		}
		fmt.Println(edKey.Type())
		rsaKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(rsaKeyStr))
		if err != nil {
			panic(err)
		}
		fmt.Println(rsaKey.Type())
	}()
}
