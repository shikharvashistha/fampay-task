package service

import (
	"github.com/shikharvashistha/fampay/pkg/store/relational/models.go"
	"gorm.io/gorm"
)

type Cron interface {
	Create(model *models.Cron) error
	Delete(model *models.Cron) error
	Get(model *models.Cron) error
	Update(model *models.Cron) error
	List(model *models.Cron) ([]models.Cron, error)
}
type cron struct {
	db *gorm.DB
}

func NewCronSvc(db *gorm.DB) Cron {
	return &cron{db}
}

func (conf *cron) Create(model *models.Cron) error {
	return model.Create(conf.db)
}

func (c *cron) Delete(model *models.Cron) error {
	return model.Delete(c.db)
}

func (c *cron) Get(model *models.Cron) error {
	return model.Get(c.db)
}

func (c *cron) Update(model *models.Cron) error {
	return model.Update(c.db)
}

func (c *cron) List(model *models.Cron) ([]models.Cron, error) {
	return model.List(c.db)
}
