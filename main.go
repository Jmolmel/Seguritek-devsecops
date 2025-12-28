// CI trigger
package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", ImageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              ":" + port,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Println("Servidor iniciado en el puerto", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
