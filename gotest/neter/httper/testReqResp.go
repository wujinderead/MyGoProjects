package main

import (
	"bytes"
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
	testRedirectResponse()
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
