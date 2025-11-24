package main

import "fmt"

func main() {
	// create a map called "modules"
	modules := map[int]string{
		104: "Datenbanken",
		117: "Betriebssysteme",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	// delete one
	delete(modules, 117)
	// add one
	modules[999] = "Einführung in Go"
	// replace one
	modules[104] = "Datenbanken (fortgeschritten)"
	fmt.Println(modules)
}
