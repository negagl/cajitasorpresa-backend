package handlers

import (
	"cajitasorpresa/db"
	"cajitasorpresa/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBoxes(c *gin.Context) {
	var boxes []models.Box

	if err := db.DB.Model(&models.Box{}).Preload("IngredientesBox").Find(&boxes).Error; err != nil {
		log.Println("Error al obtener los boxes:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener boxes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": boxes})
}

func GetBoxPorId(c *gin.Context) {
	box_id_string := c.Param("box_id")
	box_id, err := strconv.Atoi(box_id_string)

	if err != nil {
		log.Println("Error al convertir el box_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el box id no es valido"})
		return
	}

	var box models.Box

	if err := db.DB.Model(&models.Box{}).Preload("IngredientesBox").First(&box, box_id).Error; err != nil {
		log.Println("Error al obtener el box:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener el box"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": box,
	})
}

func GetIngredientesPorBox(c *gin.Context) {
	box_id_string := c.Param("box_id")
	box_id, err := strconv.Atoi(box_id_string)

	if err != nil {
		log.Println("Error al convertir el box_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el box id no es valido"})
		return
	}

	var ingredientes []models.IngredienteBox

	if err := db.DB.Where("box_id = ?", box_id).Find(&ingredientes).Error; err != nil {
		log.Println("Error al obtener los ingredientes:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudieron obtener los ingredientes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingredientes})
}
