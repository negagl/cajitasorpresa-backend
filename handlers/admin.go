package handlers

import (
	"cajitasorpresa/db"
	"cajitasorpresa/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetListaCompras(c *gin.Context) {
	desdeQuery := c.Query("desde")
	hastaQuery := c.Query("hasta")

	// Me aseguro que desde y hasta sean fechas y que sea un rango correcto
	desde, err := time.Parse("2006-01-02", desdeQuery)
	if err != nil {
		log.Println("Error al convertir fecha desde a tipo Time:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "la fecha o el formato de fecha es incorrecto"})
		return
	}

	hasta, err := time.Parse("2006-01-02", hastaQuery)
	if err != nil {
		log.Println("Error al convertir fecha hasta a tipo Time:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "la fecha o el formato de fecha es incorrecto"})
		return
	}

	// Verificamos que el rango de fechas es correcto (desde es menor que hasta)
	if hasta.Before(desde) {
		log.Println("Error al validar el rango de fechas, desde es mayor que hasta")
		c.JSON(http.StatusBadRequest, gin.H{"error": "la fecha 'hasta' no puede ser menor que la fecha 'desde'"})
		return
	}

	hasta = hasta.AddDate(0, 0, 1)

	// Buscamos la lista de pedidos para luego extraer los ingredientes de los boxes y los extra
	var pedidos []models.Pedido
	if err := db.DB.Preload("ItemsExtra.Extra").Where("fecha_entrega >= ? AND fecha_entrega < ?", desde, hasta).Find(&pedidos).Error; err != nil {
		log.Println("Error al obtener los pedidos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener la lista de compras"})
		return
	}

	listaCompras := map[string]float64{}

	// Por cada pedido buscamos los ingredientes del Box ID y los extras ya los precargamos.
	for _, pedido := range pedidos {
		var ingredientes []models.IngredienteBox
		if err := db.DB.Where("box_id = ?", pedido.BoxID).Find(&ingredientes).Error; err != nil {
			log.Println("Error al obtener la lista de ingradientes del box", pedido.BoxID, ":", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo obtener la lista de compras"})
			return
		}

		for _, ingrediente := range ingredientes {
			listaCompras[ingrediente.Nombre] += ingrediente.Cantidad
		}

		for _, itemExtra := range pedido.ItemsExtra {
			listaCompras[itemExtra.Extra.Nombre] += float64(itemExtra.Cantidad)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": listaCompras})
}
