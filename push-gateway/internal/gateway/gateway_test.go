package gateway

import (
	"fmt"
	"testing"
)

func TestNowcnGateway(t *testing.T) {

	// str := `{"msg":"SUCCESS","code":"0"}`
	// fmt.Printf("str: %v\n", str)
	// var data map[string]interface{}
	// json.Unmarshal([]byte(str), &data)
	// fmt.Printf("data: %v\n", data)
	// code := data["code"].(string)
	// fmt.Printf("code: %v\n", code)
	// fmt.Printf("code: %T\n", data["code"])
	gateway := &NowcnGateway{}

	err := gateway.SendMessage("莫明豪,优品未来居509B将于02月02日到期欠租545元,如有异议请联系17336172723", "15088132466")
	fmt.Printf("err: %v\n", err)
}
