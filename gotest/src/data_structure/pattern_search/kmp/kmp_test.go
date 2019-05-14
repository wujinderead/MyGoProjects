package kmp

import (
	"testing"
	"fmt"
)

func TestKmp(t *testing.T) {
	fmt.Println(search("ABABDABACDABABCABAB", "ABABCABAB"))
	fmt.Println(search("破釜破釜舟破釜破沉舟破釜破釜沉破釜破釜", "破釜破釜沉破釜破釜"))
	fmt.Println(search("AAAAAAAAAAAAAAAAAB", "AAAAB"))
	fmt.Println(search("ABABABCABABABCABABABC", "ABABAC"))
}