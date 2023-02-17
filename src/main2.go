package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	/*
		//funcion que recorre cuantas carpetas hay
		type walkFunc func(path string, info fs.FileInfo, err error) error
		count := 0
		countDir := 0
		var pathEmailToWalk = "/home/user/Descargas/enron_mail_20110402/maildir"
		//variable para guardar direccion en lector os
		dirFiles := os.DirFS(pathEmailToWalk)
		//funcion que recorre la direccion path
		fs.WalkDir(dirFiles, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				panic(err)
			}
			//Detecta si es carpeta de lo contrario es archivo
			if d.Type().IsDir() {

				//contador de carpeta
				countDir++
				//muestra el contado de carpeta
				println(" es carpeta N°: ", countDir, "\n")
			} else {
				//contado de archivo
				count++
				//muestra el contador de archivo
				fmt.Println("Archivo N°: ", count)
			}

			return nil

			})
			println("hay: ", countDir, " carpetas y ", count, " Archivos :)")
	*/

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome hola mundo"))
	})
	http.ListenAndServe(":3000", r)

}
