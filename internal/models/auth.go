package models

type Merchant struct {
	Id       int64  `json:"id gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
