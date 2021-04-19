package main

import "fmt"

type Action interface {
	Eat()

	Walk()
}

type Person struct {
}

func (person *Person) Eat() {
	fmt.Println("人吃饭")
}

func (person *Person) Walk() {
	fmt.Println("人走路")
}

type Student struct {
	Person
}

func (student *Student) Eat() {
	fmt.Println("学生吃饭")
}

func (student *Student) Walk() {
	student.Person.Walk()
	fmt.Println("学生走路")
}

func main() {

	var a Action
	a = &Student{}

	a.Eat()
	a.Walk()

	a = &Person{}
	a.Eat()
	a.Walk()

}
