package golangx

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

// bcrypt: generate hashed password and hide the real password
func TestBrcypt(t *testing.T) {
	passes := [][]byte{[]byte("test123"), []byte("lgq123456"), []byte("mypassword!@#$%^&*()")}
	for _, pass := range passes {
		for _, cost := range []int{4, 7, 10} {
			for i := 0; i < 3; i++ {
				// the hashed password is always 480 bits (60 bytes)
				// we can generate different hashes for the same password
				hpass, _ := bcrypt.GenerateFromPassword(pass, cost)
				// we can detect cost for hashed password
				bc, _ := bcrypt.Cost(hpass)
				fmt.Println(string(pass), cost, bc, hex.EncodeToString(hpass))
			}
			fmt.Println()
		}
	}
}

func TestCompare(t *testing.T) {
	realpass := []byte("test123")
	fakepass := []byte("mypassword!@#$%^&*()")
	hpasses := make([][]byte, 0)
	for _, cost := range []int{4, 7, 10} {
		for i := 0; i < 3; i++ {
			// generate hashed passwords multiple times and with multiple costs
			hpass, _ := bcrypt.GenerateFromPassword(realpass, cost)
			fmt.Println(hex.EncodeToString(hpass))
			hpasses = append(hpasses, hpass)
		}
	}
	// all the hashed passwords match the real password
	for _, hpass := range hpasses {
		err := bcrypt.CompareHashAndPassword(hpass, realpass)
		fmt.Println(err)
	}
	// all the hashed passwords don't match the fake password
	for _, hpass := range hpasses {
		// return error: "crypto/bcrypt: hashedPassword is not the hash of the given password"
		err := bcrypt.CompareHashAndPassword(hpass, fakepass)
		fmt.Println(err)
	}
}

// cost is the relatively time cost to generate the hashed password, range [4, 31].
// the time consumption seems increasing exponentially. so do not use large cost.
func TestCost(tt *testing.T) {
	pass := []byte("test123")
	t := time.Now()
	for cost := 4; cost <= 15; cost++ {
		_, _ = bcrypt.GenerateFromPassword(pass, cost)
		// on this computer (8 cores i7-6700, 32GB memory), time cost is:
		// cost4 1.040859ms, cost5 2.125058ms, cost6 4.023802ms, cost7 7.63987ms, cost8 15.305128ms,
		// cost9 29.774477ms, cost10 60.050374ms, cost11 113.922648ms, cost12 225.6935ms, ...
		fmt.Print("cost", cost, " ", time.Now().Sub(t), ", ")
		t = time.Now()
	}
}
