package auth

type Auth struct {
	ID       int
	Email    string `gorm:"unique"`
	Password string
	Name     string
}