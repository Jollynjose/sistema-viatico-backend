package entities

type RouteValidated struct {
	Route
	IsValid bool
	Err     error
}

func NewRouteValidated(route Route) *RouteValidated {
	if err := route.validate(); err != nil {
		return &RouteValidated{
			Route:   route,
			IsValid: false,
			Err:     err,
		}
	}

	return &RouteValidated{
		Route:   route,
		IsValid: true,
		Err:     nil,
	}
}
