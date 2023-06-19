package internal

import "strconv"
import "fmt"

// funciones
//los metodos no tiene retornos
//recibir un valor numerico y convertir en string

//parametro de tipo numerico, va devolver de salida




func VariablesVarias(){

	Nombre = "pedro"
	Estado = true
	Sueldo = 1574.12



	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	

}
func ConviertoaTexto(numero int) (bool, string) {
	
	texto := strconv.Itoa(numero)
	return true, texto

}
