package entities

type FuelValidated struct {
	Fuel    Fuel
	isValid bool
	Err     error
}

func (f *FuelValidated) IsValidated() bool {
	return f.isValid
}

func NewFuelValidated(fuel *Fuel) *FuelValidated {
	if err := fuel.validate(); err != nil {
		return &FuelValidated{
			Fuel:    *fuel,
			isValid: false,
			Err:     err,
		}
	}

	return &FuelValidated{
		Fuel:    *fuel,
		isValid: true,
		Err:     nil,
	}
}
