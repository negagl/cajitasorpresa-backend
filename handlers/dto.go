package handlers

import "time"

type ExtraInput struct {
	ExtraID  uint `json:"extra_id" binding:"required"`
	Cantidad int  `json:"cantidad" binding:"required"`
}

type PersonalizacionInput struct {
	Categoria string `json:"categoria" binding:"required"`
	Eleccion  string `json:"eleccion" binding:"required"`
}

type CrearPedidoInput struct {
	NombreCliente   string                 `json:"nombre_cliente" binding:"required"`
	TelefonoCliente string                 `json:"telefono_cliente" binding:"required"`
	BoxID           uint                   `json:"box_id" binding:"required"`
	SaborJugo       string                 `json:"sabor_jugo"`
	ConAzucar       bool                   `json:"con_azucar"`
	MensajeTarjeta  string                 `json:"mensaje_tarjeta"`
	FotoURL         string                 `json:"foto_url"`
	FechaEntrega    time.Time              `json:"fecha_entrega" binding:"required"`
	Extras          []ExtraInput           `json:"extras"`
	Personalizacion []PersonalizacionInput `json:"personalizacion"`
}
