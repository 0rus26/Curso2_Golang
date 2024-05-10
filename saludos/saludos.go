package saludos

import (
	"errors"
	"fmt"
)

// bnlos módulo exportables deben comenzar obligatoriamente con mayusculas e ave amria
func Hola(name string) (string, error) {

	if name == "" {
		err := errors.New("Nombre faltante")
		return "", err
	}

	mensaje := fmt.Sprintf("¡Hola %v, Bienvenido!", name)
	return mensaje, nil
}
