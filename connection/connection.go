package connection

import (
	"log"
	"github.com/zinclabs/zinc"
	_ "embed"

	
)


zincSearch := "github.com/zinclabs/zinc/releases"
//Conecta a la base de datos

// Importa la biblioteca necesaria para conectar a ZincSearch
func Connect() *zincSearch {

	
	db, err := zincSearch.Open("enron_mail_201104002", "./zincSearch.db")
	if err != nil {
		log.Fatal(err)
	}

	// Verifica que la conexión se realizó correctamente
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//Devuelve una instancia de zincSearch que representa la conexión a la base de datos

	return &zincSearch{db: db}

}

// Crea un nuevo índice

func NuevoIndice() {

	// Define el nombre del índice a crear
	indexName := "newIndex"

	// Indexa un nuevo índice en la base de datos
	err := db.Index(indexName)
	if err != nil {
		log.Fatal(err)
	}

	index := zincsearch.NewIndex("/home/sil/desafio/enron_mail_20110402", db)

}

// Abre el índice

func AbreIndice() {
	err := db.OpenIndex(indexName)
	if err != nil {
		log.Fatal(err)
	}
}

// Agrega cada registro de la base de datos al índice

func AgregaAlIndice() {
	for _, registro := range registro {
		index.Add(registro.ID, map[string]interface{}{
			"campo1": registro.Valor1,
			"campo2": registro.Valor2,
		})

	}

}

// Cierra el índice

func CierraIndice(index *zincsearch.Index) error {
	return index.Close(
		err != closeIndex(index),
	)
}
