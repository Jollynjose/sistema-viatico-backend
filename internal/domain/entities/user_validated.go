package entities

type UserValidated struct {
	User
	isValidated bool
}

func (u *UserValidated) IsValidated() bool {
	return u.isValidated
}

func NewUserValidated(u *User) *UserValidated {
	if err := u.validate(); err != nil {
		return &UserValidated{
			User:        *u,
			isValidated: false,
		}
	}
	return &UserValidated{
		User:        *u,
		isValidated: true,
	}
}
