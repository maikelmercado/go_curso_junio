package internal

import (
	"bufio" // lectura de escritura y archivos, todo lo que entra En bufio entra en modo texto
	"fmt"
	"os"
	"strconv" // permite convertir, albaetico a un entero 
)



var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumero(){
	scanner := bufio.NewScanner(os.Stdin)      //inicializar un objeto,   packete os stadin input es el teclado BUFIO paquete completo
	
	fmt.Println("Ingrese numero 1:")
	if  scanner.Scan(){
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("EL dato ingresado es incorreecto"+err.Error())
		}
	}

	fmt.Println("Ingrese numero 2:")
	if  scanner.Scan(){
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("EL dato ingresado es incorreecto"+err.Error())
		}
	}

	fmt.Println("Ingrese leyenda:")
	if  scanner.Scan(){
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)

}
	
