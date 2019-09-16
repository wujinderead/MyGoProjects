package stdlib

import (
	"bytes"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoderDecoder(t *testing.T) {
	// encoder
	w := os.Stdout
	en := json.NewEncoder(w)
	en.SetEscapeHTML(true)
	en.SetIndent("", "  ")
	curve := elliptic.P224()
	err := en.Encode(curve)
	fmt.Println(err)
	err = en.Encode(map[string]interface{}{"aaa": "aa&aa<aa>aa", "bbb": 1, "ccc": true})
	fmt.Println(err)

	// decoder
	r := bytes.NewReader([]byte(`{"aaa":"数据", "bbb":1, "ccc": true}{"bbb":2, "aaa":"方法", "ccc": false}`))
	de := json.NewDecoder(r)
	type pojo struct {
		Aaa string `json:"aaa"`
		Bbb int    `json:"bbb"`
		Ccc bool   `json:"ccc"`
	}
	pp := new(pojo)
	err = de.Decode(pp)
	fmt.Println(err, pp, de.More())
	if false {
		ps := pojo{}
		err = de.Decode(ps)  // got err `json: Unmarshal(non-pointer utils.pojo)`
		fmt.Println(err, ps) // must use pointer
	}
	var mp map[string]interface{}
	err = de.Decode(&mp)
	fmt.Println(err, mp, de.More())

	// token
	r = bytes.NewReader([]byte(`{"aaa":["数据", "data"], "bbb":{"ddd":1, "arr":[]}, "ccc": true}`))
	de = json.NewDecoder(r)
	for {
		token, err := de.Token()
		if token == nil || err != nil {
			break
		}
		// token is either delim (big parentheses, middle parentheses, commas), or identifiers, or real data
		fmt.Println(token)
	}
}

func TestJson(t *testing.T) {
	jsoner := []byte(`{
	"name": "james",
	"age": 35,
	"teams": ["cavaliers", "heats", "lakers"]
	}`)
	buf := new(bytes.Buffer)
	err := json.Compact(buf, jsoner) // compact json, remove unnecessary indents, then append to buffer
	fmt.Println(err, buf.String())

	buf.Reset()
	err = json.Indent(buf, jsoner, "", "  ") // add proper indents, then append to buffer
	fmt.Println(err, buf.String())

	type player struct {
		Name string   `json:"name"`
		Age  int      `json:"age"`
		Ts   []string `json:"teams"`
	}
	p := new(player)
	err = json.Unmarshal(jsoner, p) // unmarshal from json to struct
	fmt.Println(err, p)

	byter, err := json.Marshal(p) // marshal struct to json
	fmt.Println(err, string(byter))

	byter, err = json.MarshalIndent(p, "", "  ") // marshal struct to json with indents
	fmt.Println(err, string(byter))

	fmt.Println("valid:", json.Valid(jsoner)) // check if it is valid json
	jsoner[1] = ','
	fmt.Println("valid:", json.Valid(jsoner))
}
