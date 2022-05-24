package database

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
}
