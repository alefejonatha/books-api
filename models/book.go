package models

import (
	"time"

	"gorm.io/gorm"
)

//GORM definiu uma gorm.Model estrutura, que inclui campos ID, CreatedAt, UpdatedAt,DeletedAt
type Book struct {
	ID          uint           `json:"id" gorm:"primarykey"` //Upper case convenções do GORM
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MediumPrice float64        `json:"medium_price"`
	Author      string         `json:"author"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`              // Configura a hora atual se for zero na criação
	UpdatedAt   time.Time      `json:"updated"`              // Definido para os segundos atuais do unix na atualização ou se for zero na criação
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"` //soft delete or fake delete(pesquisar)
}
