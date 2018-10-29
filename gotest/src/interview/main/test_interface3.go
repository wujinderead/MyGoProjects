package main

import "fmt"

type User struct {
}
type MyUser1 User   // type definition, a whole new type
type MyUser2 = User // type alias, the are the same
func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

func main() {
	var i1 MyUser1
	var i2 MyUser2
	var usr User
	i1.m1()
	i2.m2()
	//i1.m2()
	//i2.m1()
	usr.m2()
	//usr.m1()
	/*
		MyUser2 is alias of User, so they both have m2.
		MyUser1 only has m1.
	*/
}
