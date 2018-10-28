package util

import (
	"bytes"
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	"math/big"
	"testing"
	"encoding/hex"
)

func TestDockerKey(t *testing.T) {
	type privKey struct {
		d, x, y string
	}
	keys := []privKey{
		{"R0aX_K9DkBeqL5DW_Jz2cpAHmLFqA0IiaTxUX2CPg50", "3C7kMmWVKoRcT6k-5KsQ8yugVENiSCEzUUsrXBiNQlQ", "GxT4iyMeggh9oxty8ODfD0mXtKmAF94TDDdMwRnFNlY"},
		{"LqcozJ4HPUa_7RYNOlKefcFQct7euoEdXbi-j7qTxS4", "8xTY4hPJcTIves7E8OrSdpYqECXfCMfFT1iUj-9y9Rs", "Oh53uB89rKRh8Q7CqIUJvjZ_Yo_fdXnKktp6lH10OJE"},
		{"h1j6BMfw6Z3Nv8lMkT3KGwqsoJ3pnfWb7HS0MHTXYYc", "CDLHmHWdoYTY4xM6cHQmhFQRxKsPFyAoBfEgHVYls2M", "3v8t2x-a0KPSBwpXsdzPJQB_BAC8bggt6Egi-GLRivY"},
		{"D00we1lvii5JRuD_FbunAsVxJoSurE3eMAyG-p1U_bo", "8qmksN-_VZuRMFdXhzc0kpCyOh3mnyulBFdsq0vMpUE", "uNds3LDn05Y7UOUePfOS9qATKfXsCKUPep-pBn32aE4"},
	}
	for _, key := range keys {
		d_bytes, _ := base64.RawURLEncoding.DecodeString(key.d)
		x_bytes, _ := base64.RawURLEncoding.DecodeString(key.x)
		y_bytes, _ := base64.RawURLEncoding.DecodeString(key.y)
		fmt.Printf("d: %x\nx: %x\ny: %x\n", d_bytes, x_bytes, y_bytes)
		x, y := elliptic.P256().ScalarBaseMult(d_bytes)
		xx, yy := new(big.Int).SetBytes(x_bytes), new(big.Int).SetBytes(y_bytes)
		fmt.Println(elliptic.P256().IsOnCurve(xx, yy))
		fmt.Printf("x: %x\ny: %x\n", x.Bytes(), y.Bytes())
		fmt.Println(bytes.Equal(x.Bytes(), x_bytes))
		fmt.Println(bytes.Equal(y.Bytes(), y_bytes))
		fmt.Println(elliptic.P256().IsOnCurve(x, y))
	}
}

func TestBase64Padding(t *testing.T) {
	strs := []string{
		"R0aX_K9DkBeqL5DW_Jz2cpAHmLFqA0IiaTxUX2CPg50",
		"3C7kMmWVKoRcT6k-5KsQ8yugVENiSCEzUUsrXBiNQlQ",
		"GxT4iyMeggh9oxty8ODfD0mXtKmAF94TDDdMwRnFNlY",
		"LqcozJ4HPUa_7RYNOlKefcFQct7euoEdXbi-j7qTxS4",
		"8xTY4hPJcTIves7E8OrSdpYqECXfCMfFT1iUj-9y9Rs",
		"Oh53uB89rKRh8Q7CqIUJvjZ_Yo_fdXnKktp6lH10OJE",
		"h1j6BMfw6Z3Nv8lMkT3KGwqsoJ3pnfWb7HS0MHTXYYc",
		"CDLHmHWdoYTY4xM6cHQmhFQRxKsPFyAoBfEgHVYls2M",
		"3v8t2x-a0KPSBwpXsdzPJQB_BAC8bggt6Egi-GLRivY",
		"D00we1lvii5JRuD_FbunAsVxJoSurE3eMAyG-p1U_bo",
		"8qmksN-_VZuRMFdXhzc0kpCyOh3mnyulBFdsq0vMpUE",
		"uNds3LDn05Y7UOUePfOS9qATKfXsCKUPep-pBn32aE4",
	}
	for _, str := range strs {
		std, _ := base64.StdEncoding.DecodeString(str)
		stdRaw, _ := base64.RawStdEncoding.DecodeString(str)
		url, _ := base64.URLEncoding.DecodeString(str)
		urlRaw, _ := base64.RawURLEncoding.DecodeString(str)
		fmt.Printf("%x\n%x\n%x\n%x\n\n", std, stdRaw, url, urlRaw)
	}

	bytes, _ := hex.DecodeString("474697fcaf439017aa2f90d6fc9cf672900798b16a034222693c545f608f839d")
	fmt.Println(base64.StdEncoding.EncodeToString(bytes))
	fmt.Println(base64.RawStdEncoding.EncodeToString(bytes))
	fmt.Println(base64.URLEncoding.EncodeToString(bytes))
	fmt.Println(base64.RawURLEncoding.EncodeToString(bytes))
}
