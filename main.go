package main

import (
	"fmt"
	"log"
	"updata_alerm/erdapi"
)

//
//type notifyGroup struct {
//	value int
//	//groupExist    func() ()
//	groupNotExist func() (int, error)
//}
//
//type NotifyGroupIds struct {
//	GroupL1 int
//	GroupL2 int
//}
//
//var GroupIds NotifyGroupIds

func main() {
	//获取环境的一些参数
	erdapi.GetDiceInfo("配置文件路径")

	//获取对应的组织id
	OrgNameId, _ := erdapi.GetOrgId("hyjtsc")
	fmt.Println("the org is", OrgNameId)

	//检查通知组是否存在
	groupL1, groupL2, err := erdapi.CheckNotifyGroupExistence()
	if err != nil {
		log.Println("获取告警组失败", err)
	}
	//告警通知组ID获取和创建
	//notifyGroupCheckers := []*notifyGroup{
	//	{value: groupL1, groupNotExist: erdapi.CreateNotifyGroupsL1},
	//	{value: groupL2, groupNotExist: erdapi.CreateNotifyGroupsL2},
	//}
	//for i, check := range notifyGroupCheckers {
	//	newID := check.Handle()
	//	if i == 0 {
	//		GroupIds.GroupL1 = newID
	//	} else if i == 1 {
	//		GroupIds.GroupL2 = newID
	//	}
	//}
	fmt.Println("the groupID is", groupL1, groupL2)
	////告警的获取
	//alarmInfos := erdapi.GetAlarmID()
	//for _, alarmInfo := range alarmInfos {
	//	fmt.Printf("Alarm name: %s, ID: %d\n", alarmInfo.Name, alarmInfo.ID)
	//}
	//
	////集群的获取
	//orgInfos := erdapi.GetCluster()
	//
	//for _, orgInfo := range orgInfos {
	//	fmt.Printf("orgName: %s\n", orgInfo.Name)
	//}

	//根据集群判断是创建还是更新
	erdapi.HandleClusterAndAlertGroups(erdapi.GetCluster, erdapi.GetAlarmID)

}

//func (h *notifyGroup) Handle() int {
//	if h.value == 0 {
//		newID, _ := h.groupNotExist()
//		h.value = newID
//		return newID
//	}
//	return h.value
//}
