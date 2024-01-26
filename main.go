package main

import (
	"fmt"
	"log"
	"updata_alerm/erdapi"
)

type AlertGroup struct {
	value int
	//groupExist    func() ()
	groupNotExist func() (int, error)
}

func main() {
	//获取环境的一些参数
	erdapi.GetDiceInfo("配置文件路径")

	//获取对应的组织id
	orgnameId, _ := erdapi.GetOrgId("hyjtsc")
	fmt.Println("the org is", orgnameId)

	//模版文件
	alertItemL1, alertItemL2NoProd, alertItemL2, err := erdapi.ProcessTemplateAndData()
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Println(alertItemL1)
	fmt.Println(alertItemL2NoProd)
	fmt.Println(alertItemL2)

	//告警通知组ID获取和创建
	groupL1, groupL2, err := erdapi.CheckAlertGroupExistence()
	if err != nil {
		log.Println("获取告警组失败", err)
	}
	fmt.Println("the groupID is", groupL1, groupL2)

	alertGroupCheckers := []*AlertGroup{
		{value: groupL1, groupNotExist: erdapi.CreateAlertGroupsL1},
		{value: groupL2, groupNotExist: erdapi.CreateAlertGroupsL2},
	}
	for _, check := range alertGroupCheckers {
		if err := check.Handle(); err != nil {
			log.Println(err)
		}
	}

	erdapi.HandleClusterAndAlertGroups(erdapi.GetCluster, erdapi.GetAlarmID)
	////告警的获取
	//alarmInfos, err := erdapi.GetAlarmID()
	//if err != nil {
	//	fmt.Println("Error getting alarm IDs: ", err)
	//	return
	//}
	//for _, alarmInfo := range alarmInfos {
	//	fmt.Printf("Alarm name: %s, ID: %d\n", alarmInfo.Name, alarmInfo.ID)
	//}
	////集群的获取
	//orgInfos, err := erdapi.GetCluster()
	//if err != nil {
	//	fmt.Println("Error getting org Information:", err)
	//}
	//for _, orgInfo := range orgInfos {
	//	fmt.Printf("orgName: %s\n", orgInfo.Name)
	//}

}

func (h *AlertGroup) Handle() error {
	if h.value == 0 {
		newID, _ := h.groupNotExist()
		h.value = newID
	}
	return nil
}
