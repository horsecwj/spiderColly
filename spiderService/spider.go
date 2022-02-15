package spiderService

import (
	"Spider/spiderService/schedule"
)

func Run() error {
	return schedule.Run()
}
