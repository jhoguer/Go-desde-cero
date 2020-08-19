package main

// Para importar un modulo se hace usando el siguiente comando
// go mod init github.com/jhoguer/Go-desde-cero/tree/master/paquetes

// Para usar paquetes de terceros, debemos usar el comando
// go build, que compila el proyecto y si estamos usando
// alguna dependencia que no este descargada, esta se desacarga
// se agrega la dependencia al archivo go.mod
// ademas se crea el archivo go.sum que lleva el control de las
// dependencias directas e indirectas que usamos o usan los modulos
// de terceros que hemos descargado
// El comando go build tambien crea un archivo .exe
import (
	"fmt"

	"github.com/jhoguer/Go-desde-cero/tree/master/paquetes/greet"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(greet.English())
	fmt.Println(greet.Spanish())
	fmt.Println(greet.Italian())
	fmt.Println("______________________________________")
	fmt.Println(quote.HelloV3())
	fmt.Println(quote.Concurrency())
	fmt.Println("______________________________________")

}

// go list -m -versions rsc.io/quote ****** Para listar las versiones del paquete
// go get rsc.io/sampler ******* para actualizar a la ultima version
// go get rsc.io/sampler@v1.3.1 ******* para actualizar a una version especifica, en este caso la v1.3.1
// go mod tidy **** para limpiar las dependencias que ya no usamos
