package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	//noClientCert()
	//requestClientCert()
	//verifyClientCertIfGiven()
	//requireAndVerifyClientCert()
}

func noClientCert() {
	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("keys/ec.ca.crt")
	certPool.AppendCertsFromPEM(caCert)
	config := &tls.Config{
		ClientAuth: tls.NoClientCert, // not require client cert
		//ClientAuth: tls.RequestClientCert,  // request but not mandatory
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("indexer"))
	})
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("我是李哥"))
	})
	server := http.Server{
		Addr:      "0.0.0.0:8080",
		TLSConfig: config,
	}
	go func() {
		time.Sleep(500 * time.Millisecond)
		for _, v := range []string{
			// ok, no verify server
			"--verify no https://localhost:8080/name",
			// error, verify server but no right ca file, 'tls: unknown certificate authority'
			"--verify yes https://localhost:8080/name",
			// ok, verify server with right ca file
			"--verify keys/ec.ca.crt https://localhost:8080/name",
			// error, because the server hostname is not in the server cert, the client treat server as untrusted
			"--verify keys/ec.ca.crt https://10.19.138.22:8080/name",
		} {
			byter, err := exec.Command("http", strings.Split(v, " ")...).Output()
			fmt.Println("byter:", string(byter), ", err:", err)
		}
		time.Sleep(500 * time.Millisecond)
		os.Exit(1)
	}()
	log.Fatal(server.ListenAndServeTLS("keys/ec.server.crt", "keys/ec.server.key"))
}

func requestClientCert() {
	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("keys/ec.ca.crt")
	certPool.AppendCertsFromPEM(caCert)
	config := &tls.Config{
		ClientAuth: tls.RequireAnyClientCert, // client must provider a cert no matter if it's valid
		ClientCAs:  certPool,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("indexer"))
	})
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("我是李哥"))
	})
	server := http.Server{
		Addr:      ":8080",
		TLSConfig: config,
	}
	go func() {
		time.Sleep(500 * time.Millisecond)
		for _, v := range []string{
			// error, tls: client didn't provide a certificate
			"--verify no https://localhost:8080/name",
			// ok
			"--cert keys/ec.client.crt --cert-key keys/ec.client.key --verify no https://localhost:8080/name",
			// ok, even wrong cert
			"--cert keys/rsa.client.crt --cert-key keys/rsa.client.key --verify no https://localhost:8080/name",
			// cmd error, must provide both cert and key to httpie
			"--cert keys/rsa.client.crt --verify no https://localhost:8080/name",
		} {
			byter, err := exec.Command("http", strings.Split(v, " ")...).Output()
			fmt.Println("byter:", string(byter), ", err:", err)
		}
		time.Sleep(500 * time.Millisecond)
		os.Exit(1)
	}()
	log.Fatal(server.ListenAndServeTLS("keys/ec.server.crt", "keys/ec.server.key"))
}

func verifyClientCertIfGiven() {
	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("keys/ec.ca.crt")
	certPool.AppendCertsFromPEM(caCert)
	config := &tls.Config{
		ClientAuth: tls.VerifyClientCertIfGiven, // verify client cert if given
		ClientCAs:  certPool,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("indexer"))
	})
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("我是李哥"))
	})
	server := http.Server{
		Addr:      ":8080",
		TLSConfig: config,
	}
	go func() {
		time.Sleep(500 * time.Millisecond)
		for _, v := range []string{
			// ok, client don't provide cert
			"--verify no https://localhost:8080/name",
			// ok, client provide right cert
			"--cert keys/ec.client.crt --cert-key keys/ec.client.key --verify no https://localhost:8080/name",
			// err, client privode wrong cert, 'tls: failed to verify client's certificate: x509'
			"--cert keys/rsa.client.crt --cert-key keys/rsa.client.key --verify no https://localhost:8080/name",
		} {
			byter, err := exec.Command("http", strings.Split(v, " ")...).Output()
			fmt.Println("byter:", string(byter), ", err:", err)
		}
		time.Sleep(500 * time.Millisecond)
		os.Exit(1)
	}()
	log.Fatal(server.ListenAndServeTLS("keys/ec.server.crt", "keys/ec.server.key"))
}

func requireAndVerifyClientCert() {
	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("keys/ec.ca.crt")
	certPool.AppendCertsFromPEM(caCert)
	config := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert, // client is required to give a right cert
		ClientCAs:  certPool,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("indexer"))
	})
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("我是李哥"))
	})
	server := http.Server{
		Addr:      ":8080",
		TLSConfig: config,
	}
	go func() {
		time.Sleep(500 * time.Millisecond)
		for _, v := range []string{
			// error, client don't provide cert, 'tls: client didn't provide a certificate'
			"--verify no https://localhost:8080/name",
			// ok, client provide right cert
			"--cert keys/ec.client.crt --cert-key keys/ec.client.key --verify no https://localhost:8080/name",
			// err, client privode wrong cert, 'tls: failed to verify client's certificate: x509'
			"--cert keys/rsa.client.crt --cert-key keys/rsa.client.key --verify no https://localhost:8080/name",
		} {
			byter, err := exec.Command("http", strings.Split(v, " ")...).Output()
			fmt.Println("byter:", string(byter), ", err:", err)
		}
		time.Sleep(500 * time.Millisecond)
		os.Exit(1)
	}()
	log.Fatal(server.ListenAndServeTLS("keys/ec.server.crt", "keys/ec.server.key"))
}

func toClose(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}
}
