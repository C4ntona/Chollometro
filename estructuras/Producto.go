package estructuras

type Producto struct {
	nombre     string
	definicion string
	precio     float32
	valoracion int
	version    Version
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

func (p *Producto) getVersion() int {
	return p.version.getVersion()
}

func (p *Producto) setNombre() string {
	return p.nombre
}

func (p *Producto) setDefinicion() string {
	return p.nombre
}

func (p *Producto) setPrecio() float32 {
	return p.precio
}

func (p *Producto) setValoracion() int {
	return p.valoracion
}
