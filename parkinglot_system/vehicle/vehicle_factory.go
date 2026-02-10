package vehicle

import "fmt"

func GetVehicle(vehicle VehicleType, id int) (Vehicle, error) {
	switch vehicle {
	case CarType:
		return &Car{
			Type: CarType,
			Id:   id,
		}, nil
	case TruckType:
		return &Truck{
			Type: TruckType,
			Id:   id,
		}, nil
	case BikeType:
		return &Bike{
			Type: BikeType,
			Id:   id,
		}, nil
	default:

		return nil, fmt.Errorf("unknown vehicle type")
	}
}
