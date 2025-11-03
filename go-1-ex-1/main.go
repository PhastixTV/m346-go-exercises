package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName = "Philipp"
	var lastName = "Sidler"
	var dayOfBirth = 2
	var monthOfBirth = 06
	var yearOfBirth = 2008
	var numberOfSiblings = 4
	var heightInMeters = 1.72
	var zodiacSign = '♊'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
