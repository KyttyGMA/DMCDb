package main

import "fmt"

func main() {
	// Crear una nueva instancia de la base de datos
	db := NewDualMemoryDB()

	// Insertar datos
	db.Insert("mi_clave", "Este es un valor de prueba.")

	// Recuperar y mostrar datos completos
	fullData := db.GetFullData("mi_clave")
	fmt.Println("Datos completos:", fullData)

	// Recuperar y mostrar metadatos
	metadata := db.GetMetadata("mi_clave")
	fmt.Println("Metadatos:", metadata)
}
