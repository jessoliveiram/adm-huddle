package user

type User struct {
	Username string
	Name     string
	Email    string
	Role     string
}

type Role int

const (
	Admin Role = iota
	Medical
	Cleaning
	Maintenance
)
