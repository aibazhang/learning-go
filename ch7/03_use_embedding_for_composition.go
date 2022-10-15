package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	ID      string
	Reports []Employee
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "11111",
		},
		Reports: []Employee{},
		ID:      "22222",
	}
	fmt.Println(m.ID)
	fmt.Println(m.Employee.ID)
	fmt.Println(m.Description())

	// var eFail Employee = m
	// var eOK Employee = m.Employee
}
