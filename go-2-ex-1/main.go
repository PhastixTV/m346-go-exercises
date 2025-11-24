package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	// embed full name and birth date information
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName:         FullName{FirstName: "Max", LastName: "Mustermann"},
		BirthDate:        BirthDate{Day: 1, Month: 1, Year: 2000},
		NumberOfSiblings: 0,
		ZodiacSign:       '♒',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
