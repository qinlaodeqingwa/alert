package main

import (
	"fmt"
	"log"
	"updata_alerm/erdapi"
)

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
	fmt.Println("the groupID is", groupL1, groupL2)

	//根据集群判断是创建还是更新
	erdapi.HandleClusterAndAlertGroups(erdapi.GetCluster, erdapi.GetAlarmID)

}
