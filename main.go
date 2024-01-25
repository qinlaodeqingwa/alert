package main

import (
	"fmt"
	"updata_alerm/erdapi"
)

func main() {
	alertItemL1, alertItemL2, alertItemL2NoProd, err := erdapi.ProcessTemplateAndData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(alertItemL1)
	fmt.Println(alertItemL2)
	fmt.Println(alertItemL2NoProd)

	AlertIdL1, AlertIdL2, err := erdapi.CheckAlertGroupExistence()
	if err != nil {
		fmt.Errorf("", err)
	}
	fmt.Println("jieguo", AlertIdL1, AlertIdL2)

}
