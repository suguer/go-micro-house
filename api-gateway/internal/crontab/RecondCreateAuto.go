package crontab

import (
	"fmt"
	"time"
	// "github.com/gin-gonic/gin"
)

func init() {
	// handlers = append(handlers, Crontab{
	// 	Rule: "* * * * *",
	// 	Fun:  RecondCreateAuto,
	// })

}

func RecondCreateAuto() {
	fmt.Printf("[%s] RecondCreateAuto\n", time.Now().Format("2006-01-02 15:04:05"))

}
