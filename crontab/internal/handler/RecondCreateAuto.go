package handler

import (
	"context"
	"crontab/routes"
	"fmt"
	"time"
)

func init() {
	handlers = append(handlers, Crontab{
		Rule: "* * * * *",
		Fun:  RecondCreateAuto,
	})

}

func RecondCreateAuto(c context.Context, router routes.Router) {
	fmt.Printf("[%s] RecondCreateAuto\n", time.Now().Format("2006-01-02 15:04:05"))

}
