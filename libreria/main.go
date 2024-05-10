package main

import (
	"fmt"
	"libreria/libro"
)

func main() {

	//libro1 := libro.Libro{
	//		Name:    "Quijote de la mancha",
	//Autor:   "Miguel de Cervantes",
	//Paginas: 500,
	//}

	libro2 := libro.NuevoLibro("Harry Potter", "No sé", 100)
	//libro1.Imprimir_info()
	libro2.SetName("Harry Potter 2")
	libro2.GetName()
	fmt.Println(libro2.GetName())

	libro2.Imprimir_info()

	lbTxt1 := libro.NuevoLibroTexto("Harry Potter texto", "No sé", 100, "Editorial 1", "Nivel 1")
	lbTxt1.Imprimir_info()
	fmt.Println(lbTxt1.GetName())
	fmt.Println("END")
}
