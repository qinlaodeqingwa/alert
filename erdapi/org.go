package erdapi

import (
	"encoding/json"
	"fmt"
)

type OrgAlert struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OrgResponse struct {
	Data struct {
		List []OrgAlert `json:"list"`
	} `json:"data"`
}

type ClusterResponse struct {
	Success bool `json:"success"`
	Err     struct {
		Code string      `json:"code"`
		Msg  string      `json:"msg"`
		Ctx  interface{} `json:"ctx"`
	} `json:"err"`
	Data struct {
		Ready   []string    `json:"ready"`
		UnReady interface{} `json:"unReady"`
	} `json:"data"`
}

type OrgInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetOrgId(orgname string) (int, error) {
	body, _ := GetOrg()
	data := &OrgResponse{}
	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return 0, err
	}
	for _, item := range data.Data.List {
		if item.Name == orgname {
			return item.ID, nil
		}
	}
	return 0, nil
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

// GetCluster 获取集群的id和name
func GetCluster() []OrgInfo {
	body, _ := GetClusterInfo()

	orgResult := &ClusterResponse{}
	if err := json.Unmarshal(body, orgResult); err != nil {
		fmt.Println("JSON 解析错误:", err)
		return nil
	}
	if !orgResult.Success {
		//return nil, fmt.Errorf("获取 org names 错误: %v", orgResult.Err)
		return nil
	}

	orgInfoList := make([]OrgInfo, len(orgResult.Data.Ready))
	for i, org := range orgResult.Data.Ready {
		orgInfoList[i] = OrgInfo{Name: org}
	}

	return orgInfoList
}

func GetClusterInfo() ([]byte, error) {
	clusterUrl := "http://erda-server.default.svc.cluster.local:9095/api/k8s/clusters"

	respBody, _ := DoRequest(Request{
		Method: "GET",
		URL:    clusterUrl,
		Header: map[string]string{"org-ID": "100812"},
	})
	return respBody, nil
}
