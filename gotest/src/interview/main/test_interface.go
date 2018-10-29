package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}
type Teacher struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	fmt.Println(talk)
	return
}

func (tea Teacher) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	fmt.Println(talk)
	return
}

func invoke(peo People) {
	peo.Speak("aaa")
}

func main() {
	// '*Student' and 'Teacher' implement 'People' interface,
	// so these can compile:
	var peo1 People = Teacher{}
	var peo2 People = &Student{}

	// '*Teacher' can also compile
	var peo3 People = &Teacher{}

	// 'Student' did not implement 'People' interface,
	// so this cannot compile: var peo People = Student{}

	var stu Student = Student{}
	var tea Teacher = Teacher{}
	var stuptr *Student = &Student{}
	var teaptr *Teacher = &Teacher{}
	stu.Speak("bitch")
	tea.Speak("bitch")
	stuptr.Speak("bitch")
	teaptr.Speak("bitch")

	// these can indeed compile
	invoke(peo1)
	invoke(peo2)
	invoke(peo3)

	// this connot compile: invoke(stu)

	// these can apply to People, so these can compile
	invoke(stuptr)
	invoke(tea)
	invoke(teaptr)

	// Conclusion, '*Student' and 'Teacher' implement 'People' interface,
	// then, '*Student', 'Teacher', and '*Teacher' can be applied to 'People'.
}
