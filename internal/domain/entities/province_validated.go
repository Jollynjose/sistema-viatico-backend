package entities

type ProvinceValidated struct {
	Province
	isValidated bool
}

func (p *ProvinceValidated) IsValidated() bool {
	return p.isValidated
}

func NewProvinceValidated(p *Province) *ProvinceValidated {
	if err := p.validate(); err != nil {
		return &ProvinceValidated{
			Province:    *p,
			isValidated: false,
		}
	}
	return &ProvinceValidated{
		Province:    *p,
		isValidated: true,
	}
}
