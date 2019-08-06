package neter

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"testing"
)

func TestEscape(t *testing.T) {
	// escape to utf8 encoding
	for _, s := range "~`!@#$%^&*()-_=+[]{}\\|:;\"'<>?,./ \t李王Ω≤≥µ" {
		fmt.Println(string(s), url.PathEscape(string(s)), hex.EncodeToString([]byte(string(s))))
	}
	fmt.Println(url.PathEscape("http://baidu.com/search?k=aaa&b=ccc;s"))

	// unescape url
	unescaped, err := url.PathUnescape("http:%2F%2Fbaidun.com%2Fsearch%3Fk=aaa&b=ccc%3bs")
	if err != nil {
		t.Error("path unescape err:", err)
		return
	}
	fmt.Println(unescaped)

	// escape query url
	// PathEscape won't escape `:`, `=`, `+`; but QueryEscape will.
	fmt.Println(url.QueryEscape("http://baidu.com/search?k=aaa&b=ccc+李哥;s"))
	fmt.Println()

	// unescape query url
	str, err := url.QueryUnescape("http%3A%2F%2Fbaidu.com%2Fsearch%3Fk%3Daaa%26b%3Dccc%2B%E6%9D%8E%E5%93%A5%3Bs")
	if err != nil {
		t.Error("query unescape err:", err)
		return
	}
	fmt.Println("query unescape:", str)
}

func TestUrl(t *testing.T) {
	purl, err := url.Parse("mysql://xzy:123456@10.2.2.2:3306/db/schema?aaa=bbb+ddd&ccc=李哥#table1")
	if err != nil {
		t.Error("parse url err:", err)
		return
	}
	pass, set := purl.User.Password()
	fmt.Println("Scheme    :", purl.Scheme)
	fmt.Println("Opaque    :", purl.Opaque)
	fmt.Println("User      :", purl.User.String())
	fmt.Println("User name :", purl.User.Username())
	fmt.Println("User pass :", pass, set)
	fmt.Println("Host      :", purl.Host)
	fmt.Println("Path      :", purl.Path)
	fmt.Println("RawPath   :", purl.RawPath)
	fmt.Println("ForceQuery:", purl.ForceQuery)
	fmt.Println("RawQuery  :", purl.RawQuery)
	fmt.Println("Fragment  :", purl.Fragment)

	fmt.Println("str:", purl.String())
	fmt.Println("escape path:", purl.EscapedPath())
	fmt.Println("hostname:", purl.Hostname())
	fmt.Println("is abs:", purl.IsAbs())
	fmt.Println("port:", purl.Port())
	fmt.Println("request uri:", purl.RequestURI()) // things after host, /db/schema?aaa=bbb+ddd&ccc=李哥

	// url.Values, a wrapper of map[string][]string
	fmt.Println("query:", purl.Query())                 // map[aaa:[bbb ddd] ccc:[李哥]]
	fmt.Println("query encode:", purl.Query().Encode()) // aaa=bbb+ddd&ccc=%E6%9D%8E%E5%93%A5
	fmt.Println()

	// can parse a relative path to full path based on purl
	rurl, err := purl.Parse("/another/path/query?aaa=bbb#frag2")
	if err != nil {
		t.Error("reference err:", err)
		return
	}
	fmt.Println("url parsed:", rurl.String()) // mysql://xzy:123456@10.2.2.2:3306/another/path/query?aaa=bbb#frag2
	fmt.Println()

	// unmarshal a url
	u := &url.URL{}
	err = u.UnmarshalBinary([]byte("https://example.org/foo?aaa=bbb+ddd&ccc=%e6%9d%8e%e5%93%a5#table1"))
	if err != nil {
		t.Error("url unmarshal err:", err)
		return
	}
	fmt.Println("unmarshal str", u.String())
	fmt.Println("unmarshal str", u.RequestURI())
	fmt.Println("unmarshal query", u.Query())
	fmt.Println("unmarshal query encode", u.Query().Encode())
	fmt.Println()

	// url fields
	curl := &url.URL{
		Scheme:   "https",
		Host:     "cn.bing.com",
		Path:     "/academic/search",
		RawQuery: "q=elliptic+curve+pairing",
	}
	fmt.Println(curl.String())
}
