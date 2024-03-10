package estructuras

type Producto struct {
	nombre     string
	definicion string
	precio     float32
	valoracion int
	version    Version
}

func NewProducto(nombre, definicion string, precio float32, valoracion int) *Producto {
	return &Producto{nombre, definicion, precio, valoracion, *NewVersion()}
}

func (p *Producto) getNombre() string {
	return p.nombre
}

func (p *Producto) getDefinicion() string {
	return p.nombre
}

func (p *Producto) getPrecio() float32 {
	return p.precio
}

func (p *Producto) getValoracion() int {
	return p.valoracion
}

func (p *Producto) getTimestamp() string {
	return p.version.timeStamp
}

func (p *Producto) getVersion() float32 {
	return p.version.getVersion()
}

func (p *Producto) setNombre(nombre string) {
	p.nombre = nombre
}

func (p *Producto) setDefinicion(definicion string) {
	p.definicion = definicion
}

func (p *Producto) setPrecio(precio float32) {
	p.precio = precio
}

func (p *Producto) setValoracion(valoracion int) {
	p.valoracion = valoracion
}
