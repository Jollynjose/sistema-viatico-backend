package entities

type FuelHistoryValidated struct {
	FuelHistory FuelHistory
	isValidated bool
	Err         error
}

func (f *FuelHistoryValidated) IsValidated() bool {
	return f.isValidated
}

func NewFuelHistoryValidated(f *FuelHistory) *FuelHistoryValidated {
	if err := f.validate(); err != nil {
		return &FuelHistoryValidated{
			FuelHistory: *f,
			isValidated: false,
			Err:         err,
		}
	}
	return &FuelHistoryValidated{
		FuelHistory: *f,
		isValidated: true,
		Err:         nil,
	}
}
