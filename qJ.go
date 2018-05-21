package main

import (
	"fmt"
)

type (
	subject struct {
		a int
		b int
		c int
	}
)

func (o subject) getGrade(marks int) string {
	switch {
	case marks >= o.a:
		return "A"
	case marks < o.a && marks >= o.b:
		return "B"
	case marks < o.b && marks >= o.c:
		return "C"
	default:
		return "D"
	}
}

func main() {
	var marks int = 70
	subjectChineseType := subject{80, 70, 50}
	grade := subjectChineseType.getGrade(marks)
	fmt.Println("grade is", grade)

	var subjectEnglishType subject = subject{90, 80, 70}
	grade = subjectEnglishType.getGrade(marks)
	fmt.Println("grade is", grade)

}
