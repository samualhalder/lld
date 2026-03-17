package vehicles

type VehicleType int

const (
	SUV VehicleType = iota
	Compact
	Van
	Siddan
)

type VehicleStatus int

const (
	Available VehicleStatus = iota
	Booked
	Service
)

type Vehicle struct {
	Id        int
	Type      VehicleType
	Status    VehicleStatus
	VIN       string
	BasePrice int
}
