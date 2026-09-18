package handlers

import (
	"cajitasorpresa/db"
	"cajitasorpresa/models"
	"log"
	"net/http"
	"strconv"
	"time"

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

func CrearPedido(c *gin.Context) {
	var newPedidoInput CrearPedidoInput

	if err := c.ShouldBindJSON(&newPedidoInput); err != nil {
		log.Println("Error al parsear el body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "json invalido"})
		return
	}

	// Se busca el box primero
	var box models.Box
	if err := db.DB.First(&box, newPedidoInput.BoxID).Error; err != nil {
		log.Println("Error al buscar el box:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error al buscar el box"})
		return
	}

	precioTotal := box.PrecioBase

	// Buscamos los extras ahora
	for _, extraInput := range newPedidoInput.Extras {
		var extra models.Extra
		if err := db.DB.First(&extra, extraInput.ExtraID).Error; err != nil {
			log.Println("Error al buscar los extras:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error al buscar los extras"})
			return
		}

		// Calculamos el precio de cada extra y lo sumamos al total
		precioTotal += extra.PrecioBase * float64(extraInput.Cantidad)
	}

	// Creamos el pedido
	newPedido := models.Pedido{
		NombreCliente:   newPedidoInput.NombreCliente,
		TelefonoCliente: newPedidoInput.TelefonoCliente,
		BoxID:           box.ID,
		SaborJugo:       newPedidoInput.SaborJugo,
		ConAzucar:       newPedidoInput.ConAzucar,
		MensajeTarjeta:  newPedidoInput.MensajeTarjeta,
		FotoURL:         newPedidoInput.FotoURL,
		PrecioTotal:     precioTotal,
		FechaPedido:     time.Now(),
		FechaEntrega:    newPedidoInput.FechaEntrega,
		Estado:          models.Pendiente,
	}

	if err := db.DB.Create(&newPedido).Error; err != nil {
		log.Println("Error al crear el pedido:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creando el pedido"})
		return
	}

	// Creamos los extras
	for _, extraInput := range newPedidoInput.Extras {
		pedidoExtra := models.PedidoExtra{
			PedidoID: newPedido.ID,
			ExtraID:  extraInput.ExtraID,
			Cantidad: extraInput.Cantidad,
		}

		if err := db.DB.Create(&pedidoExtra).Error; err != nil {
			log.Println("Error al crear los pedidos extra:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creando los pedidos extra"})
			return
		}
	}

	// Creamos las personalizaciones
	for _, personalizacionInput := range newPedidoInput.Personalizacion {
		personalizacion := models.PersonalizacionPedido{
			PedidoID:  newPedido.ID,
			Categoria: personalizacionInput.Categoria,
			Eleccion:  personalizacionInput.Eleccion,
		}

		if err := db.DB.Create(&personalizacion).Error; err != nil {
			log.Println("Error al crear las personalizaciones:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creando las personalizaciones"})
			return
		}
	}

	// devolvemos el objeto completo con sus relaciones
	var pedidoCreado models.Pedido
	if err := db.DB.Preload("ItemsExtra").Preload("Personalizacion").First(&pedidoCreado, newPedido.ID).Error; err != nil {
		log.Println("Error al buscar el pedido creado:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error al buscar el pedido creado"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": pedidoCreado})
}
