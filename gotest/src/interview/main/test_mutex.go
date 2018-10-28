package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	map1 := UserAges{ages: make(map[string]int)}
	go map1.Add("lgq", 21)
	go map1.Add("zw", 22)
	go map1.Add("ljc", 23)
	go map1.Add("lsy", 24)
	time.Sleep(1 * time.Second)
	fmt.Println(map1.Get("lgq"))
	fmt.Println(map1.Get("zw"))
}
