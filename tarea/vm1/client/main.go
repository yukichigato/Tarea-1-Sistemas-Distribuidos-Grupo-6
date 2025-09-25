package main

import (
	"fmt"
)

// main is the entry point of the application, displays the initial menu and handles user registration, login, and exit
func main() {
	for {
		fmt.Println("Menu")
		fmt.Println()
		fmt.Println("1. Registrarse")
		fmt.Println("2. Iniciar Sesión")
		fmt.Println("3. Salir")
		fmt.Println()

		choice := readInput("Seleccione una opción: ")
		fmt.Println()

		switch choice {
		case "1":
			registerMenu()
		case "2":
			loginMenu(&currentUser)
		case "3":
			return
		}
	}
}