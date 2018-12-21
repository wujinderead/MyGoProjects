package reflector

import (
	"crypto/elliptic"
	"fmt"
	"strconv"
)

type Aircraft struct {
	no   int8   `json:"no"`
	Name string `json:"name"`
}
type Bird struct {
	name string
	elliptic.CurveParams
}
type Fly interface {
	fly(from int, to uint) bool
	Fly(param ...int) string
}

func (aircraft *Aircraft) fly(from int, to uint) bool {
	fmt.Println(aircraft.Name, aircraft.no)
	return true
}

func (aircraft *Aircraft) Fly(param ...int) string {
	return aircraft.Name + " " + strconv.Itoa(len(param))
}

func (bird Bird) fly(from int, to uint) bool {
	fmt.Println(bird)
	return true
}

func (bird Bird) Fly(param ...int) string {
	return "birdy"
}

func (aircraft *Aircraft) String() string {
	return aircraft.Name
}
