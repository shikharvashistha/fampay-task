package relational

import (
	"github.com/shikharvashistha/fampay/pkg/store/relational/service"
	"gorm.io/gorm"
)

func NewRelational(db *gorm.DB) RL {
	return &relational{
		videos: service.NewVideosSvc(db),
		cron:   service.NewCronSvc(db),
	}
}

type relational struct {
	videos service.Videos
	cron   service.Cron
}

func (r *relational) Videos() service.Videos {
	return r.videos
}

func (r *relational) Cron() service.Cron {
	return r.cron
}
