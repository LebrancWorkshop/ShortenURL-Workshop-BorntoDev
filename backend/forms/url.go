package forms

import "gorm.io/gorm"

type ShortlyURL struct {
	gorm.Model
	OriginalURL 	string 	`gorm: "unique"`
	ShortURL 			string 	`gorm: "unique"`
}


