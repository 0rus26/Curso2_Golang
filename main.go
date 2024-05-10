package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	a := 1
	leer := bufio.NewReader(os.Stdin)
	for {
		leer.Reset(os.Stdin)
		switch a {
		case 1:
			fmt.Printf("\nDigite:")
			nombre, _ := leer.ReadString('\n')
			fmt.Printf("\nDigite2:")
			nombre2, _ := leer.ReadString('\n')
			fmt.Printf("\nDigite3:")
			nombre3, _ := leer.ReadString('\n')
			fmt.Printf("\nNOmbre: %v", nombre)
			fmt.Printf("\nNOmbre: %v", nombre2)
			fmt.Printf("\nNOmbre: %v", nombre3)

		default:

		}
		fmt.Println("<--->")
	}
}
