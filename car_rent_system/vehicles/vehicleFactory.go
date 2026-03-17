package vehicles

func GetCar(id, baesPrice int, ty VehicleType) *Vehicle {
	return &Vehicle{
		Id:        id,
		BasePrice: baesPrice,
		Type:      ty,
		Status:    Available,
	}
}
