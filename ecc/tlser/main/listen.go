package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	tlsListen()
}

func tlsListen() {
	//serverKeypair, err := tls.LoadX509KeyPair("keys/ec.server.crt", "keys/ec.server.key")
	serverKeypair, err := tls.LoadX509KeyPair("keys/ed25519.server.crt", "keys/ed25519.server.key")
	if err != nil {
		fmt.Println("get key pair err:", err)
		return
	}

	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("keys/ec.ca.crt")
	certPool.AppendCertsFromPEM(caCert)

	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{serverKeypair},
		ClientAuth:   tls.RequireAnyClientCert, // require client to provide cert
		ClientCAs:    certPool,
	}
	listener, err := tls.Listen("tcp", "localhost:12345", serverConfig)
	defer toClose1(listener)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	conn, err := listener.Accept() // (net.Listener).Accept() is just accept a tcp conn, it hasn't handshaked yet
	defer toClose1(conn)
	if err != nil {
		fmt.Println("accept err:", err)
	}
	byter := make([]byte, 128)
	n, err := conn.Read(byter) // tls.Conn.Read() or tls.Conn.Write() invoke Handshake() implicitly
	if err != nil && err != io.EOF {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("read: ", string(byter[:n]), conn.LocalAddr().String())
}

func toClose1(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}
}
