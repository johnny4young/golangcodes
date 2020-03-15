package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/golangcodes/routes"

	"github.com/golangcodes/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración")
	}
	// Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// inicia server
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}

	log.Println("Iniciando el servidor en http://localhost:8080")
	server.ListenAndServe()
	log.Println("Final del programa")
}
