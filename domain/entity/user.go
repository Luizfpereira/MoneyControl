package entity

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser() (*User, error) {
	return nil, nil
}

func (u *User) Validate() error {
	return nil
}
