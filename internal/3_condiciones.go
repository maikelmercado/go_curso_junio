package internal

import (
	"fmt"
	"runtime" // ambiente de ejecucion donde esta corriendo nuestro sistema
)

func Condicion() {

	if os := runtime.GOOS; os == "Linux." || os == "windows." {

		fmt.Println(" esto no es linux")

	} else {

		fmt.Println("Estyo es window windows")

	}

	switch os := runtime.GOOS; os {
	case "linux.":
		fmt.Println("esto es linux")
	case "darwin.":
		fmt.Println("Esto es darwin")
	case "windows.":
		fmt.Println("Esto es windows")
	default:
		fmt.Printf("%s \n", os) //permite que se formatee el texto, el argumento es un strin un salto de linea y lo va mostrar por pantalla
	}
}
