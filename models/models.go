package models

import "time"

type TipoExtra string
type EstadoPedido string

const (
	Complemento TipoExtra = "complemento"
	Adicional   TipoExtra = "adicional"

	Pendiente EstadoPedido = "pendiente"
	EnProceso EstadoPedido = "en proceso"
	Entregado EstadoPedido = "entregado"
)

type Box struct {
	ID              uint             `json:"id"`
	Nombre          string           `json:"nombre"`
	Descripcion     string           `json:"descripcion"`
	PrecioBase      float64          `json:"precio_base"`
	Imagen          string           `json:"imagen"`
	IngredientesBox []IngredienteBox `json:"ingredientes"`
}

type IngredienteBox struct {
	ID             uint    `json:"id"`
	Nombre         string  `json:"nombre"`
	BoxID          uint    `json:"box_id"`
	Cantidad       float64 `json:"cantidad"`
	UnidadDeMedida string  `json:"unidad_de_medida"`
}

type Extra struct {
	ID          uint      `json:"id"`
	Nombre      string    `json:"nombre"`
	Tipo        TipoExtra `json:"tipo"`
	Descripcion string    `json:"descripcion"`
	PrecioBase  float64   `json:"precio_base"`
}

type Pedido struct {
	ID              uint                    `json:"id"`
	NombreCliente   string                  `json:"nombre_cliente"`
	TelefonoCliente string                  `json:"telefono_cliente"`
	BoxID           uint                    `json:"box_id"`
	SaborJugo       string                  `json:"sabor_jugo"`
	ConAzucar       bool                    `json:"con_azucar"`
	MensajeTarjeta  string                  `json:"mensaje_tarjeta"`
	FotoURL         string                  `json:"foto_url"`
	ItemsExtra      []PedidoExtra           `json:"items_extra"`
	Personalizacion []PersonalizacionPedido `json:"personalizacion"`
	PrecioTotal     float64                 `json:"precio_total"`
	FechaPedido     time.Time               `json:"fecha_pedido"`
	FechaEntrega    time.Time               `json:"fecha_entrega"`
	Estado          EstadoPedido            `json:"estado"`
}

type PedidoExtra struct {
	ID       uint `json:"id"`
	PedidoID uint `json:"pedido_id"`
	ExtraID  uint `json:"extra_id"`
	Cantidad int  `json:"cantidad"`
}

type PersonalizacionPedido struct {
	ID        uint   `json:"id"`
	PedidoID  uint   `json:"pedido_id"`
	Categoria string `json:"categoria"`
	Eleccion  string `json:"eleccion"`
}
