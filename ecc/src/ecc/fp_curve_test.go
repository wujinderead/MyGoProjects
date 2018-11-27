package ecc

import (
	"crypto/elliptic"
	"crypto/rand"
	"io"
	"math/big"
	"testing"
)

func TestConst(t *testing.T) {
	if Zero.Cmp(new(big.Int).SetInt64(0)) == 0 {
		t.Log("good.")
	}
	infi := &EcPoint{new(big.Int).SetInt64(0), new(big.Int).SetInt64(0)}
	if Infinity.Equals(infi) {
		t.Log("good.")
	}
	if Infinity.Equals(NewPoint()) {
		t.Log("good.")
	}
	infiCopy := Infinity.Copy()
	np := NewPoint()
	p := &EcPoint{new(big.Int).SetInt64(0), new(big.Int).SetInt64(2)}
	pCopy := p.Copy()
	t.Logf("%p, %p, %p, %v, %v", infi, infiCopy, np, infi.Equals(np), np.Equals(infiCopy))
	t.Logf("%p, %p, %p, %p, %p, %p", infi.X, infi.Y, infiCopy.X, infiCopy.Y, np.X, np.Y)
	t.Logf("%p, %p, %v", p, pCopy, p.Equals(pCopy))
	t.Logf("%p, %p, %p, %p", p.X, p.Y, pCopy.X, pCopy.Y)
}

func TestEcCurve_Add(t *testing.T) {
	curve, _ := GetFpCurve("secp224r1")
	x1, _ := new(big.Int).SetString("19277929113566293071110308034699488026831934219452440156649784352033", 10)
	y1, _ := new(big.Int).SetString("19926808758034470970197974370888749184205991990603949537637343198772", 10)
	t.Log(curve.X.Cmp(x1) == 0, curve.Y.Cmp(y1) == 0)

	x2, _ := new(big.Int).SetString("11838696407187388799350957250141035264678915751356546206913969278886", 10)
	y2, _ := new(big.Int).SetString("2966624012289393637077209076615926844583158638456025172915528198331", 10)
	x3, _ := new(big.Int).SetString("23495795443371455911734272815198443231796705177085412225858576936196", 10)
	y3, _ := new(big.Int).SetString("17267899494408073472134592504239670969838724875111952463975956982053", 10)
	x5, _ := new(big.Int).SetString("5241180935788447299415492279837860720508896463754443826289355788714", 10)
	y5, _ := new(big.Int).SetString("4202927080198900989467433352160374810063042822299340499566031669147", 10)
	p1 := &EcPoint{x1, y1}
	p2 := &EcPoint{x2, y2}
	p3 := &EcPoint{x3, y3}
	p5 := &EcPoint{x5, y5}
	x2_re := curve.Add(p1, p1)
	t.Log(x2_re.Equals(&EcPoint{x2, y2}))
	t.Log(x2_re.ToStr())
	t.Log(p2.ToStr())
	x3_re := curve.Add(p2, &EcPoint{x1, y1})
	t.Log(x3_re.Equals(&EcPoint{x3, y3}))
	t.Log(x3_re.ToStr())
	t.Log(p3.ToStr())
	x5_re := curve.Add(p3, p2)
	t.Log(x5_re.Equals(&EcPoint{x5, y5}))
	t.Log(x5_re.ToStr())
	t.Log(p5.ToStr())
	xa, _ := new(big.Int).SetString("13767876268297836769844486473490979114236434690224307562045282460159", 10)
	ya, _ := new(big.Int).SetString("22118781371587511354922822567431197267775915011071080155979634070273", 10)
	xb, _ := new(big.Int).SetString("20099534334183616009144096100489037826469183056754563418223687238940", 10)
	yb, _ := new(big.Int).SetString("21163395713719829884372140998109152372217705929433802923814725259789", 10)
	xc, _ := new(big.Int).SetString("20265066168786098004447144135144701335001640345992457384057083230065", 10)
	yc, _ := new(big.Int).SetString("8984557987810807037190217998256774306118199921011006968275655413387", 10)

	pa := &EcPoint{xa, ya}
	pb := &EcPoint{xb, yb}
	pc := &EcPoint{xc, yc}
	k1 := []byte{0xd3, 0x64, 0x75, 0xF4}
	k2 := []byte{0xd3, 0x55, 0x0E, 0x64, 0x75, 0xF4}
	k3 := []byte{0xd3, 0x55, 0x0E, 0x64, 0x75, 0xF4, 0xd3, 0x55, 0x0E, 0x64, 0x75, 0xF4}
	t.Log(curve.ScalaMult(p1, k1).ToStr())
	t.Log(pa.ToStr())
	t.Log(curve.ScalaMult(p1, k2).ToStr())
	t.Log(pb.ToStr())
	t.Log(curve.ScalaMult(p1, k3).ToStr())
	t.Log(pc.ToStr())
}

func TestEcCurve_ScalaMult(t *testing.T) {
	curve, _ := GetFpCurve("secp521r1")
	cp := elliptic.P521()
	t.Logf("%x %x %x", cp.Params().Gx.Bytes(), cp.Params().Gy.Bytes(), cp.Params().P.Bytes())
	t.Logf("%x %x %x", curve.X.Bytes(), curve.Y.Bytes(), curve.P.Bytes())
	byteLen := (curve.P.BitLen() + 7) >> 3
	priv := make([]byte, byteLen)

	io.ReadFull(rand.Reader, priv)
	px, py := cp.ScalarBaseMult(priv)
	p := &EcPoint{px, py}
	t.Log(p.ToStr())
	point := curve.ScalaMultBase(priv)
	t.Log(point.ToStr(), point.Equals(p))
}
