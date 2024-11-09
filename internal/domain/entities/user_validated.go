package entities

type UserValidated struct {
	User
	isValidated bool
	Err         error
}

func (u *UserValidated) IsValidated() bool {
	return u.isValidated
}

func NewUserValidated(u *User) *UserValidated {
	if err := u.validate(); err != nil {
		return &UserValidated{
			User:        *u,
			isValidated: false,
			Err:         err,
		}
	}
	return &UserValidated{
		User:        *u,
		isValidated: true,
	}
}
