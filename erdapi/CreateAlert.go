package erdapi

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateAlertGroupsL1() (int, error) {
	alertGroupUrl := Url("/api/notify-groups", nil, "")

	args := os.Args
	if len(args) != 2 {
		fmt.Print("用法：用户id")
	}
	data := map[string]interface{}{
		"scopeType": "org",
		"scopeId":   os.Getenv("ORGID"),
		"name":      "Erda-L1(勿删)",
		"targets": []map[string]interface{}{
			{
				"type": "dingding",
				"values": []map[string]interface{}{
					{
						"receiver": "https://kubeprober.erda.cloud/robot/send?access_token=a841f19df03cbb7d480627d2329d3ea7ff33516b7d8314c96cae571001b18cf8",
						"secret":   "SEC494e5597ad849f36c87be531d2472a0be3eb1fa25ae9dc2d4e343f82c933a0ad",
					},
				},
			},
		},
	}
	respBody, _ := DoRequest(Request{
		Method: "POST",
		URL:    alertGroupUrl,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + os.Getenv("Token")},
		Body:   data,
	})
	var response map[string]interface{}

	if err := json.Unmarshal([]byte(respBody), &response); err != nil {
		return 0, fmt.Errorf("解析 JSON 出错: %v", err)
	}

	dataValueL1, ok := response["data"].(int)
	if !ok {
		return 0, fmt.Errorf("无法转换为 float64 类型")
	}

	return dataValueL1, nil
}
func CreateAlertGroupsL2() (int, error) {

	alertGroupUrl := Url("/api/notify-groups", nil, "")

	args := os.Args
	if len(args) != 2 {
		fmt.Print("用法：用户id")
	}
	//L1Ding:=&AlertDetail{}
	data := map[string]interface{}{
		"scopeType": "org",
		"scopeId":   os.Getenv("ORGID"),
		"name":      "Erda-L2(勿删)",
		"targets": []map[string]interface{}{
			{
				"type": "dingding",
				"values": []map[string]interface{}{
					{
						"receiver": "https://kubeprober.erda.cloud/robot/send?access_token=a841f19df03cbb7d480627d2329d3ea7ff33516b7d8314c96cae571001b18cf8",
						"secret":   "SEC494e5597ad849f36c87be531d2472a0be3eb1fa25ae9dc2d4e343f82c933a0ad",
					},
				},
			},
		},
	}
	respBody, _ := DoRequest(Request{
		Method: "POST",
		URL:    alertGroupUrl,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + os.Getenv("Token")},
		Body:   data,
	})

	var response map[string]interface{}
	if err := json.Unmarshal([]byte(respBody), &response); err != nil {
		return 0, fmt.Errorf("解析 JSON 出错: %v", err)
	}

	dataValueL2, ok := response["data"].(int)
	if !ok {
		return 0, fmt.Errorf("无法转换为 float64 类型")
	}

	return dataValueL2, nil
}
