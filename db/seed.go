package db

import (
	"cajitasorpresa/models"
	"log"
)

func Seed() {
	// Crear los boxes
	boxes := []models.Box{
		{
			Nombre: "Box amor y amistad",
			Descripcion: "- Jugo natural con o sin azúcar (mora, maracuyá o naranja según disponibilidad)" +
				"- Sandwich de jamón y queso con pan artesanal / Waffles con mermelada y fresas" +
				"- Granola, yogurt y fruta de temporada picada" +
				"- Galletas" +
				"- Dos chocolates" +
				"- Tarjeta personalizada" +
				"- Decoración" +
				"- Foto impresa de la amistad o pareja",
			PrecioBase: 70000,
			Imagen:     "",
		},
		{
			Nombre: "Box clasico",
			Descripcion: "- Jugo natural con o sin azúcar (mora, maracuyá o naranja según disponibilidad)" +
				"- Sandwich de jamón y queso con pan artesanal / Waffles con mermelada y fresas" +
				"- Fruta de temporada picada" +
				"- Postre" +
				"- Tarjeta personalizada con mensaje" +
				"- Decoración",
			PrecioBase: 67000,
			Imagen:     "",
		},
		{
			Nombre: "Box esencial",
			Descripcion: "- Jugo natural con o sin azúcar (mora, maracuyá o naranja según disponibilidad)" +
				"- Sandwich de jamón y queso con pan artesanal / Waffles con mermelada y fresas / Deditos integrales" +
				"- Postre / Mousse de mango sin azúcar" +
				"- Chocolate / Galleta integral" +
				"- Tarjeta personalizada" +
				"- Decoración",
			PrecioBase: 55000,
			Imagen:     "",
		},
		{
			Nombre: "Box saludable",
			Descripcion: "- Jugo natural sin azúcar (mora, maracuyá o naranja según disponibilidad)" +
				"- Sandwich con jamón de pollo o de pavo y queso bajo en grasa con pan integral / Deditos integrales" +
				"- Granola, yogurt griego y fruta de temporada picada" +
				"- Barra de cereal Tosh" +
				"- Mousse de mango sin azúcar" +
				"- Tarjeta personalizada" +
				"- Foto impresa",
			PrecioBase: 75000,
			Imagen:     "",
		},
		{
			Nombre: "Box sorpresa premium",
			Descripcion: "- Jugo natural con o sin azúcar (mora, maracuyá o naranja según disponibilidad)" +
				"- Sandwich de jamón y queso con pan artesanal / Waffles con mermelada y fresas / Bandejita de costeñitos (carimañolas, butifarras, bollo de mazorca y queso costeño)" +
				"- Granola, yogurt y fruta de temporada picada" +
				"- Galletas" +
				"- Dos chocolates" +
				"- Tarjeta personalizada" +
				"- Decoración" +
				"- Foto impresa de la amistad o pareja colgantes / Cuadro en madera con foto" +
				"- Mini cake de cumpleaños" +
				"- Hatsu en lata / Cerveza en lata" +
				"- Mini bouquet de flores en claveles o margaritas / Mug personalizado con sobre de café",
			PrecioBase: 150000,
			Imagen:     "",
		},
	}

	for _, box := range boxes {
		if err := DB.FirstOrCreate(&box, models.Box{Nombre: box.Nombre}).Error; err != nil {
			log.Println("Error al sembrar el box:", box.Nombre, err)
		}
	}

	log.Println("Seed de boxes completado.")
}
