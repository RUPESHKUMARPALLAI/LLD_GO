package vehicle

type VehicleType string

const (
	Car VehicleType = "Car"
	Bike VehicleType = "Bike"
	Truck VehicleType = "Truck"
)

type Vehicle struct {
	Number string
	Type VehicleType
}

func NewVehicle (Number string, Type VehicleType) *Vehicle {
	return &Vehicle{
		Number: Number,
		Type: Type,
	}
}