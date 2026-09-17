package handlers

import (
	"cajitasorpresa/db"
	"cajitasorpresa/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPersonalizacionPorId(c *gin.Context) {
	personalizacion_id_string := c.Param("personalizacion_id")
	personalizacion_id, err := strconv.Atoi(personalizacion_id_string)

	if err != nil {
		log.Println("Error al convertir personalizacion_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el personalizacion id no es valido"})
		return
	}

	var personalizacion models.PersonalizacionPedido

	if err := db.DB.First(&personalizacion, personalizacion_id).Error; err != nil {
		log.Println("Error al encontrar la personalizacion:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener la personalizacion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": personalizacion})
}

func GetPersonalizacionPorIdPedido(c *gin.Context) {
	pedido_id_string := c.Param("pedido_id")
	pedido_id, err := strconv.Atoi(pedido_id_string)

	if err != nil {
		log.Println("Error al convertir pedido_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el pedido id no es valido"})
		return
	}

	var personalizacion []models.PersonalizacionPedido

	if err := db.DB.Where("pedido_id = ?", pedido_id).Find(&personalizacion).Error; err != nil {
		log.Println("Error al encontrar la personalizacion:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener la personalizacion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": personalizacion})
}
