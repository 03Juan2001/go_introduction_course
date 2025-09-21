package structs

// Las structs en Go son colecciones de campos
// Se definen con la palabra clave "type" seguida del nombre de la struct y la palabra clave "struct"
// Los campos se definen con el nombre del campo seguido del tipo
// Se pueden anidar structs dentro de otras structs
// Se pueden agregar etiquetas a los campos para propósitos como la serialización JSON

type Product struct {
	ID    uint     `json:"id"` // Etiquetas para serialización JSON
	Name  string   `json:"name"`
	Type  Type     `json:"type"` // Campo de tipo struct anidada
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) AddTag(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

func (c *Car) AddProduct(product ...Product) {
	c.Products = append(c.Products, product...)
}

func (c Car) TotalPrice() float64 {
	var total float64
	for _, p := range c.Products {
		total += p.TotalPrice()
	}
	return total
}

func NewCar(userID uint) Car {
	return Car{
		UserID: userID,
	}
}
