package erdapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
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

func GetAlarmCluster(alarmgroupname string) (string, error) {
	var alarmgroupId int
	u := Url("/api/orgCenter/alerts", url.Values{
		"pageNo":   {"1"},
		"pageSize": {"10"},
	}, "")

	body, _ := DoRequest(Request{
		Method: "GET",
		URL:    u,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + os.Getenv("Token")},
	})

	alarmInfo := &Response{}
	if err := json.Unmarshal(body, alarmInfo); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return "", err
	}
	for _, alertid := range alarmInfo.Data.List {
		if alertid.Name == alarmgroupname {
			alarmgroupId = alertid.ID
		}
	}
	alarmID := strconv.Itoa(alarmgroupId)

	url := Url("/api/orgCenter/alerts/alarmId", nil, alarmID)
	body, _ = DoRequest(Request{
		Method: "GET",
		URL:    url,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + os.Getenv("Token")},
	})
	alarm := &AlertRule{}
	if err := json.Unmarshal(body, alarm); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return "", err
	}
	alarmcluster := alarm.Data.TriggerCondition[0].Values
	return alarmcluster, nil
}
