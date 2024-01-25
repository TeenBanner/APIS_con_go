package main

import "middlewares/funciones"

func execute(name string, f func(string)) {
	f(name)
}
func main() {
	name := "Comunidad EDteam"
	execute(name, funciones.Saludar)
	execute(name, funciones.Despedirse)
}
