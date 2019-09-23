package golangx

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256) // generate a new token with empty MapClaims
	fmt.Println("header: ", token.Header)
	signed, _ := token.SignedString([]byte("123456")) // sign the token with hmac key "123456"
	fmt.Println("signed token:", signed)
	fmt.Println()

	// parse the token and validate it
	tk, err := jwt.Parse(signed, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("123456"), nil
	})
	fmt.Println("err:", err)
	fmt.Println("valid:", tk.Valid)
	fmt.Println("signature:", tk.Signature)
	fmt.Println("raw:", tk.Raw)
	fmt.Println("header:", tk.Header)

	parts := strings.Split(signed, ".")
	header, err := jwt.DecodeSegment(parts[0])
	fmt.Println(err, string(header))
	payload, err := jwt.DecodeSegment(parts[1])
	fmt.Println(err, string(payload))
	sig, err := jwt.DecodeSegment(parts[2])
	fmt.Println(err, hex.EncodeToString(sig))
}

func TestNewWithClaims(t *testing.T) {
	type user struct { // custom claim
		Name               string `json:"name"` // field name must be capital to be exported
		Id                 int    `json:"Id"`
		Location           string `json:"loc"`
		jwt.StandardClaims        // make struct with standard claims to implement the Claims interface
	}
	myuser := &user{Name: "lgq", Id: 1, Location: "shanghai"}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myuser) // generate a new token with empty claims
	token.Header["encoding"] = "RawURL"                        // add custom header field
	toBeSigned, _ := token.SigningString()                     // the string to be singed: headerBase64+"."+payloadBase64
	signed, _ := token.SignedString([]byte("123456"))          // sign the token with hmac key "123456"
	fmt.Println("to be signed:", toBeSigned)
	fmt.Println("signed token:", signed) // the generated JWT
	fmt.Println()

	// parse the token and validate it
	tk, err := jwt.Parse(signed, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("123456"), nil
	})
	fmt.Println("err:", err)
	fmt.Println("valid:", tk.Valid)
	fmt.Println("signature:", tk.Signature)
	fmt.Println("raw:", tk.Raw)
	fmt.Println("header:", tk.Header) // with custom header fields
	fmt.Println()

	// parse the token and marshal custom claims
	myu := &user{}
	tk, err = jwt.ParseWithClaims(signed, myu, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("123456"), nil
	})
	fmt.Println("parsed user:", myu) // parse payload json to get custom claims
	fmt.Println("err:", err)
	fmt.Println("valid:", tk.Valid)
	fmt.Println("signature:", tk.Signature)
	fmt.Println("raw:", tk.Raw)
	fmt.Println("header:", tk.Header)
	fmt.Println()

	// test hmac
	parts := strings.Split(signed, ".")
	headerB64 := parts[0]
	payloadB64 := parts[1]
	sigB64 := parts[2]
	header, _ := base64.RawURLEncoding.DecodeString(headerB64)
	payload, _ := base64.RawURLEncoding.DecodeString(payloadB64)
	sig, _ := base64.RawURLEncoding.DecodeString(sigB64)
	fmt.Println(string(header))          // header json: {"alg":"HS256","encoding":"RawURL","typ":"JWT"}
	fmt.Println(string(payload))         // payload json: {"name":"lgq","Id":1,"loc":"shanghai"}
	fmt.Println(hex.EncodeToString(sig)) // sha256-hmac signature

	// signature = sha256-hmac( base64(jsonMarshal(header)) + '.' + base64(jsonMarshal(header)), hmac-password )
	hmacer := hmac.New(crypto.SHA256.New, []byte("123456"))
	hmacer.Write([]byte(headerB64))
	hmacer.Write([]byte{'.'})
	hmacer.Write([]byte(payloadB64))
	sum := hmacer.Sum(nil)
	fmt.Println(hex.EncodeToString(sum))
	fmt.Println(bytes.Equal(sig, sum)) // should be equal
}

func TestJwtEcdsa(t *testing.T) {
	// can also be ES384, ES256
	token := jwt.New(jwt.SigningMethodES512)     // generate a new token with empty claims
	token.Header["usage"] = "authorization"      // add custom header
	token.Claims.(jwt.MapClaims)["name"] = "lgq" // add payload fields to MapClaims
	token.Claims.(jwt.MapClaims)["location"] = "SHA"

	// the key should be compatible with SigningMethod. i.e., ES256 with P256, ES284 with P384, ES512 with P521
	priv, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	// put the public key to payload
	token.Claims.(jwt.MapClaims)["pubX"] = base64.URLEncoding.EncodeToString(priv.X.Bytes())
	token.Claims.(jwt.MapClaims)["pubY"] = base64.URLEncoding.EncodeToString(priv.Y.Bytes())

	signed, _ := token.SignedString(priv) // sign the token with ecdsa p256 private keys
	fmt.Println("signed token:", signed)  // the generated JWT
	fmt.Println()

	// parse the token and validate it
	tk, err := jwt.Parse(signed, func(token *jwt.Token) (i interface{}, e error) {
		// get ecdsa public key from jwt
		x, err := base64.URLEncoding.DecodeString(token.Claims.(jwt.MapClaims)["pubX"].(string))
		if err != nil {
			return nil, errors.New("decode public key err:" + err.Error())
		}
		y, err := base64.URLEncoding.DecodeString(token.Claims.(jwt.MapClaims)["pubY"].(string))
		if err != nil {
			return nil, errors.New("decode public key err:" + err.Error())
		}
		pub := &ecdsa.PublicKey{X: new(big.Int).SetBytes(x), Y: new(big.Int).SetBytes(y)}
		switch token.Header["alg"] {
		case "ES256":
			pub.Curve = elliptic.P256()
		case "ES384":
			pub.Curve = elliptic.P384()
		case "ES512":
			pub.Curve = elliptic.P521()
		default:
			return nil, errors.New("not ecdsa public key")
		}
		// use public key to validate
		return pub, nil
	})
	fmt.Println("err:", err)
	fmt.Println("valid:", tk.Valid)
	fmt.Println("signature:", tk.Signature)
	fmt.Println("raw:", tk.Raw)
	fmt.Println("header:", tk.Header) // with custom header fields
	fmt.Println()

	// display payload
	segments := strings.Split(signed, ".")
	headerB64 := segments[0]
	payloadB64 := segments[1]
	sigB64 := segments[2]
	header, _ := base64.RawURLEncoding.DecodeString(headerB64)
	payload, _ := base64.RawURLEncoding.DecodeString(payloadB64)
	sig, _ := base64.RawURLEncoding.DecodeString(sigB64)
	fmt.Println(string(header))          // header json: {"alg":"HS256","encoding":"RawURL","typ":"JWT"}
	fmt.Println(string(payload))         // payload json: {"name":"lgq","Id":1,"loc":"shanghai"}
	fmt.Println(hex.EncodeToString(sig)) // ecdsa signature
}

func TestJwtExpire(t *testing.T) {
	// can also be ES384, ES256
	type tve struct {
		Name string `json:"name"`
		jwt.StandardClaims
	}
	tv := &tve{Name: "lige"}
	tv.StandardClaims.ExpiresAt = time.Now().Add(3 * time.Second).Unix()
	tv.StandardClaims.Issuer = "lige"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tv)
	token.Header["with-expire"] = true // add custom header

	signed, _ := token.SignedString([]byte("123456")) // sign the token
	fmt.Println("signed token:", signed)              // the generated JWT
	fmt.Println()

	// parse the token and validate it
	tk, err := jwt.Parse(signed, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("123456"), nil
	})
	fmt.Println("err:", err)
	fmt.Println("valid:", tk.Valid)
	fmt.Println("signature:", tk.Signature)
	fmt.Println("raw:", tk.Raw)
	fmt.Println("header:", tk.Header) // with custom header fields
	fmt.Println("valid claim:", tk.Claims.Valid())
	fmt.Println()

	// display payload
	segments := strings.Split(signed, ".")
	headerB64 := segments[0]
	payloadB64 := segments[1]
	sigB64 := segments[2]
	header, _ := base64.RawURLEncoding.DecodeString(headerB64)
	payload, _ := base64.RawURLEncoding.DecodeString(payloadB64)
	sig, _ := base64.RawURLEncoding.DecodeString(sigB64)
	fmt.Println(string(header))          // header json: {"alg":"HS256","encoding":"RawURL","typ":"JWT"}
	fmt.Println(string(payload))         // payload json: {"name":"lgq","Id":1,"loc":"shanghai"}
	fmt.Println(hex.EncodeToString(sig)) // ecdsa signature
	fmt.Println()

	// wait the token tobe expired
	time.Sleep(6 * time.Second)
	tk, err = jwt.Parse(signed, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("123456"), nil
	})
	fmt.Println("err:", err) // err: "token is expired"
	fmt.Println("valid:", tk.Valid)
	fmt.Println("valid claim", tk.Claims.Valid())
}
