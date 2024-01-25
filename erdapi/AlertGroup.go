package erdapi

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	List []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
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

func CheckAlertGroupExistence() (int, int, error) {
	body, _ := RetrieveAlertGroups()
	data := &Response{}
	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return 0, 0, err
	}

	var erdaL1ID, erdaL2ID int
	for _, item := range data.Data.List {
		if item.Name == "Erda-L1(勿删)" {
			erdaL1ID = item.ID
		} else if item.Name == "Erda-L2(勿删)" {
			erdaL2ID = item.ID
		}
	}
	return erdaL1ID, erdaL2ID, nil
}

func RetrieveAlertGroups() ([]byte, error) {

	accessToken, err := GetAccessToken("/api/notify-groups")
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
