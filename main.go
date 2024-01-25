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
	erdapi.GetClusterInfo("配置文件路径")

	//获取对应的组织id
	orgNames, _ := erdapi.GetOrgId()
	fmt.Print("the org is", orgNames)

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

	//告警的获取和创建
	alarmInfos, err := erdapi.GetAlarmID()
	if err != nil {
		fmt.Println("Error getting alarm IDs: ", err)
		return
	}
	for _, alarmInfo := range alarmInfos {
		fmt.Printf("Alarm name: %s, ID: %d\n", alarmInfo.Name, alarmInfo.ID)
	}

}
func (h *AlertGroup) Handle() error {
	if h.value == 0 {
		newID, err := h.groupNotExist()
		if err != nil {
			return err
		}
		h.value = newID
	}
	//return h.groupNotExist(h.value)
	return nil
}
