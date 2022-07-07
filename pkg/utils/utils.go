package utils

import (
	"log"

	"github.com/shikharvashistha/fampay/pkg/common"
)

func LoggerConsole(logDetail *common.Error) {
	log.Println(logDetail.ErrorCode + " : " + logDetail.ErrorDescription)
}
