package handlers

import (
	"cajitasorpresa/db"
	"cajitasorpresa/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPedidoExtraPorId(c *gin.Context) {
	pedido_extra_id_string := c.Param("pedido_extra_id")
	pedido_extra_id, err := strconv.Atoi(pedido_extra_id_string)

	if err != nil {
		log.Println("Error al convertir pedido_extra_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el pedido extra id no es valido"})
		return
	}

	var pedido models.PedidoExtra

	if err := db.DB.First(&pedido, pedido_extra_id).Error; err != nil {
		log.Println("Error al encontrar el pedido extra:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener el pedido extra"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pedido})
}

func GetPedidoExtraPorIdPedido(c *gin.Context) {
	pedido_id_string := c.Param("pedido_id")
	pedido_id, err := strconv.Atoi(pedido_id_string)

	if err != nil {
		log.Println("Error al convertir pedido_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el pedido id no es valido"})
		return
	}

	var pedido []models.PedidoExtra

	if err := db.DB.Where("pedido_id = ?", pedido_id).Find(&pedido).Error; err != nil {
		log.Println("Error al encontrar el pedido extra:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener el pedido extra"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pedido})
}
