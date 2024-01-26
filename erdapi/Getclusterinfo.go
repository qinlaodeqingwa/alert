package erdapi

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type diceConfig struct {
	OrganizationName string        `yaml:"organization_name"`
	ClusterName      string        `yaml:"cluster_name"`
	Install          InstallConfig `yaml:"installs"`
}

type InstallConfig struct {
	Dice DiceConfig `yaml:"dice"`
}

type DiceConfig struct {
	CentralURLs CentralURLsConfig `yaml:"central_urls"`
}

type CentralURLsConfig struct {
	CollectorURL string `yaml:"collector"`
	OpenAPIURL   string `yaml:"openapi"`
}

func GetDiceInfo(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("读取文件错误: %v", err)
	}
	var config diceConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("解析YAML错误: %v", err)
	}
	os.Setenv("ORG_NAME", config.OrganizationName)
	os.Setenv("CLUSTER_NAME", config.ClusterName)
	os.Setenv("COLLECTOR_URL", config.Install.Dice.CentralURLs.CollectorURL)
	os.Setenv("OPENAPI_URL", config.Install.Dice.CentralURLs.OpenAPIURL)
	return nil
}
