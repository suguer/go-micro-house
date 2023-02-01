package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"push-gateway/pkg/http"

	"github.com/spf13/viper"
)

type NowcnGateway struct {
	httpClient http.HttpClient
}

func (g *NowcnGateway) SendMessage(content, mobile string) error {
	request := http.NewHttpRequest()
	request.SetQuery("apiType", "3")
	request.SetQuery("mobile", mobile)
	request.SetQuery("content", content)
	g.buildParam(request)
	response, _ := g.send(request)

	var data map[string]interface{}
	json.Unmarshal(response.GetBody(), &data)
	fmt.Printf("data: %v\n", data)
	code := data["code"].(string)
	if code == "0" {
		return nil
	}
	return errors.New(data["msg"].(string))
}

func (g *NowcnGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("userId", viper.GetString("sms.appid"))
	request.SetQuery("password", viper.GetString("sms.secret"))
}
func (g *NowcnGateway) send(request *http.HttpRequest, Action ...string) (*http.HttpResponse, error) {
	request.SetURL("http://ad1200.now.net.cn:2003/sms/sendSMS" + request.GetPath())
	response, err := g.httpClient.Send(request)
	return response, err
}
