package entities

type FuelHistoryValidated struct {
	FuelHistory FuelHistory
	isValidated bool
}

func (f *FuelHistoryValidated) IsValidated() bool {
	return f.isValidated
}

func NewhistoryValidated(f *FuelHistory) *FuelHistoryValidated {
	if err := f.validate(); err != nil {
		return &FuelHistoryValidated{
			FuelHistory: *f,
			isValidated: false,
		}
	}
	return &FuelHistoryValidated{
		FuelHistory: *f,
		isValidated: true,
	}
}
