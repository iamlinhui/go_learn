package main

import "fmt"

type Action interface {
	Eat()

	Walk()
}

type Person struct {
	Text string

	Name string
}

func (person *Person) Eat() {
	fmt.Printf("%v学人吃饭\n", person.Name)
}

func (person *Person) Walk() {
	fmt.Println("人走路")
}

type Student struct {
	Person
	Name string
}

func (student *Student) Eat() {
	fmt.Printf("%v学生吃饭\n", student.Name)
}

func (student *Student) Walk() {
	student.Person.Walk()
	fmt.Println("学生走路")
}

func main() {

	var a Action
	a = &Student{Name: "AA"}

	a.Eat()
	a.Walk()

	a = &Person{Name: "BB"}
	a.Eat()
	a.Walk()

	c := &Student{Name: "CC"}
	fmt.Println(c.Person.Name)
	fmt.Println(c.Name)

}
