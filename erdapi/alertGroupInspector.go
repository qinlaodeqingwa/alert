package erdapi

import (
	"encoding/json"
	"fmt"
	"os"
)

type Alert struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AlertResponse struct {
	Data struct {
		List []Alert `json:"list"`
	} `json:"data"`
}

//type AlarmInfo struct {
//	Name string
//	ID   string
//}

func CreateAlarm() {
	alarmInfos := GetAlarmID()
	//if err != nil {
	//	fmt.Println("Error getting alarm IDs: ", err)
	//	return
	//}
	for _, alarmInfo := range alarmInfos {
		fmt.Printf("Alarm name: %s, ID: %s\n", alarmInfo.Name, alarmInfo.ID)
	}

}

func GetAlarmID() []Alert {
	body, _ := RetrieveAlert()
	data := &AlertResponse{}
	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil
	}
	names := map[string]bool{
		"Erda-L1-prod(勿删)":   true,
		"Erda-L2(勿删)":        true,
		"Erda-L2-noprod(勿删)": true,
	}

	var alarms []Alert
	for _, item := range data.Data.List {
		if _, ok := names[item.Name]; ok {
			alarms = append(alarms, Alert{ID: item.ID, Name: item.Name})
		}
	}
	fmt.Println("the alarm id is", alarms)
	return alarms
}

func RetrieveAlert() ([]byte, error) {
	accessToken, err := GetAccessToken("/api/orgCenter/alerts", "GET")
	if err != nil {
		fmt.Println("Error on GetAccessToken:", err)
	}
	os.Setenv("Token", accessToken)

	AlarmUrl := "https://dice.erda.cloud/api/hyjtsc/orgCenter/alerts?pageNo=1&pageSize=10"
	//AlarmUrl := Url("/api/orgCenter/alerts", url.Values{
	//	"pageNo":   {"1"},
	//	"pageSize": {"10"},
	//}, "")

	respBody, _ := DoRequest(Request{
		Method: "GET",
		URL:    AlarmUrl,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken},
	})

	return respBody, nil
}
