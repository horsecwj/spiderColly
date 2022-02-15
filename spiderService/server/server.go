package server

import (
	"Spider/spiderService/server/controller"
)

func Run(address string, release bool) error {

	return controller.Run(address, release)
}
