package erdapi

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

// Config 解析yaml
type Config struct {
	Alerts struct {
		Update struct {
			L1       []AlertDetail `yaml:"L1"`
			L2       []AlertDetail `yaml:"L2"`
			L2NoProd []AlertDetail `yaml:"L2NoProd"`
		} `yaml:"update"`
		Delete struct {
			L1       []AlertDetail `yaml:"L1"`
			L2       []AlertDetail `yaml:"L2"`
			L2NoProd []AlertDetail `yaml:"L2NoProd"`
		} `yaml:"delete"`
		//Group struct {
		//	L1 []AlertDetail `yaml:"L1"`
		//	L2 []AlertDetail `yaml:"L2"`
		//} `yaml:"groups"`
	} `yaml:"alerts"`
}

type AlertDetail struct {
	Name      string
	Value     string
	Operation string
	Level     string
}

// Todo 解析一下对应的告警方式，设置为变量，直接获取了
//type GroupDetail struct {
//	Name     string `yaml:"name"`
//	Receiver string `yaml:"receiver"`
//	Secret   string `yaml:"secret"`
//}

// AlertMapping 映射关系
var AlertMapping = map[string]string{
	"平台组件网盘不可用告警":                 "dice_component_gfs_status",
	"注册中心实例内存使用率异常":               "register_center_mem",
	"注册中心实例CPU使用率异常":              "register_center_cpu",
	"API网关实例CPU使用率异常":             "api_gateway_cpu",
	"API网关实例内存使用率异常":              "api_gateway_mem",
	"平台组件异常退出":                    "dice_component_container_exit",
	"机器宕机":                        "machine_status",
	"机器磁盘":                        "machine_disk",
	"网盘":                          "machine_netdisk",
	"磁盘Inode异常告警":                 "machine_disk_inode",
	"机器Load5":                     "machine_load5",
	"平台中间件异常退出":                   "dice_addon_container_exit",
	"kubernetes节点异常":              "kubernetes_node",
	"kubernetes组件实例内存状态":          "kubernetes_instance_mem",
	"kubernetes组件异常退出":            "kubernetes_instance_exit",
	"kubernetes组件实例Ready状态异常":     "kubernetes_instance_ready",
	"平台组件Fink吞吐量异常告警":             "dice_component_flink_throughput",
	"平台组件实例CPU状态":                 "dice_component_container_cpu",
	"平台组件实例Ready状态异常":             "dice_component_container_ready",
	"平台组件实例内存状态":                  "dice_component_container_mem",
	"平台组件Flink任务checkpoint延迟异常告警": "dice_component_flink_checkpoint_duration",
	"平台中间件Kafka GC耗时异常":           "dice_addon_kafka_gc_time",
	"平台中间件实例CPU状态":                "dice_addon_container_cpu",
	"平台中间件Cassandra GC次数异常":       "dice_addon_cassandra_gc_count",
	"平台中间件Elasticsearch GC次数异常":   "dice_addon_elasticsearch_gc_count",
	"平台中间件Cassandra GC耗时异常":       "dice_addon_cassandra_gc_time",
	"平台中间件kafka消息堆积":              "dice_addon_kafka_lag",
	"平台中间件实例内存状态":                 "dice_addon_container_mem",
	"平台中间件Elasticsearch GC耗时异常":   "dice_addon_elasticsearch_gc_time",
	"平台中间件Kafka GC次数异常":           "dice_addon_kafka_gc_count",
	"平台中间件实例Ready状态异常":            "dice_addon_container_ready",
	"网盘容量使用量异常告警":                 "machine_netdisk_used",
	"机器CPU":                       "machine_cpu",
	"机器始终一致性异常告警":                 "machine_cluster_clock",
	"机器磁盘IO":                      "machine_disk_util",
}

func translateAlertIndex(userFriendlyName string) (string, error) {
	alertIndex, ok := AlertMapping[userFriendlyName]
	if !ok {
		return "", fmt.Errorf("未知的告警名称: %s", userFriendlyName)
	}
	return alertIndex, nil
}

func updateRule(rules []map[string]interface{}, name string, newValue interface{}) []map[string]interface{} {
	for i := range rules {
		if rules[i]["alertIndex"] == name {
			functions := rules[i]["functions"].([]map[string]interface{})
			for j := range functions {
				functions[j]["value"] = newValue
			}
			break
		}
	}
	return rules
}

func deleteRule(rules []map[string]interface{}, name string) []map[string]interface{} {
	for i, rule := range rules {
		if rule["alertIndex"] == name {
			rules = append(rules[:i], rules[i+1:]...)
			break
		}
	}
	return rules
}

func AddNewAlertRules(data []byte) ([]AlertDetail, error) {
	allAlerts := Config{}
	err := yaml.Unmarshal(data, &allAlerts)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	details := processAlertLevel(allAlerts.Alerts.Update.L1, "update", "L1")
	details = append(details, processAlertLevel(allAlerts.Alerts.Update.L2, "update", "L2")...)
	details = append(details, processAlertLevel(allAlerts.Alerts.Update.L2NoProd, "update", "L2NoProd")...)
	details = append(details, processAlertLevel(allAlerts.Alerts.Delete.L1, "delete", "L1")...)
	details = append(details, processAlertLevel(allAlerts.Alerts.Delete.L2, "delete", "L2")...)
	details = append(details, processAlertLevel(allAlerts.Alerts.Update.L2NoProd, "update", "L2NoProd")...)

	return details, nil
}

// processAlertLevel 检查名称是否符合
func processAlertLevel(alerts []AlertDetail, operation string, level string) []AlertDetail {
	result := make([]AlertDetail, len(alerts))
	for i, alert := range alerts {
		alert.Operation = operation
		alert.Level = level
		name, err := translateAlertIndex(alert.Name)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
		alert.Name = name
		result[i] = alert
	}
	return result

}
