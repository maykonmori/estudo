package main

import (
	"net/http"

	"github.com/maykonmori1993/estudo/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
