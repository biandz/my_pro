package main

import (
	"fmt"
)

func main() {
	s2()
}

func s1() {
	defer func() {
		fmt.Println("123")
	}()
	defer func() {
		fmt.Println("456")
	}()
	defer func() {
		fmt.Println("789")
	}()

	panic("xxx")
}

type People struct {
}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) showB() {
	fmt.Println("showTeacherB")
}

func s2() {
	t := Teacher{}
	t.ShowA()
}
