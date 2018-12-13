package stdlib

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFormatInt(t *testing.T) {
	// i to a
	fmt.Println(strconv.FormatInt(12346, 10))
	fmt.Println(strconv.FormatInt(12346, 2))
	fmt.Println(strconv.FormatInt(12346, 8))
	fmt.Println(strconv.FormatInt(12346, 16))
	fmt.Println(strconv.Itoa(12346))

	// a to i
	fmt.Println(strconv.Atoi("123456"))
	fmt.Println(strconv.ParseInt("-1a2b3d4e", 16, 0))
	fmt.Println(strconv.ParseInt("0777", 0, 0))

	// float, bool
	fmt.Println(strconv.FormatFloat(12.3456, 'E', 4, 64))
	fmt.Println(strconv.FormatFloat(12.3456, 'f', 4, 64))
	fmt.Println(strconv.FormatFloat(12.3456, 'g', 4, 64))
	fmt.Println(strconv.FormatFloat(12.3456, 'b', 4, 64))
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseFloat("12.3456", 64))

	// quote
	fmt.Println(strconv.Quote("dasda"))
	fmt.Println(strconv.QuoteRune('阿'))
	fmt.Println(strconv.QuoteRuneToASCII('阿'))
	fmt.Println(strconv.QuoteToASCII("我是黄旭东"))
	fmt.Println(strconv.QuoteToGraphic("我是黄旭东"))

	// graphic
	fmt.Println("graphic a: ", strconv.IsGraphic('a'))
	fmt.Println("graphic 我: ", strconv.IsGraphic('我'))
	fmt.Println("graphic 😀: ", strconv.IsGraphic('😀'))

	// append
	buf := make([]byte, 0)
	buf = strconv.AppendQuote(buf, "我是黄旭东")
	buf = strconv.AppendQuoteRune(buf, '黄')
	buf = strconv.AppendBool(buf, true)
	buf = strconv.AppendInt(buf, 123456, 16)
	buf = strconv.AppendFloat(buf, 12.3456, 'e', 4, 64)
	fmt.Println(string(buf))
}
