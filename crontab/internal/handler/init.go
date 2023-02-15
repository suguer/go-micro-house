package handler

import (
	"context"

	"github.com/robfig/cron/v3"
)

type Crontab struct {
	Rule string
	Fun  func(c context.Context)
}

var (
	handlers = make([]Crontab, 0)
)

func InitHandle(c *cron.Cron, ctx context.Context) {
	for _, v := range handlers {
		c.AddFunc(v.Rule, func() {
			v.Fun(ctx)
		})
	}
}
