package models

import (
	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/forms"
	"gorm.io/gorm"
)

type Model struct {
	db *gorm.DB
}

func NewModel(db *gorm.DB) *Model {
	return &Model{db: db}
}

func (m *Model) CreateShortURL(url *forms.ShortlyURL) (*gorm.DB, error) {
	result := m.db.Create(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (m *Model) GetOriginalURL(url *forms.ShortlyURL) (*gorm.DB, error) {
	result := m.db.Where("original_url = ?", url.OriginalURL).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (m *Model) GetShortURL(url *forms.ShortlyURL, shortURL string) (*gorm.DB, error) {
	result := m.db.Where("short_url = ?", shortURL).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}
