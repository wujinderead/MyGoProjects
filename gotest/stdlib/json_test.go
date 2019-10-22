package stdlib

import (
	"bytes"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
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

func TestMarshalUnmarshal(t *testing.T) {
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

// self-defined marshaller and unmarshaller
type monthYear time.Time

func (m monthYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Year  int `json:"year"`
		Month int `json:"month"`
	}{time.Time(m).Year(), int(time.Time(m).Month())})
}

type inter int

func (i *inter) UnmarshalJSON(in []byte) error {
	mm := struct {
		Year  int `json:"year"`
		Month int `json:"month"`
	}{}
	err := json.Unmarshal(in, &mm)
	if err != nil {
		return err
	}
	*i = inter(mm.Year*100 + mm.Month)
	return nil
}

func TestMarshalUnmarshalJSONInterface(t *testing.T) {
	m := monthYear(time.Now())
	jm, err := json.Marshal(m)
	fmt.Println(err, string(jm))
	jm, err = json.Marshal(&m)
	fmt.Println(err, string(jm))

	i := inter(0)
	err = (&i).UnmarshalJSON(jm)
	fmt.Println(err, i)
}

func TestAnnotation(t *testing.T) {
	type MM struct {
		Alpha   int    `json:"-"`              // "-" ignore this field when parsing or outputting it
		Bravo   int    `json:""`               // "" means no specified name, use "Bravo" as name
		Charlie string `json:"char,omitempty"` // do not output this field when empty
	}
	m := MM{1, 1, ""}
	b, err := json.Marshal(&m)
	fmt.Println(err, string(b)) // charlie not output

	m = MM{1, 1, "aaa"}
	b, err = json.Marshal(&m)
	fmt.Println(err, string(b)) // charlie outputs

	var m1 MM
	err = json.Unmarshal([]byte(`{"Alpha":2, "Bravo":2, "char":"aaa"}`), &m1)
	fmt.Println(err, m1) // got {0 2 aaa}, alpha ignored

	type inn1 struct {
		Mike  MM  `json:"mike,inline"`
		Delta int `json:"delta"`
	}
	in1 := &inn1{m1, 5}
	b, err = json.Marshal(in1)
	fmt.Println(err, string(b)) // wrapped struct: {"mike":{"Bravo":2,"char":"aaa"},"delta":5}

	type inn2 struct {
		MM
		Delta int `json:"delta"`
	}
	in2 := &inn2{m1, 5}
	b, err = json.Marshal(in2)
	fmt.Println(err, string(b)) // unwrapped struct: {"Bravo":2,"char":"aaa","delta":5}

	type mm struct {
		Alpha   int    `json:"-"`
		Bravo   int    `json:""`
		Charlie string `json:"char,omitempty"`
	}
	type inn3 struct {
		mm        // lowercase anonymous field can also be exported
		Delta int `json:"delta"`
	}
	in3 := &inn3{mm{2, 2, "aaa"}, 5}
	b, err = json.Marshal(in3)
	fmt.Println(err, string(b)) // {"Bravo":1,"char":"aaa","delta":5}
}
