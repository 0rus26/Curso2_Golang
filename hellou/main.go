package main

import (
	"fmt"
	"log"

	"github.com/0rus26/saludos"
)

func main() {
	log.SetPrefix("saludos: ")
	log.SetFlags(0)

	mensaje, err := saludos.Hola("")
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println(mensaje)

}
