package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	// TODO: declare a type for Class (consisting of multiple students)
	// TODO: declare a map of modules being attended by multiple classes
	// TODO: output everything using fmt.Println()

	type Student struct {
		FirstName string
		LastName  string
	}

	type Class struct {
		Name     string
		Students []Student
	}

	type Module struct {
		Number  int
		Name    string
		Classes []Class
	}

	// sample students
	classA := Class{
		Name: "A",
		Students: []Student{
			{FirstName: "Anna", LastName: "Muster"},
			{FirstName: "Ben", LastName: "Muster"},
			{FirstName: "Clara", LastName: "Muster"},
		},
	}

	classB := Class{
		Name: "B",
		Students: []Student{
			{FirstName: "Daniel", LastName: "Beispiel"},
			{FirstName: "Eva", LastName: "Beispiel"},
			{FirstName: "Fabian", LastName: "Beispiel"},
		},
	}

	// modules attended by classes
	modules := map[int]Module{
		101: {Number: 101, Name: "Mathematik", Classes: []Class{classA}},
		102: {Number: 102, Name: "Informatik", Classes: []Class{classA, classB}},
		103: {Number: 103, Name: "Physik", Classes: []Class{classB}},
	}

	// print modules, classes and students
	for _, m := range modules {
		fmt.Printf("Module %d: %s\n", m.Number, m.Name)
		for _, c := range m.Classes {
			fmt.Printf("  Class %s:\n", c.Name)
			for _, s := range c.Students {
				fmt.Printf("    - %s %s\n", s.FirstName, s.LastName)
			}
		}
	}
}
