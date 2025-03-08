package domain

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	
}

func NewCreateUser(name string, email string) *User {
	return &User{Name: name, Email: email}
}

func (user *User) GetName() string {
	return user.Name
}

func (user *User) SetName(name string) {
	user.Name = name
}
func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) SetEmail(email string) {
	user.Email = email
}
