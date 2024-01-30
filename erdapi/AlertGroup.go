package erdapi

import (
	"encoding/json"
	"fmt"
	"os"
)

type notifyGroup struct {
	value int
	//groupExist    func() ()
	groupNotExist func() (int, error)
}

type NotifyGroupIds struct {
	GroupL1 int
	GroupL2 int
}

var GroupIds NotifyGroupIds

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

func CheckNotifyGroupExistence() (int, int, error) {
	body, _ := RetrieveAlertGroups()
	data := &Response{}
	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return 0, 0, err
	}

	var notifyL1Id, notifyL2Id int
	for _, item := range data.Data.List {
		if item.Name == "Erda-L1(勿删)" {
			notifyL1Id = item.ID
		} else if item.Name == "Erda-L2(勿删)" {
			notifyL2Id = item.ID
		}
	}
	notifyGroupCheckers := []*notifyGroup{
		{value: notifyL1Id, groupNotExist: CreateNotifyGroupsL1},
		{value: notifyL2Id, groupNotExist: CreateNotifyGroupsL2},
	}
	for i, check := range notifyGroupCheckers {
		newID := check.Handle()
		if i == 0 {
			GroupIds.GroupL1 = newID
		} else if i == 1 {
			GroupIds.GroupL2 = newID
		}
	}
	return notifyL1Id, notifyL2Id, nil

}

func RetrieveAlertGroups() ([]byte, error) {

	accessToken, err := GetAccessToken("/api/notify-groups", "GET")
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
func (h *notifyGroup) Handle() int {
	if h.value == 0 {
		newID, _ := h.groupNotExist()
		h.value = newID
		return newID
	}
	return h.value
}
