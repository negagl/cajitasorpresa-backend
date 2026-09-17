package db

import (
	"cajitasorpresa/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("cajitasorpresa.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Fallo al conectar a la BD:", err)
	}

	if err := db.AutoMigrate(&models.Box{}, &models.IngredienteBox{}, &models.Extra{}, &models.Pedido{}, &models.PedidoExtra{}, &models.PersonalizacionPedido{}); err != nil {
		log.Fatal("Fallo al realizar las migraciones:", err)
	}

	DB = db
}
