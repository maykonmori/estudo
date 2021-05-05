package routes

import (
	"net/http"

	"github.com/maykonmori1993/estudo/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
