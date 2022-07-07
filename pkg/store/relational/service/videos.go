package service

import (
	"github.com/shikharvashistha/fampay/pkg/store/relational/models.go"
	"gorm.io/gorm"
)

type Videos interface {
	Create(model *models.Videos) error
	Delete(model *models.Videos) error
	Get(model *models.Videos) error
	Update(model *models.Videos) error
	List(model *models.Videos) ([]models.Videos, error)
}

type videos struct {
	db *gorm.DB
}

func NewVideosSvc(db *gorm.DB) Videos {
	return &videos{db}
}

func (dep *videos) Create(model *models.Videos) error {
	return model.Create(dep.db)
}

func (dep *videos) Delete(model *models.Videos) error {
	return model.Delete(dep.db)
}

func (dep *videos) Get(model *models.Videos) error {
	return model.Get(dep.db)
}

func (dep *videos) Update(model *models.Videos) error {
	return model.Update(dep.db)
}

func (dep *videos) List(model *models.Videos) ([]models.Videos, error) {
	return model.List(dep.db)
}
