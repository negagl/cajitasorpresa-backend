package db

import (
	_ "embed"
	"encoding/json"
	"log"

	"cajitasorpresa/models"
)

// seed_data.json viaja DENTRO del binario (go:embed), así que no depende de
// la carpeta desde donde corras el programa.
//
//go:embed seed_data.json
var seedJSON []byte

// Structs propios del seed, a propósito NO son models.Box: si le pasaras a GORM
// un Box con IngredientesBox lleno, intentaría insertar los ingredientes por su
// cuenta y se duplicarían cada vez que corras el seed.
type seedIngrediente struct {
	Nombre   string  `json:"nombre"`
	Cantidad float64 `json:"cantidad"`
	Unidad   string  `json:"unidad"`
}

type seedBox struct {
	Nombre       string            `json:"nombre"`
	Descripcion  string            `json:"descripcion"`
	PrecioBase   float64           `json:"precio_base"`
	Imagen       string            `json:"imagen"`
	Ingredientes []seedIngrediente `json:"ingredientes"`
}

type seedExtra struct {
	Nombre      string           `json:"nombre"`
	Tipo        models.TipoExtra `json:"tipo"`
	Descripcion string           `json:"descripcion"`
	PrecioBase  float64          `json:"precio_base"`
}

type seedData struct {
	Boxes  []seedBox   `json:"boxes"`
	Extras []seedExtra `json:"extras"`
}

// Seed sincroniza la base con seed_data.json. Es idempotente: se puede correr
// en cada arranque. Si no existe el registro lo crea; si existe, le actualiza
// los campos (por eso, cambiar un precio en el JSON y reiniciar basta).
func Seed() {
	var data seedData
	if err := json.Unmarshal(seedJSON, &data); err != nil {
		log.Fatal("seed_data.json no es un JSON válido:", err)
	}

	for _, sb := range data.Boxes {
		var box models.Box
		err := DB.Where(models.Box{Nombre: sb.Nombre}).
			Assign(models.Box{Descripcion: sb.Descripcion, PrecioBase: sb.PrecioBase, Imagen: sb.Imagen}).
			FirstOrCreate(&box).Error
		if err != nil {
			log.Println("Error al sembrar el box:", sb.Nombre, err)
			continue // sin box.ID no se pueden crear sus ingredientes
		}

		for _, si := range sb.Ingredientes {
			var ing models.IngredienteBox
			err := DB.Where(models.IngredienteBox{BoxID: box.ID, Nombre: si.Nombre}).
				Assign(models.IngredienteBox{Cantidad: si.Cantidad, UnidadDeMedida: si.Unidad}).
				FirstOrCreate(&ing).Error
			if err != nil {
				log.Println("Error al sembrar el ingrediente:", sb.Nombre, "/", si.Nombre, err)
			}
		}
	}

	for _, se := range data.Extras {
		var extra models.Extra
		err := DB.Where(models.Extra{Nombre: se.Nombre}).
			Assign(models.Extra{Tipo: se.Tipo, Descripcion: se.Descripcion, PrecioBase: se.PrecioBase}).
			FirstOrCreate(&extra).Error
		if err != nil {
			log.Println("Error al sembrar el extra:", se.Nombre, err)
		}
	}

	log.Println("Seed completado.")
}
