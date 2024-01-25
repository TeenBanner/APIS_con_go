package funciones

import (
	"fmt"
)

// en go  los middlewares tienen el patron de diseno de funciones de primer ordens
// el middleware va a recibir una de estas funciones como parametro y va aretornar una funcion que recibe un string
func Saludar(name string) {
	fmt.Println("Hola, ", name)
}

func Despedirse(name string) {
	fmt.Println("Adios", name)
}
