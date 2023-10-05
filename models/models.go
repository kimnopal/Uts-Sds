package models

type User struct {
	Id_user               uint `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Nama, Email, Username string
	Password              string
}

// Tulis jawab code nya di sini
type Admin struct {
	Id       uint64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
