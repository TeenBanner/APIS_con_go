package funciones

import (
	"fmt"
	"time"
)

func Saludar(name string) {
	fmt.Println(time.Now())
	fmt.Println("Hola, ", name)
}

func Despedirse(name string) {
	fmt.Println(time.Now())
	fmt.Println("Adios", name)
}
