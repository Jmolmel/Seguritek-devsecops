package main

import (
	"fmt"
	"log"
	"net/http"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("[WARN] Método HTTP no permitido:", r.Method)
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintln(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>DevSecOps</title>
		</head>
		<body>
			<img src="/static/logo.png" alt="Logo" width="300"/>
		</body>
		</html>
	`)
}
