package erdapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type Data struct {
	List []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		// 根据你的数据，可能还需要其他的字段
	} `json:"list"`
}

type Response struct {
	Data Data `json:"data"`
}

type Payload struct {
	Metadata             map[string]string   `json:"metadata"`
	AccessTokenExpiredIn string              `json:"accessTokenExpiredIn"`
	AccessibleAPIs       []map[string]string `json:"accessibleAPIs"`
}

func CheckAlertGroupExistence() (string, string, error) {
	body, _ := RetrieveAlertGroups()
	data := &Response{}
	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return "", "", err
	}

	var erdaL1ID, erdaL2ID string
	var err error
	for _, item := range data.Data.List {
		if item.Name == "Erda-L1(勿删)" {
			erdaL1ID = item.ID
		} else if item.Name == "Erda-L2(勿删)" {
			erdaL2ID = item.ID
		}
	}

	//ToDo 可以通过从yaml文件中获取钉钉得地址
	if erdaL1ID == "" {
		erdaL1ID, err = CreateAlertGroupsL1("Erda-L1(勿删)")
		if err != nil {
			fmt.Println("Unable to create Erda-L1: ", err)
			return erdaL1ID, "", err
		}
	}
	if erdaL2ID == "" {

		erdaL2ID, err = CreateAlertGroupsL1("Erda-L2(勿删)")
		if err != nil {
			fmt.Println("Unable to create Erda-L2: ", err)
			return "", erdaL2ID, err
		}
	}

	return erdaL1ID, erdaL2ID, nil
}

func RetrieveAlertGroups() ([]byte, error) {
	tokenUrl := Url("/oauth2/token", url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {"pipeline"},
		"client_secret": {"devops/pipeline"},
	}, "")

	accessToken, err := GetAccessToken(tokenUrl)
	if err != nil {
		fmt.Println("Error on GetAccessToken:", err)
	}
	os.Setenv("Token", accessToken)

	alertGroupUrl := "https://dice.erda.cloud/api/hyjtsc/notify-groups?scopeType=org&scopeId=100812&pageNo=1&pageSize=10"
	//alertGroupUrl := Url("/api/orgname/notify-groups", url.Values{
	//	"scopeType": {"org"},
	//	"scopeId":   {"100812"},
	//	"pageNo":    {"1"},
	//	"pageSize":  {"10"},
	//}, "")

	respBody, _ := DoRequest(Request{
		Method: "GET",
		URL:    alertGroupUrl,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken},
	})

	return respBody, nil
}
