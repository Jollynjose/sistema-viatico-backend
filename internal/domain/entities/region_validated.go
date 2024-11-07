package entities

type RegionValidated struct {
	Region
	isValidated bool
}

func (r *RegionValidated) IsValidated() bool {
	return r.isValidated
}

func NewRegionValidated(r *Region) *RegionValidated {
	if err := r.validate(); err != nil {
		return &RegionValidated{
			Region:      *r,
			isValidated: false,
		}
	}
	return &RegionValidated{
		Region:      *r,
		isValidated: true,
	}
}
