package store

import (
	"github.com/shikharvashistha/fampay/pkg/store/relational"
	"gorm.io/gorm"
)

type Store interface {
	RL() relational.RL
}

func NewStore(db *gorm.DB) Store {
	return &store{
		rl: relational.NewRelational(db),
	}
}
