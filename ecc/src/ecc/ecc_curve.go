package ecc

import (
	"fmt"
	"math/big"
)

type EcCurve struct {
	head                       *EcCurveHead
	Seed, P, A, B, X, Y, Order *big.Int
}

type EcPoint struct {
	X, Y *big.Int
}

// for instantiate an EcPoint
func NewPoint() *EcPoint {
	return &EcPoint{new(big.Int).SetInt64(0), new(big.Int).SetInt64(0)}
}

func (point *EcPoint) Equals(p *EcPoint) bool {
	return point.X.Cmp(p.X) == 0 && point.Y.Cmp(p.Y) == 0
}

func (point *EcPoint) Copy() *EcPoint {
	return &EcPoint{new(big.Int).Set(point.X), new(big.Int).Set(point.Y)}
}

func (point *EcPoint) ToStr() string {
	return fmt.Sprintf("[%x, %x]", point.X.Bytes(), point.Y.Bytes())
}

func parseEcCurve(head *EcCurveHead, data []byte) *EcCurve {
	ret := &EcCurve{}
	ret.head = head
	if head.seedLen > 0 {
		ret.Seed = new(big.Int).SetBytes(data[0:head.seedLen])
	}
	ret.P = new(big.Int).SetBytes(data[head.seedLen : head.seedLen+head.paramLen])
	ret.A = new(big.Int).SetBytes(data[head.seedLen+head.paramLen : head.seedLen+head.paramLen*2])
	ret.B = new(big.Int).SetBytes(data[head.seedLen+head.paramLen*2 : head.seedLen+head.paramLen*3])
	ret.X = new(big.Int).SetBytes(data[head.seedLen+head.paramLen*3 : head.seedLen+head.paramLen*4])
	ret.Y = new(big.Int).SetBytes(data[head.seedLen+head.paramLen*4 : head.seedLen+head.paramLen*5])
	ret.Order = new(big.Int).SetBytes(data[head.seedLen+head.paramLen*5 : head.seedLen+head.paramLen*6])
	return ret
}
