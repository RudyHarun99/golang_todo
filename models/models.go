package models

type Todos struct {
	ID 					uint		`gorm:"primary_key; auto_increment" json:"id"`
	Title				string	`gorm:"not null" json:"title"`
	Description	string	`json:"description"`
	Status			string	`json:"status"`
}