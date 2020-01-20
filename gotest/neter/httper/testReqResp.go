package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unsafe"

	"gotest/neter/httper/utils"
)

func main() {
	//testGet()
	//testPost()
	//testPostForm()
	//testHead()
	//testDo()
	//testRedirectResponse()
	testWriteRequest()
	//testChunkedResponse()
	//testChunkedResponseGo()
}

func testGet() {
	resp, err := http.Get("https://cn.bing.com/search?q=%E5%96%80%E7%BA%B3%E6%96%AF")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	defer utils.ToClose(resp.Body)

	utils.PrintResponse(resp)
	fmt.Println()

	utils.PrintTlsConnState(resp.TLS)
	fmt.Println()

	utils.PrintRequest(resp.Request)
	fmt.Println()

	utils.PrintTlsConnState(resp.Request.TLS)
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
	defer utils.ToClose(resp.Body)

	utils.PrintResponse(resp)
	fmt.Println()

	utils.PrintTlsConnState(resp.TLS)
	fmt.Println()

	utils.PrintRequest(resp.Request)
	fmt.Println()

	utils.PrintTlsConnState(resp.Request.TLS)
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
	defer utils.ToClose(resp.Body)

	utils.PrintResponse(resp)
	fmt.Println()

	utils.PrintTlsConnState(resp.TLS)
	fmt.Println()

	utils.PrintRequest(resp.Request)
	fmt.Println()

	utils.PrintTlsConnState(resp.Request.TLS)
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
	defer utils.ToClose(resp.Body)

	utils.PrintResponse(resp)
	fmt.Println()

	utils.PrintTlsConnState(resp.TLS)
	fmt.Println()

	utils.PrintRequest(resp.Request)
	fmt.Println()

	utils.PrintTlsConnState(resp.Request.TLS)
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
		fmt.Println("do err:", err)
		return
	}
	defer utils.ToClose(resp.Body)

	utils.PrintResponse(resp)
	fmt.Println()

	utils.PrintTlsConnState(resp.TLS)
	fmt.Println()

	fmt.Println("request:", uintptr(unsafe.Pointer(req)), "resp req:", uintptr(unsafe.Pointer(resp.Request)))
	utils.PrintRequest(resp.Request)
	fmt.Println()

	utils.PrintTlsConnState(resp.Request.TLS)
	fmt.Println()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("response body:")
	fmt.Println(string(body))
}

func testRedirectResponse() {
	req := &http.Request{}
	req.URL, _ = url.Parse("http://t.cn/AiYeOD5V")
	req.Method = http.MethodGet
	fmt.Println("=== initial request: ===", req.URL)
	fmt.Println()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// last 302 response
			redirectResp := req.Response
			fmt.Println("=== redirect response: ===")
			fmt.Println("status:", redirectResp.Status)
			utils.PrintTlsConnState(redirectResp.TLS)
			for k, v := range redirectResp.Header {
				fmt.Println("resp header", k, "=", v)
			}
			fmt.Println()

			// request to redirected destination
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
	defer utils.ToClose(resp.Body)

	fmt.Println("=== final response: ===")
	utils.PrintResponse(resp)
	fmt.Println()

	utils.PrintTlsConnState(resp.TLS)
	fmt.Println()
}

// serialize request to bytes and write to writer
func testWriteRequest() {
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
	req.Body = ioutil.NopCloser(strings.NewReader(`我的天呐12345678901234567`))
	req.Header = make(map[string][]string)
	req.Header["Accept"] = []string{"text/html", "application/xhtml+xml", "application/xml", "image/webp"}
	req.Header["Accept-Language"] = []string{"zh-CN,zh;q=0.9"}
	req.Header["Sec-Fetch-Mode"] = []string{"navigate"}
	req.Header["Sec-Fetch-Site"] = []string{"same-origin"}
	req.Header["Accept-Language"] = []string{"zh-CN,zh;q=0.9"}
	req.Header["Transfer-Encoding"] = []string{"chunked"}
	req.Trailer = make(map[string][]string)
	req.Trailer["key1"] = []string{"value1", "value2"}
	req.Trailer["key2"] = []string{"value3"}
	req.AddCookie(&http.Cookie{Name: "username", Value: "lgq", Path: "/", Domain: ".baidu.com"})
	req.AddCookie(&http.Cookie{Name: "password", Value: "12345678"})

	buf := &bytes.Buffer{}
	_ = req.Write(buf)
	fmt.Println(buf.Bytes())
	fmt.Printf("%q\n", string(buf.Bytes()))

	fmt.Println("================")
	fmt.Println()
	buf.Reset()
	req.Body = nil
	_ = req.Write(buf)
	fmt.Println(buf.Bytes())
	fmt.Println(string(buf.Bytes()))
	fmt.Printf("%q\n", string(buf.Bytes()))
}

/*
request format:
PUT /put HTTP/1.1\r\n
Host: httpbin.org\r\n
User-Agent: Go-http-client/1.1\r\n
Transfer-Encoding: chunked\r\n
Trailer: Key1,Key2\r\n
Accept: text/html\r\n
Accept: application/xhtml+xml\r\n
Accept: application/xml\r\n
Accept: image/webp\r\n
Accept-Language: zh-CN,zh;q=0.9\r\n
Cookie: username=lgq; password=12345678\r\n
Sec-Fetch-Mode: navigate\r\n
Sec-Fetch-Site: same-origin\r\n
\r\n
1d\r\n
我的天呐12345678901234567\r\n
0\r\n
key1: value1\r\n
key1: value2\r\n
key2: value3\r\n
\r\n"

*/

func testChunkedResponse() {
	conn, err := tls.Dial("tcp", "www.baidu.com:443", nil)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()

	// write a http request to the tls (tcp) conn
	bw := bufio.NewWriter(conn)
	_, _ = bw.WriteString("GET / HTTP/1.1\r\n")
	_, _ = bw.WriteString("Host: www.baidu.com\r\n")
	_, _ = bw.WriteString("Connection: keep-alive\r\n")
	_, _ = bw.WriteString("Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\r\n")
	// do not compress
	// _, _ = bw.WriteString("Accept-Encoding: gzip, deflate, br\r\n")
	_, _ = bw.WriteString("Accept-Language: zh-CN,zh;q=0.9\r\n")
	// set user-agent as browser, otherwise server won't sent chunked data
	_, _ = bw.WriteString("User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36\r\n")
	_, _ = bw.WriteString("\r\n")
	fmt.Println("flush:", bw.Flush())

	// read the http response from the tls (tcp) conn
	buf := make([]byte, 16*1024)
	var ns []int
	i := 0
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println("====read:", i, n)
		ns = append(ns, n)
		/*if n<20 {
			fmt.Println(buf[:n])
		} else {
			fmt.Println(buf[n-10:n])
		}*/
		fmt.Printf("%q\n\n", string(buf[:n]))
		if n >= 5 && string(buf[n-5:n]) == "0\r\n\r\n" {
			break
		}
	}
	fmt.Println(ns)
}

/*
response format:
HTTP/1.1 200 OK\r\n
Header1: xxx\r\n
Header2: xxx\r\n
\r\n
9ee\r\n                                // chunk length
chunked-body...\r\n
16a0\r\n
chunked-body...\r\n
0                                      // terminal
\r\n\r\n
*/

func testChunkedResponseGo() {
	req := &http.Request{}
	req.URL, _ = url.Parse("https://www.baidu.com")
	req.Header = make(map[string][]string)
	req.Header["Host"] = []string{"www.baidu.com"}
	req.Header["Connection"] = []string{"keep-alive"}
	req.Header["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"}
	req.Header["Accept-Language"] = []string{"zh-CN,zh;q=0.9"}
	req.Header["User-Agent"] = []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("request err:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	fmt.Println(resp.Trailer)
	fmt.Println(resp.TransferEncoding)
}
