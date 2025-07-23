package server

import (
	"mystanford/logger"
	"strconv"
)

func Start(port int) {
	e := Server.Listen("0.0.0.0:" + strconv.Itoa(port))
	if e != nil {
		logger.Logger.Error(e.Error())
		return
	}
}
