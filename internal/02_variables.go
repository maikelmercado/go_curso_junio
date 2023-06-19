package internal

import (
	"fmt"
	"time"
)

//variables o funcion deben ser en mayuscula la primera letras
// importar

func Variables(){
//declaratia
var intcomun int 

//asignacion
intde32 := int32(10)


println(intcomun, intde32) 

} 

//variables 
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func VariablesVarias1(){

	Nombre = "pedro"
	Estado = true
	Sueldo = 1574.12
	Fecha = time.Now()
	//now tiene funciones adicionarles fechas o que devuelva fecha el objeto time


	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)


}