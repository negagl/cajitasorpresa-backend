package main

import (
	"cajitasorpresa/db"
	"cajitasorpresa/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Base de datos: conectar, migrar y sembrar (antes de aceptar peticiones)
	db.Conectar()
	db.Seed()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	api := router.Group("/api")

	// ---------- Públicas ----------

	// Boxes
	api.GET("/boxes", handlers.GetBoxes)
	api.GET("/boxes/:box_id", handlers.GetBoxPorId)
	api.GET("/boxes/:box_id/ingredientes", handlers.GetIngredientesPorBox)

	// Extras
	api.GET("/extras", handlers.GetExtras)
	api.GET("/extras/:extra_id", handlers.GetExtraByID)

	// Crear pedido
	api.POST("/pedidos", handlers.CrearPedido)

	// ---------- Admin ----------
	admin := api.Group("/admin")

	admin.GET("/lista-compras", handlers.GetListaCompras)

	// Pedidos
	admin.GET("/pedidos", handlers.GetPedidos)
	admin.GET("/pedidos/:pedido_id", handlers.GetPedidoPorId)
	admin.GET("/pedidos/:pedido_id/extras", handlers.GetPedidoExtraPorIdPedido)
	admin.GET("/pedidos/:pedido_id/personalizaciones", handlers.GetPersonalizacionPorIdPedido)
	admin.PATCH("/pedidos/:pedido_id/estado", handlers.ActualizarEstadoPedido)
	admin.GET("/boxes/:box_id/pedidos", handlers.GetPedidosPorBoxId)

	// Detalle por ID
	admin.GET("/pedidos_extra/:pedido_extra_id", handlers.GetPedidoExtraPorId)
	admin.GET("/personalizaciones/:personalizacion_id", handlers.GetPersonalizacionPorId)

	if err := router.Run(); err != nil {
		log.Fatal("No se pudo iniciar el servidor:", err)
	}
}
