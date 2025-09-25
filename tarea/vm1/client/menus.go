package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Global variable to hold current user session (moved from main.go)
var currentUser User

// registerMenu handles user registration by collecting user information and creating a new account
func registerMenu() {
	nombre := readInput("Ingrese su nombre: ")
	apellido := readInput("Ingrese su apellido: ")
	email := readInput("Ingrese su email: ")
	password := readInput("Ingrese su contraseña: ")

	err := CreateUser(User{Name: nombre, Surname: apellido, Email: email, Password: password, USM_Pesos: 100})
	if err != nil {
		fmt.Println("Error al crear el usuario:", err)
	} else {
		fmt.Println()
		fmt.Println("Usuario creado con éxito")
		fmt.Println()
	}
}

// loginMenu handles user authentication and sets the current user session
func loginMenu(currentUser *User) {
	email := readInput("Ingrese su email: ")
	fmt.Println()
	password := readInput("Ingrese su contraseña: ")
	fmt.Println()
	user, err := FetchUser(email, password)
	if err != nil {
		fmt.Println("Error al iniciar sesión:", err)
	} else {
		fmt.Println("Sesión iniciada correctamente")
		fmt.Println()
		*currentUser = user
		mainMenu()
	}
}

// mainMenu displays the main application menu after successful login with options for catalog, cart, loans, account, and popular books
func mainMenu() {
	for {
		fmt.Println("Menu")
		fmt.Println()
		fmt.Println("1. Ver catálogo")
		fmt.Println("2. Carro de compras")
		fmt.Println("3. Mis préstamos")
		fmt.Println("4. Mi cuenta")
		fmt.Println("5. Populares")
		fmt.Println("6. Salir")
		fmt.Println()

		choice := readInput("Seleccione una opción: ")
		fmt.Println()
		switch choice {
		case "1":
			catalogMenu()
		case "2":
			cartMenu()
		case "3":
			transactionsMenu()
		case "4":
			accountMenu()
		case "5":
			popularMenu()
		case "6":
			return
		}
	}
}

// accountMenu provides account management options including balance inquiry, adding funds, and viewing transaction history
func accountMenu() {
	for {
		fmt.Println()
		fmt.Println("1. Consultar saldo")
		fmt.Println("2. Abonar usm pesos")
		fmt.Println("3. Ver historial de compras y arriendos")
		fmt.Println("4. Salir")
		fmt.Println()

		choice := readInput("Seleccione una opción: ")
		fmt.Println()
		switch choice {
		case "1":
			fmt.Println("Su saldo es de", currentUser.USM_Pesos, "usm pesos.")
		case "2":
			amountStr := readInput("Ingrese la cantidad de usm pesos a abonar: ")
			fmt.Println()
			amount, err := strconv.Atoi(amountStr)
			if err != nil || amount <= 0 {
				fmt.Println("Cantidad inválida, por favor intente de nuevo.")
				continue
			}
			err = UpdateUserUSMPesos(currentUser.ID, currentUser.USM_Pesos+amount)
			if err != nil {
				fmt.Println("Error al actualizar usm pesos:", err)
			} else {
				currentUser.USM_Pesos += amount
				fmt.Println("Nuevo saldo de", currentUser.USM_Pesos, "usm pesos.")
			}
		case "3":
			transactions, err := FetchUserTransactions(currentUser.ID)
			if err != nil {
				fmt.Println("Error al obtener el historial de transacciones:", err)
				continue
			}

			fmt.Println("------------------------------------------------------------")
			fmt.Println("| ID transacción | Nombre          | Tipo     | Fecha        | Valor   |")
			fmt.Println("------------------------------------------------------------")
			for _, transaction := range transactions {
				fmt.Printf("| %-14d | %-15s | %-8s | %-12s | %-7d |\n",
					transaction.ID, transaction.Title, transaction.Type, transaction.Transaction_date, transaction.Price)
			}
			fmt.Println("------------------------------------------------------------")
		case "4":
			return
		}
	}
}

// popularMenu displays the top 5 most popular books sorted by popularity score
func popularMenu() {
	books, err := FetchBooks()
	if err != nil {
		fmt.Println("Error al obtener el catálogo:", err)
		return
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].Popularity > books[j].Popularity
	})

	fmt.Println("------------------------------------------------------------")
	fmt.Println("| ID libro  | Nombre          | Categoría   | Popularidad  |")
	fmt.Println("------------------------------------------------------------")
	limit := 5
	if len(books) < limit {
		limit = len(books)
	}
	for i := 0; i < limit; i++ {
		book := books[i]
		fmt.Printf("| %-9d | %-15s | %-12s | %-7d |\n",
			book.ID, book.Title, book.Category, book.Popularity)
	}
	fmt.Println("------------------------------------------------------------")
}

// catalogMenu displays all available books with their details in a formatted table
func catalogMenu() {
	books, err := FetchBooks()
	if err != nil {
		fmt.Println("Error al obtener el catálogo:", err)
		return
	}

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("| ID libro  | Nombre          | Categoría   | Modalidad    | Valor   |")
	fmt.Println("----------------------------------------------------------------------")
	for _, book := range books {
		fmt.Printf("| %-9d | %-15s | %-12s | %-12s | %-7d |\n",
			book.ID, book.Title, book.Category, book.Transaction_type, book.Price)
	}
	fmt.Println("----------------------------------------------------------------------")
}

// transactionsMenu displays user's loan history with return dates, status, and allows book returns
func transactionsMenu() {
	transactions, err := FetchUserLoans(currentUser.ID)
	if err != nil {
		fmt.Println("Error al obtener los préstamos:", err)
		return
	}

	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("| ID préstamo | Nombre          | Fecha inicio  | Fecha fin    | Días restantes | Estado	 |")
	fmt.Println("-------------------------------------------------------------------------------------------")
	for _, transaction := range transactions {
		book, err := FetchBook(transaction.Book_ID)
		if err != nil {
			fmt.Println("Error al obtener el libro para el préstamo", transaction.ID, ":", err)
			continue
		}
		daysLeft, err := calculateDaysLeft(transaction.Start_date, transaction.Return_date)
		if err != nil {
			fmt.Println("Error al calcular los días restantes:", err)
			continue
		}
		fmt.Printf("| %-12d | %-15s | %-13s | %-12s | %-14d | %-7s |\n",
			transaction.ID, book.Title, transaction.Start_date, transaction.Return_date, daysLeft, transaction.Status)
	}
	fmt.Println("-------------------------------------------------------------------------------------------")

	fmt.Println()
	for {
		fmt.Println("1. Devolver libro")
		fmt.Println("2. Salir al menú principal")
		fmt.Println()
		choice := readInput("Seleccione una opción: ")
		fmt.Println()
		switch choice {
		case "1":
			loanID := readInput("Ingrese el ID del préstamo a devolver: ")
			fmt.Println()
			date := readInput("Ingrese la fecha de devolución (DD-MM-YYYY): ")
			fmt.Println()
			loanIDInt, err := strconv.Atoi(loanID)
			if err != nil {
				fmt.Println("ID inválido, por favor intente de nuevo.")
				continue
			}
			err = updateLoanStatus(loanIDInt)
			if err != nil {
				fmt.Println("Error al actualizar el estado del préstamo:", err)
			} else {
				fmt.Println("Préstamo devuelto con éxito en la fecha", date)
			}
		case "2":
			return
		}
	}
}

// cartMenu handles the shopping cart functionality including adding books, displaying cart contents, and processing orders with cart optimization
func cartMenu() {
	var cart []int

	for {
		id := readInput("Ingrese el ID del libro a agregar al carrito (o nada para salir): ")
		if id == "" {
			break
		}
		bookID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("ID inválido, por favor intente de nuevo.")
			continue
		}
		cart = append(cart, bookID)
	}
	cartBooks, err := fetchCartInfo(cart)
	if err != nil {
		fmt.Println("Error al obtener información del carrito:", err)
		return
	}

	fmt.Println("Su carro es de", len(cartBooks), "libros para un total de", calculateBookTotal(cartBooks), "usm pesos.")
	fmt.Println()
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("| Nombre          | Modalidad		| Valor   | Fecha devolución |")
	fmt.Println("---------------------------------------------------------------")
	for _, book := range cartBooks {
		if book.Transaction_type == "Arriendo" {
			fmt.Printf("| %-15s | %-12s | %-7d | %-16s |\n",
				book.Title, book.Transaction_type, book.Price, "7 días desde hoy")
		} else {
			fmt.Printf("| %-15s | %-12s | %-7d | %-16s |\n",
				book.Title, book.Transaction_type, book.Price, "-")
		}
	}
	fmt.Println("---------------------------------------------------------------")

	fmt.Println()
	for {
		confirm := readInput("Confirmar pedido (s/n): ")

		if strings.ToLower(confirm) == "n" {
			fmt.Println("Pedido cancelado. Regresando al menú principal.")
			return
		} else if strings.ToLower(confirm) != "s" {
			fmt.Println("Opción inválida, por favor intente de nuevo.")
		}

		if currentUser.USM_Pesos >= calculateBookTotal(cartBooks) {
			fmt.Println("Pedido confirmado. Gracias por su compra.")
			for _, book := range cartBooks {
				switch book.Transaction_type {
				case "Sale":
					err := CreateSale(currentUser.ID, book.ID)
					if err != nil {
						fmt.Println("Error al crear la venta para el libro", book.Title, ":", err)
					}
				case "Loan":
					err := CreateLoan(currentUser.ID, book.ID)
					if err != nil {
						fmt.Println("Error al crear el arriendo para el libro", book.Title, ":", err)
					}
				}
			}
			err := UpdateUserUSMPesos(currentUser.ID, currentUser.USM_Pesos-calculateBookTotal(cartBooks))
			if err != nil {
				fmt.Println("Error al actualizar USM_Pesos:", err)
			} else {
				currentUser.USM_Pesos -= calculateBookTotal(cartBooks)
				fmt.Println("Su nuevo saldo de USM_Pesos es:", currentUser.USM_Pesos)
			}
			return
		}

		if currentUser.USM_Pesos < calculateBookTotal(cartBooks) {
			fmt.Println("No tienes los fondos suficientes para realizar el pedido, tienes", currentUser.USM_Pesos, "usm pesos solamente pero podemos optimizar tu carrito para que no te vayas con las manos vacías.")
			fmt.Println()
			fmt.Println("1. Optimizar carrito")
			fmt.Println("2. Salir al menú principal")
			fmt.Println()
			choice := readInput("Seleccione una opción: ")

			if choice == "1" {
				optimizedCart := optimizeCart(cartBooks, currentUser.USM_Pesos)
				fmt.Println("Su carro es de", len(optimizedCart), "libros para un total de", calculateBookTotal(optimizedCart), "usm pesos.")
				fmt.Println("---------------------------------------------------------------")
				fmt.Println("| Nombre          | Modalidad		| Valor   | Fecha devolución |")
				fmt.Println("---------------------------------------------------------------")
				for _, book := range optimizedCart {
					if book.Transaction_type == "Arriendo" {
						fmt.Printf("| %-15s | %-12s | %-7d | %-16s |\n",
							book.Title, book.Transaction_type, book.Price, "7 días desde hoy")
					} else {
						fmt.Printf("| %-15s | %-12s | %-7d | %-16s |\n",
							book.Title, book.Transaction_type, book.Price, "-")
					}
				}
				fmt.Println("---------------------------------------------------------------")
				fmt.Println()
				confirm := readInput("Confirmar pedido optimizado (s/n): ")
				if strings.ToLower(confirm) == "s" {
					fmt.Println("Pedido confirmado. Gracias por su compra.")
					for _, book := range optimizedCart {
						switch book.Transaction_type {
						case "Sale":
							err := CreateSale(currentUser.ID, book.ID)
							if err != nil {
								fmt.Println("Error al crear la venta para el libro", book.Title, ":", err)
							}
						case "Loan":
							err := CreateLoan(currentUser.ID, book.ID)
							if err != nil {
								fmt.Println("Error al crear el arriendo para el libro", book.Title, ":", err)
							}
						}
					}
					err := UpdateUserUSMPesos(currentUser.ID, currentUser.USM_Pesos-calculateBookTotal(optimizedCart))
					if err != nil {
						fmt.Println("Error al actualizar usm pesos:", err)
					} else {
						currentUser.USM_Pesos -= calculateBookTotal(optimizedCart)
						fmt.Println("Su nuevo saldo de", currentUser.USM_Pesos, "usm pesos.")
					}
				}
			}
			return
		}
	}
}