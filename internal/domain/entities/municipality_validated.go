package entities

type MunicipalityValidated struct {
	Municipality
	isValidated bool
}

func (m *MunicipalityValidated) IsValidated() bool {
	return m.isValidated
}

func NewMunicipalityValidated(m *Municipality) *MunicipalityValidated {
	if err := m.validate(); err != nil {
		return &MunicipalityValidated{
			Municipality: *m,
			isValidated:  false,
		}
	}
	return &MunicipalityValidated{
		Municipality: *m,
		isValidated:  true,
	}
}
