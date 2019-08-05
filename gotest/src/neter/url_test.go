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
}

func TestUrl(t *testing.T) {
	purl, err := url.Parse("mysql://xzy:123456@10.2.2.2:3306/db/schema?aaa=bbb&ccc=ddd#eee=%E6%9D%8E")
	if err != nil {
		t.Error("parse url err:", err)
		return
	}
	pass, set := purl.User.Password()
	fmt.Println("Scheme    :", purl.Scheme    )
	fmt.Println("Opaque    :", purl.Opaque    )
	fmt.Println("User      :", purl.User.String())
	fmt.Println("User name :", purl.User.Username())
	fmt.Println("User pass :", pass, set)
	fmt.Println("Host      :", purl.Host      )
	fmt.Println("Path      :", purl.Path      )
	fmt.Println("RawPath   :", purl.RawPath   )
	fmt.Println("ForceQuery:", purl.ForceQuery)
	fmt.Println("RawQuery  :", purl.RawQuery  )
	fmt.Println("Fragment  :", purl.Fragment  )
}