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
		// Rule: "* * * * *",
		Rule: "30 10 * * *",
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
		var record_list []model.Record
		model.DB.Model(&model.Record{}).
			Where("user_id =?", sc.UserId).
			Where("start_at in ?", StartAt).
			Where("status =?", "normal").
			Where("type =?", "rent").
			Find(&record_list)
		for _, r := range record_list {
			var contract model.Contract
			var house model.House
			model.DB.First(&contract, r.ContractID)
			model.DB.First(&house, r.HouseID)
			if *contract.TargetPhone == "" {
				continue
			}
			var req service.SmsCreateRequest
			req.Content = fmt.Sprintf("%v,%v于%v到期,租金%v元,请联系并缴纳租金,如已缴纳请忽略",
				*contract.TargetName,
				*house.Title,
				r.StartAt.Format("01月02日"),
				r.Money)
			req.UserId = uint32(r.UserID)
			req.HouseId = uint32(r.HouseID)
			req.ContractId = uint32(r.ContractID)
			req.RecordId = uint32(r.ID)
			req.Mobile = *contract.TargetPhone
			ser.Create(context.Background(), &req)
		}
	}
}
