package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		http.Handle("/", http.FileServer(http.Dir("public")))
		log.Println("ejecutando server en http://localhost:8080")
		log.Println(http.ListenAndServe(":8080", nil))
	*/

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	log.Println("ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))

	// http.template()
	// {{}} para insertar codigo GO dentro del Html
	// {{.Nombre}}  //el punto (.) hace referencia a la variable pasada

}
