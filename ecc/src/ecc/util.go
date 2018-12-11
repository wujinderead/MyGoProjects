package ecc

import (
	"math/big"
	"os/exec"
	"fmt"
	"encoding/pem"
	"encoding/asn1"
)

var (
	Zero     = new(big.Int).SetInt64(0)
	ONE      = new(big.Int).SetInt64(1)
	TWO      = new(big.Int).SetInt64(2)
	NEG_ONE  = new(big.Int).SetInt64(-1)
	Infinity = &EcPoint{Zero, Zero}
)

// a/b mod p
func ModFraction(a, b, p *big.Int) *big.Int {
	b_inv := new(big.Int).ModInverse(b, p)
	b_inv.Mul(b_inv, a)
	b_inv.Mod(b_inv, p)
	return b_inv
}

func ModFractionInt64(a, b int64, p *big.Int) *big.Int {
	aa := new(big.Int).SetInt64(a)
	bb := new(big.Int).SetInt64(b)
	return ModFraction(aa, bb, p)
}

func zForAffine(x, y *big.Int) *big.Int {
	z := new(big.Int)
	if x.Sign() != 0 || y.Sign() != 0 {
		z.SetInt64(1)
	}
	return z
}

func getOpensslEcPrivateKey(curve string) (error, []byte, *big.Int, *big.Int) {
	cmd := exec.Command("openssl", "ecparam", "-name", curve, "-genkey")
	data, err := cmd.Output()
	if err != nil {
		fmt.Println("exec openssl error: ", err.Error())
		return err, nil, nil, nil
	}
	_, rest := pem.Decode(data)
	keydata, _ := pem.Decode(rest)
	type ecPrivateKey struct {
		Version       int
		PrivateKey    []byte
		NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
		PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
	}
	priv := new(ecPrivateKey)
	_, err = asn1.Unmarshal(keydata.Bytes, priv)
	if err != nil {
		fmt.Println("unmashal error: ", err.Error())
		return err, nil, nil, nil
	}
	px := new(big.Int).SetBytes(priv.PublicKey.Bytes[1:1+len(priv.PrivateKey)])
	py := new(big.Int).SetBytes(priv.PublicKey.Bytes[1+len(priv.PrivateKey):])
	return nil, priv.PrivateKey, px, py
}
