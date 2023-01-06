package main

import (
	_ "embed"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zinclabs/zinc"
)




func main() {
    
	router := mux.NewRouter().StrictSlash(value true)
	
	//ruta principal
	router.HandleFunc("/", Index)

	router.HandleFunc("/connect", connection.connection).Methods(Connect)

	// Crea un nuevo índice
	router.HandleFunc("/nuevoIndice", connection.connection).Methods(NuevoIndice)

	// Abre el índice
	router.HandleFunc("/abreIndice", connection.connection).Methods(AbreIndice)

	// Agrega cada registro de la base de datos al índice
	router.HandleFunc("/agregaAlIndice", connection.connection).Methods(AgregaAlIndice)

	// Cierra el índice
	router.HandleFunc("/cierraIndice", connection.connection).Methods(CierraIndice)

	server := http.ListenAndServe(":4080", router)
	  
	  log.Fatal(server)

	fmt.Println("¡Base de datos indexada con éxito!")
}
