package erdapi

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	Group1 = "Erda-L1-prod(勿删)"
	Group2 = "Erda-L2-noprod(勿删)"
	Group3 = "Erda-L2(勿删)"
)

var nameToGroupID = map[string]*int{
	"Erda-L1-prod(勿删)":   &GroupIds.GroupL1,
	"Erda-L2(勿删)":        &GroupIds.GroupL2,
	"Erda-L2-noprod(勿删)": &GroupIds.GroupL2,
}

type GetClustersFunc func() []OrgInfo
type GetAlarmInfo func() []Alert

func HandleClusterAndAlertGroups(getClusters GetClustersFunc, Alarm GetAlarmInfo) {
	clusters := getClusters()
	alertGroups := Alarm()
	requiredAlertGroups := []string{Group1, Group2, Group3}

	switch len(clusters) {
	case 1:
		if len(alertGroups) < 2 {
			fmt.Println("未满足要求，需要创建缺失的告警组 ")
			for _, group := range requiredAlertGroups {
				if !Contains(alertGroups, group) {
					fmt.Println("1需要创建的告警组为 ", group)
					err := CreateAlarmGroup(group)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		} else {
			fmt.Println("告警组数量满足要求，无需创建，开始更新告警项 ")
		}
	case 2:
		fallthrough
	default:
		if len(alertGroups) < 3 {
			for _, group := range requiredAlertGroups {
				if !Contains(alertGroups, group) {
					fmt.Println("未满足要求，需要创建缺失的告警组 ", group)
					err := CreateAlarmGroup(group)
					if err != nil {
						fmt.Println("创建失败", err)
					}
				}
			}
		} else {
			fmt.Println("告警组数量满足要求，无需创建 可以选择更新了 ")
			UpdateAlarm(len(alertGroups), alertGroups)
		}
	}
}

// Contains 判断切片中是否包含某个字符串的函数
func Contains(slice []Alert, item string) bool {
	//set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		if s.Name == item {
			return true
		}
	}
	return false
}

func CreateAlarmGroup(name string) error {
	var notifyId int
	if groupID, ok := nameToGroupID[name]; ok {
		notifyId = *groupID
		log.Println(notifyId)
	}

	accessToken, _ := GetAccessToken("/api/orgCenter/alerts", "POST")
	alertGroupUrl := Url("/api/orgCenter/alerts", nil, "")
	template, _ := getTemplate("hyjtsc-prod", name, notifyId)

	if template != nil {
		_, err := DoRequest(Request{
			Method: "POST",
			URL:    alertGroupUrl,
			Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken},
			Body:   template,
		})
		if err != nil {
			fmt.Println("执行请求时出错: %w", err)
		}
	} else {
		log.Println("找不到与以下内容匹配的模板: ", name)
		return fmt.Errorf("找不到与 %v 相匹配的模板", name)
	}
	return nil
}

func getTemplate(orgname, targetname string, notifyid int) (map[string]interface{}, error) {
	alertItemL1, alertItemL2NoProd, alertItemL2, err := ProcessTemplateAndData(orgname, targetname, notifyid)
	if err != nil {
		return nil, fmt.Errorf("处理模板和数据时出错: %w", err)
	}
	switch {
	case strings.Contains(targetname, "L1"):
		return alertItemL1, nil
	case strings.Contains(targetname, "L2-noprod"):
		return alertItemL2NoProd, nil
	case strings.Contains(targetname, "L2"):
		return alertItemL2, nil
	default:
		return nil, nil
	}
}

func UpdateAlarm(num int, alarm []Alert) {
	fmt.Println("更新操作")
	var alartname []string
	var notifyid int

	if num == 1 {
		alartname = []string{Group1, Group2}
	} else if num > 1 {
		alartname = []string{Group1, Group2, Group3}
	}
	for _, targetName := range alartname {
		for _, alert := range alarm {
			if alert.Name == targetName {
				if groupID, ok := nameToGroupID[targetName]; ok {
					notifyid = *groupID
				}
				orgname, _ := GetAlarmClusterName(alert.ID)
				template, _ := getTemplate(orgname, targetName, notifyid)
				PutAlarm(template, alert.ID)
			}
		}
	}
}

func PutAlarm(template map[string]interface{}, alarmId int) {
	alarmID := strconv.Itoa(alarmId)
	u := Url("/api/orgCenter/alerts/alarmId", nil, alarmID)
	fmt.Println(u)
	accesstoken, _ := GetAccessToken("/api/OrgCenter/alerts/<id>", "PUT")
	_, err := DoRequest(Request{
		Method: "PUT",
		URL:    u,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accesstoken},
		Body:   template,
	})
	if err != nil {
		fmt.Println("出错了", err)
	}
}
