package main

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	// return nil, false
	// cannot return nil as string, the default value of string is ""
	var a string
	return a, false
}
func main() {
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	v, err := GetValue(intmap, 4)
	fmt.Println(v, err)
	fmt.Println(v == "")
}
