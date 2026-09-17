package handlers

import (
	"cajitasorpresa/db"
	"cajitasorpresa/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetExtras(c *gin.Context) {
	var extras []models.Extra

	if err := db.DB.Find(&extras).Error; err != nil {
		log.Println("Error al encontrar los extras:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudieron obtener los extras"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": extras})
}

func GetExtraByID(c *gin.Context) {
	extra_id_string := c.Param("extra_id")
	extra_id, err := strconv.Atoi(extra_id_string)

	if err != nil {
		log.Println("Error al convertir el extra_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el extra id no es valido"})
		return
	}

	var extra models.Extra

	if err := db.DB.First(&extra, extra_id).Error; err != nil {
		log.Println("Error al encontrar el extra:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener el extra"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": extra})
}
