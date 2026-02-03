package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// estructura de contactos

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guardar contactos en un archivo json

func SaveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

func LoadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
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

func main() {
	// slice de contactos
	var contacts []Contact
	// cargar contactos existentes desde el archivo
	err := LoadContactsFromFile(&contacts)
	if err != nil {
		fmt.Print("Error al cargar los contactos:", err)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar opciones al usuario
		fmt.Println("==== Gestor de cotactos ====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opcion")

		var opcion int
		_, err = fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("error al leer la opcion", err)
			return
		}

		switch opcion {
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			// Agregar un contacto al slice
			contacts = append(contacts, c)

			// guadarlo en un archivo json

			if err := SaveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto")
			}
		case 2:
			// Mostrar todos los contactos
			fmt.Println("============================================")
			for index, contact := range contacts {
				fmt.Printf("%d.Nombre: %s Email: %s Telefono: %s",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("============================================")
		case 3:
			//salir dle progranma
			fmt.Println("Has salido del programa")
			return
		default:
			fmt.Println("Opcion invlida")
		}
	}
}
