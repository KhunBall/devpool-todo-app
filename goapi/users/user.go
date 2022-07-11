package users

type User struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Todo string `json:"todo"`
}
