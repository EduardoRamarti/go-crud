package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//creando las funciones

func crearUsuario() {
	fmt.Println("Usuario creado exitosamente")
}
func listarUsuario() {
	fmt.Println("Usuario listado exitosamente")
}
func actualizarUsuario() {
	fmt.Println("Usuario actualizado exitosamente")
}
func eliminarUsuario() {
	fmt.Println("Usuario eliminado exitosamente")
}

func main() {

	var reader *bufio.Reader
	var option string
	reader = bufio.NewReader(os.Stdin)

	//Opciones
	fmt.Println("A) Crear")
	fmt.Println("B) Alistar")
	fmt.Println("C) Actualizar")
	fmt.Println("D) Eliminar")

	//capturando la opcion
	fmt.Println("Ingresa una opcion:")
	option, err := reader.ReadString('\n')

	if err != nil {
		panic("No es posible obtener el valor!")
	}

	option = strings.TrimSuffix(option, "\n")

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
