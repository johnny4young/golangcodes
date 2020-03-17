package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/golangcodes/commons"

	"github.com/urfave/negroni"

	"github.com/golangcodes/routes"

	"github.com/golangcodes/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto para el servidor web")
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
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Iniciando el servidor en http://localhost:%d", commons.Port)
	server.ListenAndServe()
	log.Println("Final del programa")
}
