package relational

import "github.com/shikharvashistha/fampay/pkg/store/relational/service"

type RL interface {
	Videos() service.Videos
	Cron() service.Cron
}
