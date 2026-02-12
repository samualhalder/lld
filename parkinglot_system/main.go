package main

import (
	"fmt"
	"time"

	"github.com/samualhalder/lld/parkinglot_system/parkinglot"
	"github.com/samualhalder/lld/parkinglot_system/payment"
	"github.com/samualhalder/lld/parkinglot_system/ticket"
	"github.com/samualhalder/lld/parkinglot_system/vehicle"
)

func main() {
	floor1 := &parkinglot.Floor{
		Id:      1,
		Space:   make(map[vehicle.VehicleType]int),
		Tickets: make(map[int]*ticket.Ticket),
	}
	floor1.AddSpace(vehicle.CarType, 0)
	floor1.AddSpace(vehicle.BikeType, 10)
	floor1.AddSpace(vehicle.TruckType, 1)
	floor2 := &parkinglot.Floor{
		Id:      2,
		Space:   make(map[vehicle.VehicleType]int),
		Tickets: make(map[int]*ticket.Ticket),
	}
	floor2.AddSpace(vehicle.CarType, 0)
	floor2.AddSpace(vehicle.BikeType, 10)
	floor2.AddSpace(vehicle.TruckType, 1)
	rateChart := &parkinglot.Rate{
		Rates: make(map[vehicle.VehicleType]int),
	}
	rateChart.AddRate(vehicle.BikeType, 1)
	rateChart.AddRate(vehicle.CarType, 10)
	rateChart.AddRate(vehicle.TruckType, 50)
	parkinglot := &parkinglot.ParkingLot{
		Floors: make(map[int]*parkinglot.Floor),
		Rate:   rateChart,
	}
	parkinglot.AddFloor(floor1)
	parkinglot.AddFloor(floor2)

	car, err := vehicle.GetVehicle(vehicle.CarType, 1)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	bike, err := vehicle.GetVehicle(vehicle.BikeType, 2)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	truck1, err := vehicle.GetVehicle(vehicle.TruckType, 3)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	truck2, err := vehicle.GetVehicle(vehicle.TruckType, 4)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	_, err = parkinglot.ParkVehicle(car)
	if err != nil {
		fmt.Println(err, car.GetId())
	}
	parkinglot.ParkVehicle(bike)
	parkinglot.ParkVehicle(truck1)
	time.Sleep(time.Second * 2)
	upi := &payment.Upi{}
	tkt, _ := parkinglot.UnParkVehicle(truck1, upi)
	fmt.Println("amoumt", tkt.Amount, "status", tkt.PaymentStatus)
	_, err = parkinglot.ParkVehicle(truck2)
	if err != nil {
		fmt.Println(err.Error())
	}

}
