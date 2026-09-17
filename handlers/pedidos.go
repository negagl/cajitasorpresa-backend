package handlers

import (
	"cajitasorpresa/db"
	"cajitasorpresa/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPedidos(c *gin.Context) {
	var pedidos []models.Pedido

	if err := db.DB.Model(&models.Pedido{}).Preload("ItemsExtra").Preload("Personalizacion").Find(&pedidos).Error; err != nil {
		log.Println("Error al encontrar los pedidos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudieron obtener los pedidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pedidos})
}

func GetPedidoPorId(c *gin.Context) {
	pedido_id_string := c.Param("pedido_id")
	pedido_id, err := strconv.Atoi(pedido_id_string)

	if err != nil {
		log.Println("Error al convertir el pedido_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el id del pedido no es valido"})
		return
	}

	var pedido models.Pedido

	if err := db.DB.Model(&models.Pedido{}).Preload("ItemsExtra").Preload("Personalizacion").First(&pedido, pedido_id).Error; err != nil {
		log.Println("Error al encontrar el pedido:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener el pedido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pedido})
}

func GetPedidosPorBoxId(c *gin.Context) {
	box_id_string := c.Param("box_id")
	box_id, err := strconv.Atoi(box_id_string)

	if err != nil {
		log.Println("Error al convertir el box_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "el id del box no es valido"})
		return
	}

	var pedidos []models.Pedido

	if err := db.DB.Model(&models.Pedido{}).Preload("ItemsExtra").Preload("Personalizacion").Where("box_id = ?", box_id).Find(&pedidos).Error; err != nil {
		log.Println("Error al encontrar los pedidos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener los pedidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pedidos})
}
