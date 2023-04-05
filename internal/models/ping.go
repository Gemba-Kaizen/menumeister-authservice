package models

type Ping struct {
	UserId      int64  `json:"id" gorm:"primaryKey"`
	Message string `json:"message"`
}
