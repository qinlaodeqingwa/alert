package erdapi

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ProcessTemplateAndData() (map[string]interface{}, map[string]interface{}, map[string]interface{}, error) {
	templateL1 := map[string]interface{}{
		//"name":   name,
		//"domain": "https://dice.erda.cloud",
		//"notifies": []map[string]interface{}{
		//	{
		//		"silence": map[string]interface{}{
		//			"value":  15,
		//			"unit":   "minutes",
		//			"policy": "doubled",
		//		},
		//		"groupId":   notifyGroupID,
		//		"groupType": "dingding",
		//		"level":     "Fatal",
		//	},
		//},
		//
		//"triggerCondition": []map[string]interface{}{
		//	{
		//		"condition": "cluster_name",
		//		"operator":  "in",
		//		"values":    orgname,
		//	},
		//},
		"rules": []map[string]interface{}{
			{
				"alertIndex": "dice_component_gfs_status",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "status",
						"aggregator": "values",
						"operator":   "all",
						"value":      "Disconnected",
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"level":      "Fatal",
				"alertIndex": "register_center_mem",
				"functions": []map[string]interface{}{
					{
						"field":      "mem_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      90,
					},
				},
				"window":    3,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "register_center_cpu",
				"functions": []map[string]interface{}{
					{
						"field":      "cpu_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      90,
					},
				},
				"window":    3,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "api_gateway_cpu",
				"functions": []map[string]interface{}{
					{
						"field":      "cpu_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      90,
					},
				},
				"window":    3,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "api_gateway_mem",
				"functions": []map[string]interface{}{
					{
						"field":      "mem_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      90,
					},
				},
				"window":    3,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "dice_component_container_exit",
				"functions": []map[string]interface{}{
					{
						"field":      "exitcode",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"window":    1,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_status",
				"functions": []map[string]interface{}{
					{
						"field":      "status",
						"aggregator": "values",
						"operator":   "all",
						"value":      "not_ready",
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_disk",
				"functions": []map[string]interface{}{
					{
						"field":      "used_percent",
						"aggregator": "max",
						"operator":   "gte",
						"value":      86,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_netdisk",
				"functions": []map[string]interface{}{
					{
						"field":      "used_percent",
						"aggregator": "max",
						"operator":   "gte",
						"value":      85,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_disk_inode",
				"functions": []map[string]interface{}{
					{
						"field":      "inode_used_percent",
						"aggregator": "max",
						"operator":   "gte",
						"value":      95,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_load5",
				"functions": []map[string]interface{}{
					{
						"field":      "load5",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      50,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "dice_addon_container_exit",
				"functions": []map[string]interface{}{
					{
						"field":      "exitcode",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"window":    1,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "kubernetes_node",
				"functions": []map[string]interface{}{
					{
						"field":      "ready",
						"aggregator": "values",
						"operator":   "all",
						"value":      false,
					},
				},
				"window":    1,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "kubernetes_instance_mem",
				"functions": []map[string]interface{}{
					{
						"field":      "mem_usage_percent",
						"aggregator": "avg",
						"operator":   "gt",
						"value":      90,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "kubernetes_instance_exit",
				"functions": []map[string]interface{}{
					{
						"field":      "exitcode",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"window":    1,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "kubernetes_instance_ready",
				"functions": []map[string]interface{}{
					{
						"field":      "not_ready",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"window":    1,
				"isRecover": true,
			},
		},
	}
	templateL2Noprod := map[string]interface{}{
		//"name":   name,
		//"domain": "https://dice.erda.cloud",
		//"notifies": []map[string]interface{}{
		//	{
		//		"silence": map[string]interface{}{
		//			"value":  15,
		//			"unit":   "minutes",
		//			"policy": "doubled",
		//		},
		//		"groupId":   notifyGroupID,
		//		"groupType": "dingding",
		//		"level":     "Fatal",
		//	},
		//},
		//
		//"triggerCondition": []map[string]interface{}{
		//	{
		//		"condition": "cluster_name",
		//		"operator":  "in",
		//		"values":    orgname,
		//	},
		//},
		"rules": []map[string]interface{}{
			{
				"alertIndex": "dice_component_gfs_status",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "status",
						"aggregator": "values",
						"operator":   "all",
						"value":      "Disconnected",
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"level":      "Fatal",
				"alertIndex": "register_center_mem",
				"functions": []map[string]interface{}{
					{
						"field":      "mem_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      90,
					},
				},
				"window":    3,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "register_center_cpu",
				"functions": []map[string]interface{}{
					{
						"field":      "cpu_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      90,
					},
				},
				"window":    3,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "api_gateway_cpu",
				"functions": []map[string]interface{}{
					{
						"field":      "cpu_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      90,
					},
				},
				"window":    3,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "api_gateway_mem",
				"functions": []map[string]interface{}{
					{
						"field":      "mem_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      90,
					},
				},
				"window":    3,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "dice_component_container_exit",
				"functions": []map[string]interface{}{
					{
						"field":      "exitcode",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"window":    1,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_status",
				"functions": []map[string]interface{}{
					{
						"field":      "status",
						"aggregator": "values",
						"operator":   "all",
						"value":      "not_ready",
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_disk",
				"functions": []map[string]interface{}{
					{
						"field":      "used_percent",
						"aggregator": "max",
						"operator":   "gte",
						"value":      86,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_netdisk",
				"functions": []map[string]interface{}{
					{
						"field":      "used_percent",
						"aggregator": "max",
						"operator":   "gte",
						"value":      85,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_disk_inode",
				"functions": []map[string]interface{}{
					{
						"field":      "inode_used_percent",
						"aggregator": "max",
						"operator":   "gte",
						"value":      95,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "machine_load5",
				"functions": []map[string]interface{}{
					{
						"field":      "load5",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      50,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "dice_addon_container_exit",
				"functions": []map[string]interface{}{
					{
						"field":      "exitcode",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"window":    1,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "kubernetes_node",
				"functions": []map[string]interface{}{
					{
						"field":      "ready",
						"aggregator": "values",
						"operator":   "all",
						"value":      false,
					},
				},
				"window":    1,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "kubernetes_instance_mem",
				"functions": []map[string]interface{}{
					{
						"field":      "mem_usage_percent",
						"aggregator": "avg",
						"operator":   "gt",
						"value":      90,
					},
				},
				"window":    5,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "kubernetes_instance_exit",
				"functions": []map[string]interface{}{
					{
						"field":      "exitcode",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"window":    1,
				"isRecover": true,
			},
			{
				"level":      "Fatal",
				"alertIndex": "kubernetes_instance_ready",
				"functions": []map[string]interface{}{
					{
						"field":      "not_ready",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"window":    1,
				"isRecover": true,
			},
		},
	}
	templateL2 := map[string]interface{}{
		//"name":   name,
		//"domain": "https://dice.erda.cloud",
		//"notifies": []map[string]interface{}{
		//	{
		//		"silence": map[string]interface{}{
		//			"value":  30,
		//			"unit":   "minutes",
		//			"policy": "doubled",
		//		},
		//		"groupId":   notifyGroupID,
		//		"groupType": "dingding",
		//		"level":     "Fatal",
		//	},
		//},
		//
		//"triggerCondition": []map[string]interface{}{
		//	{
		//		"condition": "cluster_name",
		//		"operator":  "in",
		//		"values":    orgname,
		//	},
		//},
		"rules": []map[string]interface{}{
			{
				"alertIndex": "dice_component_flink_throughput",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "numRecordsOutPerSecond_count",
						"aggregator": "max",
						"operator":   "eq",
						"value":      0,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_component_container_cpu",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "cpu_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      95,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_component_container_ready",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "not_ready",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_component_container_mem",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "mem_usage_percent",
						"aggregator": "avg",
						"operator":   "gt",
						"value":      90,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_component_flink_checkpoint_duration",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "lastCheckpointDuration",
						"aggregator": "max",
						"operator":   "gte",
						"value":      3000,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_kafka_gc_time",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "CollectionTime_win",
						"aggregator": "value",
						"operator":   "gte",
						"value":      15000,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_container_cpu",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "cpu_usage_percent",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      95,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_cassandra_gc_count",
				"window":     10,
				"functions": []map[string]interface{}{
					{
						"field":      "CollectionCount_win",
						"aggregator": "value",
						"operator":   "gte",
						"value":      10,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_elasticsearch_gc_count",
				"window":     15,
				"functions": []map[string]interface{}{
					{
						"field":      "gc_collectors_old_collection_count_win",
						"aggregator": "value",
						"operator":   "gte",
						"value":      10,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_cassandra_gc_time",
				"window":     10,
				"functions": []map[string]interface{}{
					{
						"field":      "CollectionTime_win",
						"aggregator": "value",
						"operator":   "gte",
						"value":      15000,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_kafka_lag",
				"window":     10,
				"functions": []map[string]interface{}{
					{
						"field":      "lag",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      300000,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_container_mem",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "mem_usage_percent",
						"aggregator": "avg",
						"operator":   "gt",
						"value":      90,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_elasticsearch_gc_time",
				"window":     15,
				"functions": []map[string]interface{}{
					{
						"field":      "gc_collectors_old_collection_time_in_millis_win",
						"aggregator": "value",
						"operator":   "gte",
						"value":      15000,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_kafka_gc_count",
				"window":     10,
				"functions": []map[string]interface{}{
					{
						"field":      "CollectionCount_win",
						"aggregator": "value",
						"operator":   "gte",
						"value":      10,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "dice_addon_container_ready",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "not_ready",
						"aggregator": "max",
						"operator":   "gt",
						"value":      0,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "machine_netdisk_used",
				"window":     10,
				"functions": []map[string]interface{}{
					{
						"field":      "used",
						"aggregator": "max",
						"operator":   "gte",
						"value":      400,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "machine_cpu",
				"window":     5,
				"functions": []map[string]interface{}{
					{
						"field":      "cpu_usage_active",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      95,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "machine_cluster_clock",
				"window":     10,
				"functions": []map[string]interface{}{
					{
						"field":      "elapsed_abs",
						"aggregator": "avg",
						"operator":   "gte",
						"value":      5000,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
			{
				"alertIndex": "machine_disk_util",
				"window":     10,
				"functions": []map[string]interface{}{
					{
						"field":      "pct_util",
						"aggregator": "p75",
						"operator":   "gte",
						"value":      95,
					},
				},
				"isRecover": true,
				"level":     "Fatal",
			},
		},
	}

	data, err := ioutil.ReadFile("erdapi/test.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	config := map[string]interface{}{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML data: %v", err)
	}

	alertDetails, err := AddNewAlertRules(data)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Error when adding new alert rules: %v", err)
	}

	for _, detail := range alertDetails {
		switch detail.Operation {
		case "update":
			if detail.Level == "L1" {
				templateL1["rules"] = updateRule(templateL1["rules"].([]map[string]interface{}), detail.Name, detail.Value)
			} else if detail.Level == "L2" {
				templateL2["rules"] = updateRule(templateL2["rules"].([]map[string]interface{}), detail.Name, detail.Value)
			} else if detail.Level == "L2NoProd" {
				templateL2Noprod["rules"] = updateRule(templateL2Noprod["rules"].([]map[string]interface{}), detail.Name, detail.Value)
			}
		case "delete":
			if detail.Level == "L1" {
				templateL1["rules"] = deleteRule(templateL1["rules"].([]map[string]interface{}), detail.Name)
			} else if detail.Level == "L2" {
				templateL2["rules"] = deleteRule(templateL2["rules"].([]map[string]interface{}), detail.Name)
			} else if detail.Level == "L2NoProd" {
				templateL2Noprod["rules"] = deleteRule(templateL2Noprod["rules"].([]map[string]interface{}), detail.Name)
			}
		}
	}
	return templateL1, templateL2, templateL2Noprod, nil

}
