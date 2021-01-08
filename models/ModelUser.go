package models

// User type
type User struct {
	ID    uint   `gorm:"primaryKey;column:id_user" json:"id"`
	Email string `json:"email"`
}

// TableName - Convertendo o nome da tabela
func (user User) TableName() string {
	return "tb_user"
}
