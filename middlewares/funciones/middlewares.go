package funciones

import (
	"fmt"
	"time"
)

// podemos crear un tipo para hacer mas entendible el middleware
type MyFunction func(string) // es una funcion que recibe un string

// funcion de primer orden esta recibe una funcion y devuelve una funcion
func MiddlewareLog(f MyFunction) /*recibe una funcion que recibe un string*/ MyFunction /* retorna una funcion que usa como parametro un string*/ {
	//retorna una funcion
	return func(name string) /* esta funcion recibe un parametro name de tipo string*/ { // cuando retornemos esta funcion va a ejecutar su funcionalidAD
		// cuando la funcion sea llamada imprime la fecha
		fmt.Println("inicio: ", time.Now().Format("  2006-01-02 15:04:05"))
		// Ejecuta la funcion que recibe de parametro
		f(name) // se le pasa el parametro name de la funcion que esta retornando
		fmt.Println("fin de la funcion: ", time.Now())
	}
}
