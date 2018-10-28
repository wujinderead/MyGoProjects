package main

import (
	"ecc"
	"fmt"
)

func main() {
	spec := ecc.GetEcCurveSpec(ecc.SN_brainpoolP160r1)
	fmt.Println(spec.Name, spec.Desc)
	curve := ecc.GetEcCurve(ecc.SN_brainpoolP160r1)
	fmt.Printf("%x %x", curve.X, curve.Y)
	//fmt.Println(curve.IsOnCurve(curve.X, curve.Y))
}
