package crontab

import (
	"github.com/robfig/cron/v3"
)

type Crontab struct {
	Rule string
	Fun  func()
}

var (
	handlers = make([]Crontab, 0)
)

func InitCrontab() {
	c := cron.New()
	for _, v := range handlers {
		c.AddFunc(v.Rule, v.Fun)
	}
	c.Start()
}
