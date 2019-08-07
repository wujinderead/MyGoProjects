package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	//testHandleServe()
	//testServeFile()
	//testServeContent()
	testServeTls()
	//createKetAndCert()
}

func testHandleServe() {
	// http get :8080/hello?name=aaa
	// http -f post :8080/hello name=aaa
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet && request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// parse parameter in url, only parse body as form when request is post, put, or patch
		_ = request.ParseForm()
		name := request.Form.Get("name")
		if name == "" {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte("please provide 'name' parameter."))
			return
		}
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("hello, "))
		_, _ = writer.Write([]byte(name))
		return
	})
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("index"))
		return
	}
	// http get :8080/
	http.Handle("/", http.HandlerFunc(indexHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testServeFile() {
	http.HandleFunc("/files/txt", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(writer, request, "/home/xzy/conn.test7.pipeline.sh")
		return
	})
	http.HandleFunc("/files/pdf", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(writer, request, "/home/xzy/图片/schoof algorithm.pdf")
		return
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testServeContent() {
	// ServeContent deal with modtime and io.ReadSeeker, so that the server can handle cache-control
	// or 'resume from break point' request, such as "If-Modified-Since" and "Range".
	// the responses may be like:
	// 304 not modified (the content hasn't been modified, browser can use cache),
	// 206 partial content (use io.ReadSeeker to randomized read partial data).
	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		// os.File is a io.ReadSeeker that support randomized access
		filename := "/usr/local/go/src/runtime/runtime2.go"
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("open files err")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stat, err := file.Stat()
		if err != nil {
			fmt.Println("stat files err")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.ServeContent(w, r, "runtime2.go", stat.ModTime(), file)
	})
	http.HandleFunc("/content/somefile", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/neter/httper/index.html")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	// http :8080/content 'Range:bytes=0-100,200-300'
	// http :8080/content 'If-Modified-Since:Tue, 11 Jun 2019 17:27:51 GMT'
}

func testServeTls() {
	key, cert, err := createKeyAndCert()
	if err != nil {
		fmt.Println("create key cert err:", err.Error())
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/neter/httper/index.html")
	})
	log.Fatal(http.ListenAndServeTLS("10.20.38.188:8080", cert, key, nil))
}

func createKeyAndCert() (keyFile, certFile string, err error) {
	// generate ca ec key
	caKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("generate caKey err:", err)
		return
	}

	// generate ca self certificate
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		fmt.Println("failed to generate serial number:", err)
		return
	}
	//fmt.Println("ca serial number: ", hex.EncodeToString(serialNumber.Bytes()))

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"SelfCA"},
		},
		NotBefore: time.Now().Add(-time.Hour * 24 * 365),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// ca self sign
	caCertBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, caKey.Public(), caKey)
	if err != nil {
		fmt.Println("ca self sign certificate:", err)
		return
	}

	// get caCert instance
	caCert, err := x509.ParseCertificate(caCertBytes)
	if err != nil {
		fmt.Println("parse certificate err:", err)
		return
	}

	// generate server key
	serverKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("generate server key err:", err)
		return
	}

	serverTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"MyServer"},
		},
		NotBefore: time.Now().Add(-time.Hour * 24 * 365),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
		IPAddresses:           []net.IP{net.ParseIP("10.20.38.188"), net.ParseIP("127.0.0.1")},
		DNSNames:              []string{"node1", "localhost"},
	}

	// ca sign server cert
	serverCertBytes, err := x509.CreateCertificate(rand.Reader, &serverTemplate, caCert, serverKey.Public(), caKey)
	if err != nil {
		fmt.Println("sign server cert err:", err)
		return
	}

	// write server key and cert to pem file
	serverKeyBytes, err := x509.MarshalECPrivateKey(serverKey)
	if err != nil {
		fmt.Println("marshal server key err:", err)
		return
	}

	// write to file as pem
	certFile = "/tmp/server.cert"
	certF, err := os.OpenFile(certFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open cert file err:", err)
		return
	}
	defer toClose(certF)

	certB := &pem.Block{Type: "CERTIFICATE", Bytes: serverCertBytes}
	err = pem.Encode(certF, certB)
	if err != nil {
		fmt.Println("write cert pem err:", err)
		return
	}

	keyFile = "/tmp/server.key"
	keyF, err := os.OpenFile(keyFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open key file err:", err)
		return
	}
	defer toClose(keyF)

	keyB := &pem.Block{Type: "EC PRIVATE KEY", Bytes: serverKeyBytes}
	err = pem.Encode(keyF, keyB)
	if err != nil {
		fmt.Println("write key pem err:", err)
		return
	}
	return
}

func toClose(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}
}
