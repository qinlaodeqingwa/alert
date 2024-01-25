package erdapi

import (
	"encoding/json"
	"fmt"
)

func GetOrgId() (string, error) {
	body, _ := GetOrg()
	data := &Response{}
	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return "", err
	}
	for _, item := range data.Data.List {
		if item.Name == "ca" {
			return item.Name, nil
		}
	}
	return " ", nil
}

func GetOrg() ([]byte, error) {
	OrgUrl := Url("/api/orgs", nil, "")

	accessToken, err := GetAccessToken("/api/orgs")
	if err != nil {
		fmt.Println("Error on GetAccessToken:", err)
	}
	respBody, _ := DoRequest(Request{
		Method: "GET",
		URL:    OrgUrl,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken},
	})

	return respBody, nil
}

//func Getcluser() {
//	body, _ := GetcluterInfo()
//
//}

func GetcluterInfo() ([]byte, error) {
	//https://dice.erda.cloud/api/hyjtsc/clusters?orgID=100812
	clusterUrl := Url("/api/cluster", nil, "")
	accessToken, err := GetAccessToken("/api/cluster")
	if err != nil {
		fmt.Println("Error on GetAccessToken:", err)
	}
	respBody, _ := DoRequest(Request{
		Method: "GET",
		URL:    clusterUrl,
		Header: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken},
	})

	return respBody, nil
}
