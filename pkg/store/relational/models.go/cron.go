package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Cron struct {
	CronModel
}

type CronModel struct {
	Model
	CronID string `gorm:"size:191"`
}

func (c *Cron) Get(db *gorm.DB) error {

	return db.Preload(clause.Associations).Where(c).First(c).Error
}

func (c *Cron) Create(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Create(c).Error
}

func (c *Cron) Delete(db *gorm.DB) error {
	return db.Where(&Cron{CronModel: CronModel{CronID: c.CronID}}).Delete(c).Error
}

func (c *Cron) Update(db *gorm.DB) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Where(&Cron{CronModel: CronModel{Model: Model{ID: c.Model.ID}}}).Updates(c).Error
}

func (c *Cron) List(db *gorm.DB) ([]Cron, error) {

	Cron := []Cron{}
	err := db.Where(c).Find(&Cron).Error

	if err != nil {
		return nil, err
	}
	return Cron, nil
}
