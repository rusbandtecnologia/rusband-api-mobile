package models

type Client struct {
	ID       uint   `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (Client) TableName() string {
	return "clients"
}
