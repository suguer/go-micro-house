package handler

import (
	"context"
	"crontab/internal/model"
	"crontab/internal/service"
	"crontab/routes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func init() {
	handlers = append(handlers, Crontab{
		Rule: "* * * * *",
		Fun:  RecondCreateAuto,
	})

}

func RecondCreateAuto(c context.Context) {
	ser := routes.RData.GetService("sms").(service.SmsServiceClient)
	fmt.Printf("[%s] RecondCreateAuto\n", time.Now().Format("2006-01-02 15:04:05"))
	var list []model.SmsConfig
	model.DB.Model(&model.SmsConfig{}).
		Where("status =?", "active").
		Where("available > ?", 0).
		Find(&list)

	fmt.Printf("list: %v\n", len(list))
	for _, sc := range list {
		Day := strings.Split(sc.Day, ",")
		var StartAt []string
		for _, v := range Day {
			day, _ := strconv.Atoi(v)
			StartAt = append(StartAt, time.Now().AddDate(0, 0, day).Format("2006-01-02"))
		}
		fmt.Printf("Day: %v\n", Day)
		fmt.Printf("StartAt: %v\n", StartAt)
		var record_list []model.Record
		model.DB.Model(&model.Record{}).
			Where("user_id =?", sc.UserId).
			Where("start_at in ?", StartAt).
			Where("status =?", "normal").
			Where("type =?", "rent").
			Find(&record_list)
		fmt.Printf("len(record_list): %v\n", len(record_list))
		for _, r := range record_list {
			fmt.Printf("r.ID: %v\n", r.ID)
			var req service.SmsCreateRequest
			req.Content = "asdasd"
			req.UserId = 1
			resp, err := ser.Create(context.Background(), &req)
			fmt.Printf("err: %v\n", err)
			if err != nil {
				continue
			}
			fmt.Printf("resp.Code: %v\n", resp.Code)

		}
	}
}
