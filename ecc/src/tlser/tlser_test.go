package tlser

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestX509LeyPair(t *testing.T) {
	serverKeyPem, _ := ioutil.ReadFile("../../keys/ec.server.key")
	serverCertPem, _ := ioutil.ReadFile("../../keys/ec.server.crt")
	certificate, err := tls.X509KeyPair(serverCertPem, serverKeyPem)
	if err != nil {
		fmt.Println("get key pair err:", err)
		return
	}
	fmt.Println("cert key type:", reflect.TypeOf(certificate.PrivateKey))
	for i := range certificate.Certificate {
		fmt.Println("certificate byte[]:", hex.EncodeToString(certificate.Certificate[i]))
	}
	if certificate.Leaf != nil {
		fmt.Println("leaf issuer:", certificate.Leaf.Issuer)
		fmt.Println("leaf subject:", certificate.Leaf.Subject)
	}
}
