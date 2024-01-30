package erdapi

import (
	"encoding/json"
	"fmt"
	"os"
)

type notifybody struct {
	Success bool `json:"success"`
	Data    int  `json:"data"`
}

func CreateNotifyGroupsL1() (int, error) {
	notifyGroupUrl := "http://erda-server.default.svc.cluster.local:9529/api/notify-groups"
	//notifyGroupUrl := Url("/api/notify-groups", nil, "")

	//args := os.Args
	//if len(args) != 2 {
	//	fmt.Print("用法：用户id")
	//}
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
		URL:    notifyGroupUrl,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + os.Getenv("Token")},
		Body:   data,
	})

	var response notifybody
	if err := json.Unmarshal(respBody, &response); err != nil {
		return 0, fmt.Errorf("解析 JSON 出错: %v", err)
	}
	return response.Data, nil
}
func CreateNotifyGroupsL2() (int, error) {
	fmt.Println("开始创建告警通知组咯")
	notifyGroupUrl := "http://erda-server.default.svc.cluster.local:9529/api/notify-groups"
	//notifyGroupUrl := Url("/api/notify-groups", nil, "")
	//L1Ding:=&AlertDetail{}
	data := map[string]interface{}{
		"scopeType": "org",
		"scopeId":   "100812",
		"name":      "Erda-L2(勿删)",
		"targets": []map[string]interface{}{
			{
				"type": "dingding",
				"values": []map[string]interface{}{
					{
						"receiver": "https://kubeprober.erda.cloud/robot/send?access_token=44861bb01a06f1d676410a5d66f3c7ca98bfe4a5a783e69cd0c4a83cba26c9df",
						"secret":   "SEC7d7ab227054771565344a3727bc71798f7d98cda86b991c255cea347536cc47d",
					},
				},
			},
		},
	}
	accessToken, _ := GetAccessToken("/api/notify-groups", "POST")
	respBody, _ := DoRequest(Request{
		Method: "POST",
		URL:    notifyGroupUrl,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken},
		Body:   data,
	})

	var response notifybody
	if err := json.Unmarshal(respBody, &response); err != nil {
		return 0, fmt.Errorf("解析 JSON 出错: %v", err)
	}

	return response.Data, nil
}
