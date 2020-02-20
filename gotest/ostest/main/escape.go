package main

/*
# -m print optimization decisions
# -N disable optimization
# -l disable inlining
go build -gcflags="-N -m -l" ostest/main/escape.go

*/
func main() {
	testEscape1()
	testEscape2()
	testEscape3()
}

func testEscape1() {
	_ = newUser()
}

func testEscape2() {
	a := &user{}
	b := "lgq"
	a.name = &b // indirect assign: move to heap: b

	c := "test"
	aa := &user{ // direct assign: c no escape
		name: &c,
	}
	_ = aa
}

type user struct {
	name *string
}

func newUser() *user {
	a := user{} // moved to heap: a
	return &a
}

func testEscape3() {
	clo := 10      // clo shared between goroutines, so must escape to heap: moved to heap: clo
	go func() {    // the function also escape: func literal escapes to heap
		clo += 10
	}()
	_ = clo
}