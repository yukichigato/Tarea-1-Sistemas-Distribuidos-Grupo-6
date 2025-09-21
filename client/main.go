package main

import (
	"bufio"
	"client/utils"
	"errors"
	"fmt"
	"os"
	"sort"
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
			catalog, err := utils.GetBooks()
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
			scanner := bufio.NewScanner(os.Stdin)
			var cart []int

			for {
				fmt.Print("Ingrese el ID del libro a agregar al carro: ")
				if !scanner.Scan() {
					break
				}
				text := strings.TrimSpace(scanner.Text())
				if text == "" {
					break
				}
				id, err := strconv.Atoi(text)
				if err != nil {
					fmt.Println("ID inválido, intentalo otra vez.")
					continue
				}
				cart = append(cart, id)
			}

			total := utils.CalculateTotal(cart)
			fmt.Printf("Su carro es de %d libros para un total de %d usm pesos.\n", len(cart), total)

			fmt.Println(utils.CARRO_COMPRA_TABLE_HEADER)
			for _, id := range cart {
				book, err := utils.GetBookByID(id)
				if err != nil {
					fmt.Println("Error al obtener el libro:", err)
					continue
				}
				fmt.Printf("| %-20s | %-17s | %-11d | %-20s |\n", book.Title, book.Transaction_type, book.Price, "-")
			}
			fmt.Println(utils.CARRO_COMPRA_TABLE_FOOTER)

			fmt.Print("Confirmar pedido: ")
			text := strings.TrimSpace(scanner.Text())
			if (text == "") {
				fmt.Println("Pedido realizado con éxito.")
				return
			}
			fmt.Println("Pedido cancelado.")
		case 3:
			loans, err := utils.GetUserLoans(user.ID)
			if err != nil {
				fmt.Println("Error al obtener los préstamos:", err)
				continue
			}
			fmt.Println(utils.PRESTAMOS_TABLE_HEADER)
			for _, loan := range loans {
				fmt.Printf("| %-12d | %-20s | %-15s | %-13s | %-15s | %-9s |\n", loan.ID, "N/A", loan.Start_date, loan.Return_date, "N/A", loan.Status)
			}
			fmt.Println(utils.PRESTAMOS_TABLE_FOOTER)
		case 4:
			fmt.Print(utils.MENU_MI_CUENTA_TEXT)
			var subOption int
			fmt.Print("Seleccione una opción: ")
			fmt.Scanln(&subOption)
			switch subOption {
			case 1:
				fmt.Println("Su saldo actual es de:", user.USM_Pesos, "usm pesos.")
			case 2:
				var amount int
				fmt.Print("Ingrese la cantidad de usm pesos a abonar: ")
				fmt.Scanln(&amount)
				if amount > 0 {
					err := utils.UpdateUserPesos(user.ID, user.USM_Pesos + amount)
					if err != nil {
						fmt.Println("Error al actualizar el saldo:", err)
						continue
					}
					user.USM_Pesos += amount
					fmt.Println("Nuevo saldo es de", user.USM_Pesos, "usm pesos.")
				} else {
					fmt.Println("Cantidad inválida. El abono debe ser mayor a 0.")
				}
			case 3:
				transactions , err := utils.GetUserTransactions(user.ID)
				if err != nil {
					fmt.Println("Error al obtener el historial de transacciones:", err)
					continue
				}

				fmt.Print(utils.HISTORIAL_COMPRAS_ARRIENDOS_TABLE_HEADER)
				for _, transaction := range transactions {
					fmt.Printf("| %-14d | %-9d | %-20s | %-11s | %-20s | %-9d |\n", transaction.ID, transaction.Book_ID, transaction.Title, transaction.Type, transaction.Transaction_date, transaction.Price)
				}
				fmt.Print(utils.HISTORIAL_COMPRAS_ARRIENDOS_TABLE_FOOTER)
			case 4:
				return
			default:
				return
			}
		case 5:
			books, err := utils.GetBooks()
			if err != nil {
				fmt.Println("Error al obtener los libros:", err)
				continue
			}
			
			sort.Slice(books, func(i, j int) bool {
				return books[i].Popularity > books[j].Popularity
			})
			
			if len(books) > 5 {
				books = books[:5]
			}
			
			fmt.Println(utils.POPULARES_TABLE_HEADER)
			for _, book := range books {
				fmt.Printf("| %-8d | %-20s | %-12s | %-11d |\n", book.ID, book.Title, book.Category, book.Popularity)
			}
			fmt.Println(utils.POPULARES_TABLE_FOOTER)
		case 6:
			fmt.Println("Cerrando sesión...")
			return
		default:
			fmt.Println("Opción no válida. Ingrese una opción correcta.")
			continue
		}
	}
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