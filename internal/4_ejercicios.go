package internal

import (
	"strconv"
)

func ConvNumerito(letras string) (int, string) {

	numero1, err := strconv.Atoi(letras) //_ no asigna a ninguna variable
	if err != nil {
		return 0, "Hubo un error" + err.Error()

	}
	if numero1 > 100 {
		return numero1, "Es mayor a 100"
	} else {
		return numero1, "es menor de 100"
	}

}
