package main

import (
	"bufio"
	"client/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main is the entry point of the application that displays the home menu
// and handles user navigation between registration, login, and exit options.
func main() {
	var option int

	for {
		fmt.Println(utils.MENU_HOME_TEXT)
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&option)
		switch option {
		case 1:
			err := handleRegister()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			user, err := handleLogin()
			if err != nil {
				fmt.Println(err)
			}
			viewMainMenu(user)
		case 3:
			fmt.Println("Terminando ejecución...")
			os.Exit(0)
			return
		default:
			fmt.Println("Opción no válida. Ingrese una opción correcta.")
			continue
		}
	}
}

// viewMainMenu displays the main application menu for logged-in users.
// It takes a pointer to the authenticated user and shows the principal menu options.
func viewMainMenu(user *utils.User) {
	var option int

	for {
		fmt.Println(utils.MENU_PRINCIPAL_TEXT)
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&option)
		switch option {
		case 1:
			catalog, err := utils.GetCatalog()
			if err != nil {
				fmt.Println("Error al obtener el catálogo:", err)
				continue
			}
			fmt.Println(utils.CATALOGO_TABLA_HEADER)
			for _, book := range catalog {
				fmt.Printf("| %-2d | %-20s | %-17s | %-11s | %-5d |\n", book.ID, book.Title, book.Category, book.Transaction_type, book.Price)
			}
			fmt.Println(utils.CATALOGO_TABLE_FOOTER)
		case 2:
			// TODO : Carros de compras
		case 3:
			// TODO : Mis prestamos
		case 4:
			// TODO : Mi cuenta
		case 5:
			// TODO : Populares
		case 6:
			fmt.Println("Cerrando sesión...")
			return
		default:
			fmt.Println("Opción no válida. Ingrese una opción correcta.")
			continue
		}
	}
}

func viewShoppingCart(user *utils.User) {
	var cartItemIDs []int

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Ingrese el ID del libro a agregar al carro: ")

		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println("Exiting...")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("ID inválida, intentalo otra vez.")
			continue
		}

		if num >= 0 {
			cartItemIDs = append(cartItemIDs, num)
		} else if num < 0 {
			fmt.Println("ID inválida, intentalo otra vez.")
		}
	}

	var books []utils.Book
	for _, id := range cartItemIDs {
		book, err := utils.GetBookByID(id)
		if err != nil {
			fmt.Println("Error al obtener el libro:", err)
			continue
		}
		books = append(books, book)
	}

	fmt.Println(utils.CARRO_COMPRA_TABLE_HEADER)
	for _, book := range books {
		fmt.Printf("| %-20s | %-17s | %-11d | %-20s |\n", book.Title, book.Transaction_type, book.Price, "N/A")
	}
	fmt.Println(utils.CARRO_COMPRA_TABLE_FOOTER)
}

// handleRegister handles user registration by collecting user information
// and sending it to the server for account creation.
// Returns an error if registration fails, nil on success.
func handleRegister() (error) {
	var user utils.User = utils.User{}

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&user.Name)
	fmt.Print("Ingrese su apellido: ")
	fmt.Scanln(&user.Surname)
	fmt.Print("Ingrese su email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Ingrese su contraseña: ")
	fmt.Scanln(&user.Password)
	
	if !utils.RegisterUser(user) {
		fmt.Println("Error al registrar el usuario. Intente nuevamente.")
		return errors.New("registration failed")
	}
	
	fmt.Println("Usuario registrado exitosamente.")
	return nil
}

// handleLogin handles user authentication by validating credentials
// and populating the user struct with authenticated user data.
// Returns an error if login fails, nil on success.
func handleLogin() (*utils.User, error) {
	var email, contrasena string

	fmt.Print("Ingrese su email: ")
	fmt.Scanln(&email)
	fmt.Print("Ingrese su contraseña: ")
	fmt.Scanln(&contrasena)

	user, err := utils.LoginUser(email, contrasena)
	if err != nil {
		fmt.Println("Error al iniciar sesión.")
		return nil, errors.New("login failed")
	}

	fmt.Println("Inicio de sesión exitoso.")
	return user, nil
}