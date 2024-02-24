package user

type User struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Username  string `db:"user_name" json:"username"`
	Password  string `db:"password" json:"password"`
	IsAdmin   bool   `db:"is_admin" json:"isAdmin"`
}

type UserResponse struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Username  string `db:"user_name" json:"username"`
}
