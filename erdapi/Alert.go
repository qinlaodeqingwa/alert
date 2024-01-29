package erdapi

import (
	"fmt"
	"log"
	"strings"
)

type GetClustersFunc func() []OrgInfo
type GetAlarmInfo func() []Alert

func HandleClusterAndAlertGroups(getClusters GetClustersFunc, Alarm GetAlarmInfo) {
	clusters := getClusters()
	alertGroups := Alarm()
	requiredAlertGroups := []string{"Erda-L1-prod(勿删)", "Erda-L2-noprod(勿删)", "Erda-L2(勿删)"}

	fmt.Printf("当前有 %d 个集群\n", len(clusters))
	fmt.Printf("当前有 %d 个告警组\n", len(alertGroups))

	switch len(clusters) {
	case 1:
		if len(alertGroups) < 2 {
			fmt.Println("未满足要求，需要创建缺失的告警组 ")
			for _, group := range requiredAlertGroups {
				if !Contains(alertGroups, group) {
					fmt.Println("需要创建的告警组为 ", group)
				}
			}
		} else {
			fmt.Println("告警组数量满足要求，无需创建 ")

		}
	case 2:
		fallthrough
	default:
		if len(alertGroups) < 3 {
			fmt.Println("未满足要求，需要创建缺失的告警组 ")
			for _, group := range requiredAlertGroups {
				if !Contains(alertGroups, group) {
					fmt.Println("需要创建的告警组为 ", group)
					err := CreateAlarmGroup(group)
					if err != nil {
						fmt.Println("创建失败", err)
					}
				}
			}
		} else {
			fmt.Println("告警组数量满足要求，无需创建 ")
		}
	}
}

// Contains 判断切片中是否包含某个字符串的函数
func Contains(slice []Alert, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s.Name] = struct{}{}
	}
	_, ok := set[item]
	return ok
}
func CreateAlarmGroup(name string) error {
	alertItemL1, alertItemL2NoProd, alertItemL2, err := ProcessTemplateAndData(name)
	if err != nil {
		return fmt.Errorf("处理模板和数据时出错: %w", err)
	}

	accessToken, _ := GetAccessToken("/api/orgCenter/alerts")
	alertGroupUrl := Url("/api/orgCenter/alerts", nil, "")
	template := getTemplate(name, alertItemL1, alertItemL2NoProd, alertItemL2)

	if template != nil {
		_, err := DoRequest(Request{
			Method: "POST",
			URL:    alertGroupUrl,
			Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken},
			Body:   template,
		})
		if err != nil {
			return fmt.Errorf("执行请求时出错: %w", err)
		}
	} else {
		log.Println("找不到与以下内容匹配的模板: ", name)
		return fmt.Errorf("找不到与 %v 相匹配的模板", name)
	}

	return nil
}

func getTemplate(name string, alertItemL1, alertItemL2NoProd, alertItemL2 map[string]interface{}) map[string]interface{} {
	switch {
	case strings.Contains(name, "L1"):
		return alertItemL1
	case strings.Contains(name, "L2-noprod"):
		return alertItemL2NoProd
	case strings.Contains(name, "L2"):
		return alertItemL2
	default:
		return nil
	}
}
