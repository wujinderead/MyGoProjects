package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

func main() {
	//testGet()
	//testPost()
	//testPostForm()
	//testHead()
	//testDo()
	testRedirect()
}

func testGet() {
	resp, err := http.Get("https://cn.bing.com/search?q=%E5%96%80%E7%BA%B3%E6%96%AF")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	defer toClose(resp.Body)

	printResponse(resp)
	fmt.Println()

	printTlsConnState(resp.TLS)
	fmt.Println()

	printRequest(resp.Request)
	fmt.Println()

	printTlsConnState(resp.Request.TLS)
	fmt.Println()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("response body:")
	fmt.Println(string(body))
}

func testPost() {
	buf := bytes.NewBufferString("good good good")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", buf)
	if err != nil {
		fmt.Println("post err:", err)
		return
	}
	defer toClose(resp.Body)

	printResponse(resp)
	fmt.Println()

	printTlsConnState(resp.TLS)
	fmt.Println()

	printRequest(resp.Request)
	fmt.Println()

	printTlsConnState(resp.Request.TLS)
	fmt.Println()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("response body:")
	fmt.Println(string(body))
}

func testPostForm() {
	parameters := map[string][]string{
		"country": {"mexico"},
		"cities":  {"mexico city", "morelia", "guadalajara"},
	}
	resp, err := http.PostForm("http://httpbin.org/post", url.Values(parameters))
	if err != nil {
		fmt.Println("post err:", err)
		return
	}
	defer toClose(resp.Body)

	printResponse(resp)
	fmt.Println()

	printTlsConnState(resp.TLS)
	fmt.Println()

	printRequest(resp.Request)
	fmt.Println()

	printTlsConnState(resp.Request.TLS)
	fmt.Println()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("response body:")
	fmt.Println(string(body))
}

func testHead() {
	// only return response head
	resp, err := http.Head("https://cn.bing.com/search?q=not+really")
	if err != nil {
		fmt.Println("head err:", err)
		return
	}
	defer toClose(resp.Body)

	printResponse(resp)
	fmt.Println()

	printTlsConnState(resp.TLS)
	fmt.Println()

	printRequest(resp.Request)
	fmt.Println()

	printTlsConnState(resp.Request.TLS)
	fmt.Println()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("response body:")
	fmt.Println(string(body))
}

func testDo() {
	req := &http.Request{}
	req.Method = "PUT"
	req.URL = &url.URL{
		Scheme: "http",
		Host:   "httpbin.org",
		Path:   "/put",
	}
	req.Proto = "HTTP/1.1"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.Body = ioutil.NopCloser(strings.NewReader("我的天呐"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("head err:", err)
		return
	}
	defer toClose(resp.Body)

	printResponse(resp)
	fmt.Println()

	printTlsConnState(resp.TLS)
	fmt.Println()

	fmt.Println("request:", uintptr(unsafe.Pointer(req)), "resp req:", uintptr(unsafe.Pointer(resp.Request)))
	printRequest(resp.Request)
	fmt.Println()

	printTlsConnState(resp.Request.TLS)
	fmt.Println()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("response body:")
	fmt.Println(string(body))
}

func testRedirect() {
	req := &http.Request{}
	req.URL, _ = url.Parse("http://t.cn/AiYeOD5V")
	req.Method = http.MethodGet
	fmt.Println("=== initial request: ===", req.URL)
	fmt.Println()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// redirect response
			redirectResp := req.Response
			fmt.Println("=== redirect response: ===")
			fmt.Println("status:", redirectResp.Status)
			printTlsConnState(redirectResp.TLS)
			for k, v := range redirectResp.Header {
				fmt.Println("resp header", k, "=", v)
			}
			fmt.Println()

			// redirected request
			fmt.Println("=== redirected request: ===")
			fmt.Println("req url:", req.URL.String())
			for k, v := range req.Header {
				fmt.Println("req header", k, "=", v)
			}
			fmt.Println()
			return nil
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("do err:", err)
		return
	}
	defer toClose(resp.Body)

	fmt.Println("=== final response: ===")
	printResponse(resp)
	fmt.Println()

	printTlsConnState(resp.TLS)
	fmt.Println()
}

func toClose(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}
}

func printTlsConnState(state *tls.ConnectionState) {
	if state == nil {
		return
	}
	fmt.Println("tls Version                    :", state.Version)
	fmt.Println("tls HandshakeComplete          :", state.HandshakeComplete)
	fmt.Println("tls DidResume                  :", state.DidResume)
	fmt.Println("tls CipherSuite                :", cipherSuite[state.CipherSuite])
	fmt.Println("tls NegotiatedProtocol         :", state.NegotiatedProtocol)
	fmt.Println("tls NegotiatedProtocolIsMutual :", state.NegotiatedProtocolIsMutual)
	fmt.Println("tls ServerName                 :", state.ServerName)
	for i := range state.PeerCertificates {
		fmt.Println(i, "issuer:", state.PeerCertificates[i].Issuer.CommonName)
		fmt.Println(i, "subject:", state.PeerCertificates[i].Subject.CommonName)
	}
	for i := range state.VerifiedChains {
		for j := range state.VerifiedChains[i] {
			fmt.Println(i, j, "chain issuer:", state.VerifiedChains[i][j].Issuer.CommonName)
			fmt.Println(i, j, "chain subject:", state.VerifiedChains[i][j].Subject.CommonName)
		}
	}
}

var cipherSuite = map[uint16]string{
	// TLS 1.0 - 1.2 cipher suites.
	0x0005: "TLS_RSA_WITH_RC4_128_SHA",
	0x000a: "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
	0x002f: "TLS_RSA_WITH_AES_128_CBC_SHA",
	0x0035: "TLS_RSA_WITH_AES_256_CBC_SHA",
	0x003c: "TLS_RSA_WITH_AES_128_CBC_SHA256",
	0x009c: "TLS_RSA_WITH_AES_128_GCM_SHA256",
	0x009d: "TLS_RSA_WITH_AES_256_GCM_SHA384",
	0xc007: "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA",
	0xc009: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
	0xc00a: "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
	0xc011: "TLS_ECDHE_RSA_WITH_RC4_128_SHA",
	0xc012: "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA",
	0xc013: "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
	0xc014: "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
	0xc023: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256",
	0xc027: "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
	0xc02f: "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
	0xc02b: "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
	0xc030: "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
	0xc02c: "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
	0xcca8: "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
	0xcca9: "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
	// TLS 1.3 cipher suites.
	0x1301: "TLS_AES_128_GCM_SHA256",
	0x1302: "TLS_AES_256_GCM_SHA384",
	0x1303: "TLS_CHACHA20_POLY1305_SHA256",
	0x5600: "TLS_FALLBACK_SCSV",
}

func printRequest(req *http.Request) {
	fmt.Println("req Method          :", req.Method)
	fmt.Println("req URL             :", req.URL.String())
	fmt.Println("req Proto           :", req.Proto)
	fmt.Println("req ProtoMajor      :", req.ProtoMajor)
	fmt.Println("req ProtoMinor      :", req.ProtoMinor)
	fmt.Println("req Header          :", req.Header)
	fmt.Println("req ContentLength   :", req.ContentLength)
	fmt.Println("req TransferEncoding:", req.TransferEncoding)
	fmt.Println("req Close           :", req.Close)
	fmt.Println("req Host            :", req.Host)
	fmt.Println("req Trailer         :", req.Trailer)
	fmt.Println("req RemoteAddr      :", req.RemoteAddr)
	fmt.Println("req RequestURI      :", req.RequestURI)
	fmt.Println()

	printCookies(req.Cookies())

	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("read request body err:", err)
			return
		}
		fmt.Println("request body:")
		fmt.Println(string(body))
	}
}

func printResponse(resp *http.Response) {
	fmt.Println("resp Status          :", resp.Status)
	fmt.Println("resp StatusCode      :", resp.StatusCode)
	fmt.Println("resp Proto           :", resp.Proto)
	fmt.Println("resp ProtoMajor      :", resp.ProtoMajor)
	fmt.Println("resp ProtoMinor      :", resp.ProtoMinor)
	fmt.Println("resp ContentLength   :", resp.ContentLength)
	fmt.Println("resp TransferEncoding:", resp.TransferEncoding)
	fmt.Println("resp Close:          :", resp.Close)
	fmt.Println("resp Uncompressed    :", resp.Uncompressed)
	location, err := resp.Location()
	if err != nil {
		fmt.Println("resp location err:", err)
	} else {
		fmt.Println("rest Location    :", location.String())
	}
	fmt.Println()

	printCookies(resp.Cookies())

	for k, v := range resp.Header {
		fmt.Println("resp header", k, "=", v)
	}
	for k, v := range resp.Trailer {
		fmt.Println("resp trailer header", k, "=", v)
	}
}

func printCookie(cookie *http.Cookie) {
	fmt.Println("Name      :", cookie.Name)
	fmt.Println("Value     :", cookie.Value)
	fmt.Println("Path      :", cookie.Path)
	fmt.Println("Domain    :", cookie.Domain)
	fmt.Println("Expires   :", cookie.Expires)
	fmt.Println("RawExpires:", cookie.RawExpires)
	fmt.Println("MaxAge    :", cookie.MaxAge)
	fmt.Println("Secure    :", cookie.Secure)
	fmt.Println("HttpOnly  :", cookie.HttpOnly)
	fmt.Println("SameSite  :", cookie.SameSite)
	fmt.Println("Raw       :", cookie.Raw)
	fmt.Println("Unparsed  :", cookie.Unparsed)
}

func printCookies(cookies []*http.Cookie) {
	for i, cookie := range cookies {
		fmt.Println("cookies", i, ":")
		fmt.Println("Name      :", cookie.Name)
		fmt.Println("Value     :", cookie.Value)
		//fmt.Println("Path      :", cookie.Path)
		fmt.Println("Domain    :", cookie.Domain)
		//fmt.Println("Expires   :", cookie.Expires)
		//fmt.Println("RawExpires:", cookie.RawExpires)
		//fmt.Println("MaxAge    :", cookie.MaxAge)
		//fmt.Println("Secure    :", cookie.Secure)
		//fmt.Println("HttpOnly  :", cookie.HttpOnly)
		//fmt.Println("SameSite  :", cookie.SameSite)
		//fmt.Println("Raw       :", cookie.Raw)
		//fmt.Println("Unparsed  :", cookie.Unparsed)
		fmt.Println()
	}
}
