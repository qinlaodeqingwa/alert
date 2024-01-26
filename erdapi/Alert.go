package erdapi

import "fmt"

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
