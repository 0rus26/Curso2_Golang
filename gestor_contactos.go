package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	c1 := contact{
		Name:  "aa",
		Email: "jgmail.com",
		Phone: "12345666",
	}

	c2 := contact{
		Name:  "bbbb",
		Email: "bbb@gmail.com",
		Phone: "3356662",
	}

	lista_contactosa := []contact{c1, c2}
	lista_contactosb := []contact{}
	saveContactsToFile(lista_contactosa)

	loadContactsFromFile(&lista_contactosb, "contacts.json")

	for _, v := range lista_contactosb {
		fmt.Println("Nombre conacto: ", v.Name)
	}
}

type contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactsToFile(contacts []contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		fmt.Println("Error al crear archivo .json")
		return err
	}

	defer file.Close()

	//encoder, _ := json.MarshalIndent(contacts, "", " ")
	//err = ioutil.WriteFile("contacts.json", encoder, 0644)
	//encoder.
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)

	if err != nil {
		return err
	}
	return nil

}

func loadContactsFromFile(contacts *[]contact, archivo string) error {
	file, err := os.Open(archivo)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil

}
