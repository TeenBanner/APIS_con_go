package main

import "middlewares/funciones"

// en go los middlewares son funciones de primer orden estas nos permiten agregarle funcionalidades extra a una funcion
// hablando en un contexto web los middlewares se usan como intermediarios entre una solicitud entrente y la logica de la aplicacion
// estos pueden realizar tareas especificas antes de que la solicitud llegue a su respectivo controlador

func execute(name string, f funciones.MyFunction) /* recibe una funcion que tiene una firma de string*/ {
	f(name)
}
func main() {
	name := "Comunidad EDteam"
	// estos middlewares nos van a permitir escribir funciones y enforcarnos en su verdadera funcionalidad estandarizando funcionalidades comunes entre varias funciones
	execute(name, funciones.MiddlewareLog(funciones.Saludar)) // la funcion execute recibe una funcion que tiene como parametro un string
	execute(name, funciones.MiddlewareLog(funciones.Despedirse))
}
