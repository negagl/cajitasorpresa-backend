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
			Descripcion: "- Jugo natural con o sin azúcar (mora, maracuyá o naranja según disponibilidad) \n" +
				"- Sandwich de jamón y queso con pan artesanal / Waffles con mermelada y fresas \n" +
				"- Granola, yogurt y fruta de temporada picada \n" +
				"- Galletas \n" +
				"- Dos chocolates \n" +
				"- Tarjeta personalizada \n" +
				"- Decoración \n" +
				"- Foto impresa de la amistad o pareja \n",
			PrecioBase: 70000,
			Imagen:     "",
		},
		{
			Nombre: "Box clasico",
			Descripcion: "- Jugo natural con o sin azúcar (mora, maracuyá o naranja según disponibilidad) \n" +
				"- Sandwich de jamón y queso con pan artesanal / Waffles con mermelada y fresas \n" +
				"- Fruta de temporada picada \n" +
				"- Postre \n" +
				"- Tarjeta personalizada con mensaje \n" +
				"- Decoración \n",
			PrecioBase: 67000,
			Imagen:     "",
		},
		{
			Nombre: "Box esencial",
			Descripcion: "- Jugo natural con o sin azúcar (mora, maracuyá o naranja según disponibilidad) \n" +
				"- Sandwich de jamón y queso con pan artesanal / Waffles con mermelada y fresas / Deditos integrales \n" +
				"- Postre / Mousse de mango sin azúcar \n" +
				"- Chocolate / Galleta integral \n" +
				"- Tarjeta personalizada \n" +
				"- Decoración \n",
			PrecioBase: 55000,
			Imagen:     "",
		},
		{
			Nombre: "Box saludable",
			Descripcion: "- Jugo natural sin azúcar (mora, maracuyá o naranja según disponibilidad) \n" +
				"- Sandwich con jamón de pollo o de pavo y queso bajo en grasa con pan integral / Deditos integrales \n" +
				"- Granola, yogurt griego y fruta de temporada picada \n" +
				"- Barra de cereal Tosh \n" +
				"- Mousse de mango sin azúcar \n" +
				"- Tarjeta personalizada \n" +
				"- Foto impresa \n",
			PrecioBase: 75000,
			Imagen:     "",
		},
		{
			Nombre: "Box sorpresa premium",
			Descripcion: "- Jugo natural con o sin azúcar (mora, maracuyá o naranja según disponibilidad) \n" +
				"- Sandwich de jamón y queso con pan artesanal / Waffles con mermelada y fresas / Bandejita de costeñitos (carimañolas, butifarras, bollo de mazorca y queso costeño) \n" +
				"- Granola, yogurt y fruta de temporada picada \n" +
				"- Galletas \n" +
				"- Dos chocolates \n" +
				"- Tarjeta personalizada \n" +
				"- Decoración \n" +
				"- Foto impresa de la amistad o pareja colgantes / Cuadro en madera con foto \n" +
				"- Mini cake de cumpleaños \n" +
				"- Hatsu en lata / Cerveza en lata \n" +
				"- Mini bouquet de flores en claveles o margaritas / Mug personalizado con sobre de café \n",
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
