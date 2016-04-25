package main

import (
	"fmt"
	"time"
	"zhangmingkai4315/calc"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}
func (p *Person) ChangeName(fn string) {
	p.FirstName = fn
}

type Admin struct {
	Person
	Roles []string
}

func (a Admin) ChangeName(n string) {
	a.Person.ChangeName(n)
	fmt.Println("Admin Change the name")
}

type Member struct {
	Person
	Skill []string
}

func (m Member) ChangeName(n string) {
	m.Person.ChangeName(n)
	fmt.Println("Member Change the name")
}

// Interface

type User interface {
	PrintName()
	ChangeName(string)
}

type Team struct {
	Name  string
	users []User
}

func (t *Team) PrintTeam() {
	fmt.Println(t.Name)
	for _, v := range t.users {
		v.PrintName()
	}
}

func main() {
	var x, y int = 10, 5
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))
	p := &Person{
		FirstName: "Mike",
		LastName:  "Z",
		Dob:       time.Date(2016, time.February, 1, 0, 0, 0, 0, time.UTC),
		Email:     "Z@gmail.com",
		Location:  "China",
	}

	p.PrintName()
	p.ChangeName("Alice")
	p.PrintName()

	mike := Admin{
		Person{
			"Mike",
			"Z",
			time.Date(2016, time.February, 1, 0, 0, 0, 0, time.UTC),
			"z@gmail.com",
			"Beijing",
		},
		[]string{"Administrator", "Manager"},
	}
	alice := Member{
		Person{
			"Alice",
			"J",
			time.Date(2011, time.February, 1, 0, 0, 0, 0, time.UTC),
			"j@gmail.com",
			"guest",
		},
		[]string{"go", "docker"},
	}

	users := []User{mike, alice}
	for _, v := range users {
		v.PrintName()
	}

	mike.PrintName()
	alice.PrintName()
	alice.ChangeName("AAlice")
	alice.PrintName()

	t := &Team{
		"Go Project Team",
		[]User{alice, mike},
	}

	t.PrintTeam()

}
