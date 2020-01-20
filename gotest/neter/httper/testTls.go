package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/http2"
	"gopkg.in/yaml.v2"
)

func main() {
	//testReadKubeApiserverHttp2()
	//testReadKubeApiserverHttp11()
	//testBingHttp11()
}

func testBingHttp11() {
	conn, err := tls.Dial("tcp", "cn.bing.com:443", nil)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()
	err = conn.Handshake()
	if err != nil {
		fmt.Println("handshake err:", err)
		return
	}
	fmt.Println("tls conn local:", conn.LocalAddr(), "remote:", conn.RemoteAddr())

	// through bing is a http2 server, it can still handle http1.1 request
	w := bufio.NewWriter(conn)
	_, _ = w.WriteString("GET /search?q=http2 HTTP/1.1\r\n")
	_, _ = w.WriteString("HOST: cn.bing.com\r\n")
	_, _ = w.WriteString("\r\n")
	_ = w.Flush()
	buf := make([]byte, 16*1024)
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("===========", n)
		fmt.Println(string(buf[:n]))
	}
}

func testReadKubeApiserverHttp11() {
	key, cert, cacert, serverurl := readKubeAdminCerts()
	kubeAdminKeypair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		fmt.Println("get key pair err:", err)
		return
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cacert)

	// tls handshake
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{kubeAdminKeypair},
		RootCAs:      certPool,
	}
	conn, err := tls.Dial("tcp", serverurl[8:], tlsConfig) // serverurl[8:] to strip "https://"
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()
	err = conn.Handshake()
	if err != nil {
		fmt.Println("handshake err:", err)
		return
	}
	fmt.Println("tls conn local:", conn.LocalAddr(), "remote:", conn.RemoteAddr())

	// write http1.1 request manually
	w := bufio.NewWriter(conn)
	_, _ = w.WriteString("GET /api/v1/nodes HTTP/1.1\r\n")
	_, _ = w.WriteString("HOST: " + serverurl[8:] + "\r\n")
	_, _ = w.WriteString("\r\n")
	_ = w.Flush()
	buf := make([]byte, 16*1024)
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func testReadKubeApiserverHttp2() {
	key, cert, cacert, serverurl := readKubeAdminCerts()
	kubeAdminKeypair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		fmt.Println("get key pair err:", err)
		return
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cacert)

	// to install client certificates for http.Client, configure http.Transport.TLSClientConfig
	tp := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{kubeAdminKeypair},
			RootCAs:      certPool,
		},
	}
	// need to configure if we want http2 client, otherwise can only use http1.1
	err = http2.ConfigureTransport(tp)
	fmt.Println("configure transport:", err)
	client := &http.Client{Transport: tp}

	nodesurl, _ := url.Parse(serverurl + "/api/v1/nodes")
	req := &http.Request{
		URL: nodesurl,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request err:", err)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		fmt.Println(k, ":", v)
	}
	fmt.Println("req proto:", resp.Request.Proto)
	fmt.Println("resp proto:", resp.Proto)
	fmt.Println("status:", resp.Status)
	fmt.Println("trailer:", resp.Trailer)
	fmt.Println("transfer encoding:", resp.TransferEncoding)
	buf := make([]byte, 64*1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read err:", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func readKubeAdminCerts() (key, cert, cacert []byte, serverurl string) {
	kubeAdminConfigPath := "/home/xzy/.kube/config"
	clustername := "kubedog"
	username := "kubernetes-admin"
	f, err := os.Open(kubeAdminConfigPath)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()
	config := &struct {
		ApiVersion string `yaml:"apiVersion,omitempty"`
		Clusters   []struct {
			Name    string
			Cluster struct {
				Server                   string
				CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
			}
		}
		Users []struct {
			Name string
			User struct {
				ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
				ClientKeyData         string `yaml:"client-key-data,omitempty"`
			}
		}
	}{}
	err = yaml.NewDecoder(f).Decode(config)
	if err != nil {
		fmt.Println("unmarshal err:", err)
		return
	}
	for i := range config.Clusters {
		if config.Clusters[i].Name == clustername {
			serverurl = config.Clusters[i].Cluster.Server
			cacert, err = base64.StdEncoding.DecodeString(config.Clusters[i].Cluster.CertificateAuthorityData)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	for i := range config.Users {
		if config.Users[i].Name == username {
			key, err = base64.StdEncoding.DecodeString(config.Users[i].User.ClientKeyData)
			if err != nil {
				fmt.Println(err)
			}
			cert, err = base64.StdEncoding.DecodeString(config.Users[i].User.ClientCertificateData)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}
