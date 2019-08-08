package stdlib

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
)

func TestJson(t *testing.T) {
	bi, _ := new(big.Int).SetString("7fffffffffffffffffffffffed", 16)
	biJson, err := json.Marshal(bi) // marshal to decimal string
	if err != nil {
		fmt.Println("big.Int marshal err: ", err.Error())
	}
	fmt.Println(string(biJson))

	biLoad := new(big.Int)
	err = json.Unmarshal(biJson, biLoad)
	if err != nil {
		fmt.Println("big.Int unmarshal err: ", err.Error())
	}
	fmt.Println(biLoad)

	key, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	ecKeyJson, err := json.Marshal(key) // marshal to json object
	if err != nil {
		fmt.Println("ecdsa key json err: ", err.Error())
	}
	fmt.Println(string(ecKeyJson))

	ecKeyLoad := new(ecdsa.PrivateKey)
	err = json.Unmarshal(ecKeyJson, ecKeyLoad)
	if err != nil {
		fmt.Println("ecdsa key unmarshal err: ", err.Error())  // PrivateKey.Curve can't be unmarshalled
	}
	fmt.Println(ecKeyLoad.D.Cmp(key.D))
	fmt.Println(ecKeyLoad.X.Cmp(key.X))
	fmt.Println(ecKeyLoad.Y.Cmp(key.Y))
}