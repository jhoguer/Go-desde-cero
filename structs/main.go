package main

import "fmt"

func main() {
	type course struct {
		Name      string
		Professor string
		Country   string
	}

	db := course{
		Name:      "Bases de datos",
		Professor: "Alexys",
		Country:   "Colombia",
	}

	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", db.Name)

	// Para declarar una instacia por medio de su estructura literal
	// se deben declarar en el mismo orden en que se delcaro
	// Name - Professor - Country
	git := course{"Git", "Beto", "Bolivia"}
	fmt.Printf("%+v\n", git)
	fmt.Printf("%+v\n", git.Country)

	// Para declarar una instancia de un nuevo curso con solo el nombre del professor por ejemplo
	css := course{Professor: "Alvaro"}
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v\n", css.Professor)

	fmt.Println("================================================")
	// Declaracion de puntero p que apunta a la struct db
	p := &db
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", &db)
	fmt.Printf("%+v\n", &p)
	fmt.Printf("%+v\n", &p)
	fmt.Printf("%+v\n", p.Name)

	fmt.Println("================================================")
	p.Name = "Beto"
	fmt.Printf("%+v\n", p.Name)
}
