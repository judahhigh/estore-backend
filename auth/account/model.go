package account

type User struct {
	ID       string `json:"id,omitempty" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
