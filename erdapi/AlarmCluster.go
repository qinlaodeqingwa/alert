package erdapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type AlertRule struct {
	Success bool      `json:"success"`
	Data    AlertData `json:"data"`
}
type TriggerCondition struct {
	Condition string `json:"condition"`
	Operator  string `json:"operator"`
	Values    string `json:"values"`
}
type AlertData struct {
	TriggerCondition []TriggerCondition `json:"triggerCondition"`
}

func GetAlarmClusterName(alarmgroupId int) (string, error) {
	alarmID := strconv.Itoa(alarmgroupId)
	//获取clusterName
	url := Url("/api/orgCenter/alerts/alarmId", nil, alarmID)
	fmt.Println("", url)

	accessToken, _ := GetAccessToken("/api/OrgCenter/alerts/<id>", "GET")
	body, _ := DoRequest(Request{
		Method: "GET",
		URL:    url,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken},
	})

	alarm := &AlertRule{}
	if err := json.Unmarshal(body, alarm); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return "", err
	}
	alarmcluster := alarm.Data.TriggerCondition[0].Values
	return alarmcluster, nil
}
