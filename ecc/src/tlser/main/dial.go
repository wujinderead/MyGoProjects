package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	tlsDial()
}

func tlsDial() {
	clientKeyPair, err := tls.LoadX509KeyPair("keys/ec.client.crt", "keys/ec.client.key")
	if err != nil {
		fmt.Println("get key pair err:", err)
		return
	}
	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("keys/ec.ca.crt")
	certPool.AppendCertsFromPEM(caCert)
	clientConfig := &tls.Config{
		Certificates: []tls.Certificate{clientKeyPair},
		RootCAs:      certPool,
	}
	conn, err := tls.Dial("tcp", "localhost:12345", clientConfig)
	defer toClose2(conn)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	_, err = conn.Write([]byte("my name is van from " + conn.LocalAddr().String()))
	if err != nil {
		fmt.Println("write err:", err)
	}
}

func toClose2(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}
}
