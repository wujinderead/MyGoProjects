package ecc

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

var names = []string{
	SN_secp112r1,
	SN_secp112r2,
	SN_secp128r1,
	SN_secp128r2,
	SN_secp160k1,
	SN_secp160r1,
	SN_secp160r2,
	SN_secp192k1,
	SN_secp224k1,
	SN_secp224r1,
	SN_secp256k1,
	SN_secp384r1,
	SN_secp521r1,
	SN_X9_62_prime192v1,
	SN_X9_62_prime192v2,
	SN_X9_62_prime192v3,
	SN_X9_62_prime239v1,
	SN_X9_62_prime239v2,
	SN_X9_62_prime239v3,
	SN_X9_62_prime256v1,
	SN_sect113r1,
	SN_sect113r2,
	SN_sect131r1,
	SN_sect131r2,
	SN_sect163k1,
	SN_sect163r1,
	SN_sect163r2,
	SN_sect193r1,
	SN_sect193r2,
	SN_sect233k1,
	SN_sect233r1,
	SN_sect239k1,
	SN_sect283k1,
	SN_sect283r1,
	SN_sect409k1,
	SN_sect409r1,
	SN_sect571k1,
	SN_sect571r1,
	SN_X9_62_c2pnb163v1,
	SN_X9_62_c2pnb163v2,
	SN_X9_62_c2pnb163v3,
	SN_X9_62_c2pnb176v1,
	SN_X9_62_c2tnb191v1,
	SN_X9_62_c2tnb191v2,
	SN_X9_62_c2tnb191v3,
	SN_X9_62_c2pnb208w1,
	SN_X9_62_c2tnb239v1,
	SN_X9_62_c2tnb239v2,
	SN_X9_62_c2tnb239v3,
	SN_X9_62_c2pnb272w1,
	SN_X9_62_c2pnb304w1,
	SN_X9_62_c2tnb359v1,
	SN_X9_62_c2pnb368w1,
	SN_X9_62_c2tnb431r1,
	SN_wap_wsg_idm_ecid_wtls1,
	SN_wap_wsg_idm_ecid_wtls3,
	SN_wap_wsg_idm_ecid_wtls4,
	SN_wap_wsg_idm_ecid_wtls5,
	SN_wap_wsg_idm_ecid_wtls6,
	SN_wap_wsg_idm_ecid_wtls7,
	SN_wap_wsg_idm_ecid_wtls8,
	SN_wap_wsg_idm_ecid_wtls9,
	SN_wap_wsg_idm_ecid_wtls10,
	SN_wap_wsg_idm_ecid_wtls11,
	SN_wap_wsg_idm_ecid_wtls12,
	SN_brainpoolP160r1,
	SN_brainpoolP160t1,
	SN_brainpoolP192r1,
	SN_brainpoolP192t1,
	SN_brainpoolP224r1,
	SN_brainpoolP224t1,
	SN_brainpoolP256r1,
	SN_brainpoolP256t1,
	SN_brainpoolP320r1,
	SN_brainpoolP320t1,
	SN_brainpoolP384r1,
	SN_brainpoolP384t1,
	SN_brainpoolP512r1,
	SN_brainpoolP512t1,
}

func TestParams(t *testing.T) {
	initEc.Do(initEcCurves)
	for _, name := range names {
		spec := EcCurveSpecs[name]
		fmt.Println(strings.Contains(spec.Desc, "prime"), strings.Contains(spec.Desc, "binary"), name, spec.Desc)
	}
}

func TestParams_Prime(t *testing.T) {
	initEc.Do(initEcCurves)
	for _, name := range names {
		spec := EcCurveSpecs[name]
		if spec.Curve.head.fieldType != NID_X9_62_prime_field {
			continue // skip non prime field curve
		}
		curve := spec.Curve
		three := new(big.Int).SetInt64(3)
		featured := false
		if curve.A.Cmp(ZERO) == 0 {
			fmt.Print("a=0, ")
			featured = true
		}
		if curve.B.IsInt64() {
			fmt.Printf("b=%d, ", curve.B.Int64())
			featured = true
		}
		if new(big.Int).Add(curve.A, three).Cmp(curve.P) == 0 {
			fmt.Printf("a=-3, b=%s, ", hex.EncodeToString(curve.B.Bytes()))
			featured = true
		}
		if testPrimeBits(curve.P) {
			fmt.Printf("p=%s, ", hex.EncodeToString(curve.P.Bytes()))
			featured = true
		}
		if !featured {
			fmt.Println("other, ", name)
		} else {
			fmt.Println(name)
		}
	}
}

func TestParams_Binary(t *testing.T) {
	initEc.Do(initEcCurves)
	for _, name := range names {
		spec := EcCurveSpecs[name]
		if spec.Curve.head.fieldType != NID_X9_62_characteristic_two_field {
			continue // skip non binary field curve
		}
		curve := spec.Curve
		featured := false
		if curve.A.IsInt64() {
			fmt.Printf("a=%d, ", curve.A.Int64())
			featured = true
		}
		if curve.B.IsInt64() {
			fmt.Printf("b=%d, ", curve.B.Int64())
			featured = true
		}
		if !featured {
			fmt.Println("other, ", name, hex.EncodeToString(curve.A.Bytes()), hex.EncodeToString(curve.B.Bytes()))
		} else {
			fmt.Println(name)
		}
	}
}

func TestParams_BinaryFieldP(t *testing.T) {
	initEc.Do(initEcCurves)
	upper := []rune{8304, 185, 178, 179, 8308, 8309, 8310, 8311, 8312, 8313}
	for _, name := range names {
		spec := EcCurveSpecs[name]
		if !strings.Contains(spec.Desc, "binary") {
			continue // skip non binary field curve
		}
		curve := spec.Curve
		bl := curve.P.BitLen()
		fmt.Println(name, ":")
		for i := 0; i < bl; i++ {
			if curve.P.Bit(bl-i-1) == 1 {
				pos := bl - i - 1
				if pos != 0 {
					str := strconv.Itoa(pos)
					fmt.Print("x")
					for i := 0; i < len(str); i++ {
						fmt.Print(string(upper[str[i]-'0']))
					}
					fmt.Print(" + ")
				} else {
					fmt.Println(1)
				}
			}
		}
		fmt.Println()
	}
}

func testPrimeBits(p *big.Int) bool {
	count := 0
	for i := 0; i < p.BitLen(); i++ {
		if p.Bit(i) == 1 {
			count++
		}
	}
	if float32(count)/float32(p.BitLen()) > 0.8 {
		return true
	} else {
		return false
	}
}
