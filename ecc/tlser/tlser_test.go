package tlser

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestX509LeyPair(t *testing.T) {
	// get key pair from pem bytes
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

	// get key pair
	certificate, err = tls.LoadX509KeyPair("../../keys/ec.server.crt", "../../keys/ec.server.key")
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

func TestTlsConfig(t *testing.T) {
	serverKeypair, err := tls.LoadX509KeyPair("../../keys/ec.server.crt", "../../keys/ec.server.key")
	if err != nil {
		t.Error("get key pair err:", err)
		t.FailNow()
	}
	_, err = tls.LoadX509KeyPair("../../keys/ec.client.crt", "../../keys/ec.client.key")
	if err != nil {
		t.Error("get key pair err:", err)
		t.FailNow()
	}
	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("../../ec.ca.crt")
	certPool.AppendCertsFromPEM(caCert)
	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{serverKeypair},
		ClientAuth:   tls.NoClientCert,
		ClientCAs:    certPool,
		ServerName:   "node1",
	}
	clientConfig := &tls.Config{
		RootCAs: certPool,
	}
	listener, err := tls.Listen("tcp", "localhost:8080", serverConfig)
	defer toClose(listener)
	if err != nil {
		t.Error("listen err:", err)
		t.FailNow()
	}
	ch := make(chan struct{})
	go func() {
		time.Sleep(500 * time.Second)
		conn, err := tls.Dial("tcp", "localhost:8080", clientConfig)
		defer toClose(conn)
		if err != nil {
			fmt.Println("dial err:", err)
		}
		_, err = conn.Write([]byte("my name is van from " + conn.LocalAddr().String()))
		if err != nil {
			fmt.Println("write err:", err)
		}
		ch <- struct{}{}
	}()
	conn, err := listener.Accept()
	defer toClose(conn)
	if err != nil {
		fmt.Println("accept err:", err)
	}
	byter := make([]byte, 128)
	n, err := conn.Read(byter)
	if err != nil {
		fmt.Println("read err:", err)
	}
	fmt.Println("read: ", string(byter[:n]), conn.LocalAddr().String())
	<-ch
}

func toClose(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}
}
