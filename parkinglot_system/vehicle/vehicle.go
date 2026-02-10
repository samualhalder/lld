package vehicle

type VehicleType string

const (
	CarType   VehicleType = "car"
	TruckType VehicleType = "truck"
	BikeType  VehicleType = "bike"
)

type Vehicle interface {
	IsType() VehicleType
	GetId() int
}

type Car struct {
	Type VehicleType
	Id   int
}

func (c *Car) IsType() VehicleType {
	return c.Type
}
func (c *Car) GetId() int {
	return c.Id
}

type Bike struct {
	Type VehicleType
	Id   int
}

func (b *Bike) IsType() VehicleType {
	return b.Type
}
func (c *Bike) GetId() int {
	return c.Id
}

type Truck struct {
	Type VehicleType
	Id   int
}

func (t *Truck) IsType() VehicleType {
	return t.Type
}
func (c *Truck) GetId() int {
	return c.Id
}
