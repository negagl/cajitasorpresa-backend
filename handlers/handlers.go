package handlers

import (
	"cajitasorpresa/db"
	"cajitasorpresa/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBoxes(c *gin.Context) {
	var boxes []models.Box

	if err := db.DB.Find(&boxes).Error; err != nil {
		log.Println("Error al obtener los boxes:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener boxes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"boxes": boxes,
	})
}
