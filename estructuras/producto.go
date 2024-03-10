package estructuras

type Producto struct {
	Nombre     string  `json:"nombre"`
	Definicion string  `json:"definicion"`
	Precio     float32 `json:"precio"`
	Valoracion int     `json:"valoracion"`
	Version    Version `json:"version"`
}

func NewProducto(nombre, definicion string, precio float32, valoracion int) *Producto {
	return &Producto{nombre, definicion, precio, valoracion, *NewVersion()}
}

func (p *Producto) GetNombre() string {
	return p.Nombre
}

func (p *Producto) GetDefinicion() string {
	return p.Definicion
}

func (p *Producto) GetPrecio() float32 {
	return p.Precio
}

func (p *Producto) GetValoracion() int {
	return p.Valoracion
}

func (p *Producto) GetTimestamp() string {
	return p.Version.TimeStamp
}

func (p *Producto) GetVersion() float32 {
	return p.Version.GetVersion()
}

func (p *Producto) SetNombre(nombre string) {
	p.Nombre = nombre
}

func (p *Producto) SetDefinicion(definicion string) {
	p.Definicion = definicion
}

func (p *Producto) SetPrecio(precio float32) {
	p.Precio = precio
}

func (p *Producto) SetValoracion(valoracion int) {
	p.Valoracion = valoracion
}
