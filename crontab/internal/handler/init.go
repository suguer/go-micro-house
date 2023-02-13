package handler

import (
	"context"
	"crontab/routes"

	"github.com/robfig/cron/v3"
)

type Crontab struct {
	Rule string
	Fun  func(c context.Context, router routes.Router)
}

var (
	handlers = make([]Crontab, 0)
)

func InitHandle(c *cron.Cron, ctx context.Context, router routes.Router) {
	for _, v := range handlers {
		c.AddFunc(v.Rule, func() {
			v.Fun(ctx, router)
		})
	}
}
