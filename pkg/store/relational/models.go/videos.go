package models

import (
	"gorm.io/gorm"
)

type Videos struct {
	Model
	Genere string `json:"genere,omitempty"`
	VideoTitle    string `json:"video_title,omitempty"`
	Description   string `json:"description,omitempty"`
	Publishing    string `json:"publishing,omitempty"`
	ThumnailsURLs string `json:"thumbnails_urls,omitempty"`
}

func (d *Videos) Get(db *gorm.DB) error {

	return db.Where(d).First(d).Error
}

func (d *Videos) Create(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Create(d).Error
}

func (d *Videos) Delete(db *gorm.DB) error {
	return db.Where(&Videos{Model: Model{ID: d.Model.ID}}).Delete(d).Error
}

func (d *Videos) Update(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Where(&Videos{Model: Model{ID: d.Model.ID}}).Updates(d).Error
}

func (d *Videos) List(db *gorm.DB) ([]Videos, error) {

	Videoss := []Videos{}
	err := db.Where(d).Find(&Videoss).Error

	if err != nil {
		return nil, err
	}
	return Videoss, nil
}
