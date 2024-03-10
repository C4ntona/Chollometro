package main

import (
	"chollometro/estructuras"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	coche := estructuras.NewProducto("Renault Megane", "Coche de pruebas", 1000.0, 5)
	if err := ejemploWriteJSON(coche); err != nil {
		fmt.Println("Error:", err)
	}
}

func ejemploWriteJSON(p *estructuras.Producto) error {
	dir := "products"
	fileName := strings.Replace(p.GetNombre(), " ", "", -1) + ".json"
	if err := createDirIfNotExist(dir); err != nil {
		return err
	}
	file, err := os.Create(dir + "/" + fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(p); err != nil {
		return err
	}
	return nil
}

func createDirIfNotExist(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			return fmt.Errorf("Error al crear el directorio: %s", err)
		}
	}
	return nil
}
