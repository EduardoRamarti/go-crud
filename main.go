package main

import (
	"bufio" //Esta librería proporciona una interfaz conveniente para leer datos de entrada de manera más eficiente y con menos sobrecarga
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// variables globales
var reader *bufio.Reader
var id int

// bufio.Reader es un tipo de dato en Go (Golang) que proporciona una funcionalidad adicional para realizar operaciones de lectura más eficientes y flexibles en comparación con las funciones de entrada estándar.

// Estructuras
type User struct {
	id       int
	username string
	email    string
	age      int
}

// Mapas
var users map[int]User

// creando las funciones
func crearUsuario() {
	clearConsole()
	fmt.Print("Ingresa el username: ")
	username := readLines()
	fmt.Print("Ingresa el email: ")
	email := readLines()
	fmt.Print("Ingresa la edad: ")
	age, err := strconv.Atoi(readLines())

	if err != nil {
		panic("No es posible convertir de un string a un entero")
	}

	id++
	user := User{id, username, email, age}
	users[id] = user

	fmt.Println("Usuario creado exitosamente")
	fmt.Println()
}

func listarUsuario() {
	clearConsole()
	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}
	fmt.Println()
}

func actualizarUsuario() {
	fmt.Print("Ingresa el id del usuario a actualizar: ")
	id, err := strconv.Atoi(readLines())
	if err != nil {
		panic("No es posible convertir de un string a un entero")
	}

	currentUser := users[id]

	fmt.Print("Ingresa el username: ")
	currentUser.username = readLines()
	fmt.Print("Ingresa el email: ")
	currentUser.email = readLines()
	fmt.Print("Ingresa la edad: ")
	currentUser.age, _ = strconv.Atoi(readLines())

	users[id] = currentUser

	fmt.Println("Usuario actualizado exitosamente")
	fmt.Println()
}

func eliminarUsuario() {
	fmt.Println("Ingresa el id del usuario a eliminar")
	id, err := strconv.Atoi(readLines())
	if err != nil {
		panic("No es posible convertir de un string a un entero")
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	} else {
		fmt.Println("No existe ese usuario")
	}

	fmt.Println("Usuario eliminado exitosamente")
	fmt.Println()
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLines() string {

	//Se lee la opción ingresada por el usuario utilizando el lector reader. La función ReadString('\n') lee una línea completa de texto ingresada por el usuario hasta que se presiona la tecla Enter. La opción ingresada se asigna a la variable option y cualquier error que ocurra se asigna a la variable err.

	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor!")
	} else {

		return strings.TrimSuffix(option, "\n")
		//Se elimina el carácter de nueva línea ("\n") del final de la opción ingresada por el usuario. Esto se hace utilizando la función TrimSuffix del paquete strings
	}
}

func main() {

	var option string
	users = make(map[int]User)

	//Se inicializa la variable reader con un nuevo lector de entrada estándar (os.Stdin). Esto permite leer los datos ingresados por el usuario desde la consola.
	reader = bufio.NewReader(os.Stdin)

	for {
		//Opciones
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")

		//capturando la opcion
		fmt.Print("Ingresa una opcion: ")
		option = readLines()

		if option == "quit" || option == "q" {
			break
		}

		//evaluacion de la opcion seleccionada
		switch option {
		case "a", "crear", "A", "Crear":
			crearUsuario()
		case "b", "alistar", "B", "Alistar":
			listarUsuario()
		case "c", "actualizar", "C", "Actualizar":
			actualizarUsuario()
		case "d", "eliminar", "D", "Eliminar":
			eliminarUsuario()
		default:
			fmt.Println("Opcion no valida")
		}
	}

	fmt.Println("Adios")

}
