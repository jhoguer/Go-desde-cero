package main

import "fmt"

func main() {
	// bool, string, numeric
	var a bool = true
	var b string = "EDteam"
	var c uint16 = 300

	// byte es un alias para uint8
	var d byte = 255

	// rune es un alias para int32
	var e rune = -342323535
	var f int32 = 353547574

	// Si pasamos un caracter se asigna su representacion unicode
	// funciona en cualquier tipo numerico
	var g int64 = 'a'

	// tipo flotante float64
	var h float64 = 23.456

	// Operaciones
	var i uint8 = 100
	var j uint16 = 23000

	// operacion no valida por operar tipos diferentes
	// l := i + j
	// fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", l, l)

	// se debe convertir las variables al mismo tipo de dato
	// en este caso covertimos i a uint16

	l := uint16(i) + j

	// para declarar variables que planeamos usar despues
	// usamos Blank Identifier(underscore) _
	_ = 234
	var _ string = "test"

	// todas las variables que se declaran y no se inicializan
	// toman un valor por defecto valor cero
	// para string en valor cero es ""
	var m string

	// para numeros en valor cero es el numero 0
	var n uint
	var o float32

	// para booleanos en valor cero es false
	var p bool

	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", a, a)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", b, b)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", c, c)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", d, d)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", e, e)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", f, f)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", g, g)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", h, h)
	fmt.Printf("-------------------------------------\n")
	fmt.Printf("Cast i de uint8 a uint16-------------\n")
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", l, l)
	fmt.Printf("uint8--------------------------------\n")
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", i, i)
	fmt.Printf("uint16-------------------------------\n")
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", j, j)
	fmt.Printf("-------------------------------------\n")
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %q\n", m, m)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", n, n)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", o, o)
	fmt.Printf("Tipo:\t\t %T \nValor:\t\t %v\n", p, p)
}
