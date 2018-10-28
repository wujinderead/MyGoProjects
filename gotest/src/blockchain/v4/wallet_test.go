package v4

import (
	"crypto/sha256"
	"testing"
)

func TestGetAddress(t *testing.T) {
	w := NewWallet()
	t.Logf("private key: %x", w.PrivateKey.D)
	t.Logf("public  key: %x", w.PublicKey)
	t.Logf("public keyX: %x", w.PrivateKey.X)
	t.Logf("public keyY: %x", w.PrivateKey.Y)
	t.Logf("length of D, X, Y: %d", len(w.PrivateKey.D.Bytes()))
	pubKeyHash := HashPubKey(w.PublicKey)
	t.Logf("hash: %x, len: %d.", pubKeyHash, len(pubKeyHash))

	versionedPayload := append([]byte{version}, pubKeyHash...)
	t.Logf("payload: %x, len: %d", versionedPayload, len(versionedPayload))
	checksum := checksum(versionedPayload)
	firstSHA := sha256.Sum256(versionedPayload)
	secondSHA := sha256.Sum256(firstSHA[:])
	t.Logf("checkSHA: %x, len: %d", secondSHA, len(secondSHA))
	t.Logf("checksum: %x, len: %d", checksum, len(checksum))
	fullPayload := append(versionedPayload, checksum...)
	t.Logf("fullPayload: %x, len: %d", fullPayload, len(fullPayload))
	address := Base58Encode(fullPayload)
	t.Logf("address: %s, len: %d", address, len(address))
}
