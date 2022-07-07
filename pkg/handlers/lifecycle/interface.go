package lifecycle

import (
	"github.com/shikharvashistha/fampay/pkg/types"
)

type Deployment interface {
	GetData(string, string, string) (types.Response, error)
	SearchData(title string, description string) (types.Response, error)
	CronSchedule(types.CronRequest) (types.CronResponse, error)
	CronDelete(types.CronDeleteRequest) error
}
