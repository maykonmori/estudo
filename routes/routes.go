package routes

import "net/http"

func CarregaRotas() {
	http.HandleFunc("/", Index)
}
