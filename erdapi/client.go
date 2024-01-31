package erdapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

var client = &http.Client{}

type Request struct {
	Method string
	URL    string
	Header map[string]string
	Body   interface{}
}

// DoRequest performs HTTP request
func DoRequest(r Request) ([]byte, error) {
	var buf *bytes.Buffer
	if r.Body != nil {
		body, _ := json.Marshal(r.Body)
		buf = bytes.NewBuffer(body)
	} else {
		buf = &bytes.Buffer{}
	}

	req, err := http.NewRequest(r.Method, r.URL, buf)
	if err != nil {
		return nil, err
	}
	for key, val := range r.Header {
		req.Header.Add(key, val)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// GetAccessToken 获取token
func GetAccessToken(path, fangsh string) (string, error) {
	tokenUrl := "https://openapi.erda.cloud/oauth2/token?grant_type=client_credentials&client_id=pipeline&client_secret=devops%2Fpipeline"
	//tokenUrl := Url("/oauth2/token", url.Values{
	//	"grant_type":    {"client_credentials"},
	//	"client_id":     {"pipeline"},
	//	"client_secret": {"devops/pipeline"},
	//}, "")

	payloadData := Payload{
		Metadata: map[string]string{
			"Internal-Client": "bundle",
			"USER-ID":         "1004235",
			"Org-ID":          "100812",
		},
		AccessTokenExpiredIn: "0",
		AccessibleAPIs: []map[string]string{
			{
				"path":   path,
				"method": fangsh,
				"schema": "http",
			},
		},
	}
	respBody, err := DoRequest(Request{
		Method: "POST",
		URL:    tokenUrl,
		Header: map[string]string{"Content-Type": "application/json"},
		Body:   payloadData,
	})

	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", errors.New("Access token not found in the response")
	}
	return accessToken, nil
}

func Url(p string, q url.Values, alarmId string) string {
	BaseUrl := "https://openapi.erda.cloud"
	//BaseUrl := os.Getenv("OPENAPI_URL")
	if BaseUrl == "" {
		BaseUrl = "http://erda-server.default.svc.cluster.local:9529"
	}
	u, err := url.Parse(BaseUrl)
	if err != nil {
		log.Panicln(err)
	}
	if strings.Contains(p, "alarmId") {
		p = strings.Replace(p, "alarmId", alarmId, -1)
	}
	u.Path = path.Join(u.Path, p)

	if q != nil {
		if u.RawQuery == "" {
			u.RawQuery = q.Encode()
		} else {
			raw := u.Query()
			for k, v := range q {
				raw[k] = v
			}
			u.RawQuery = raw.Encode()
		}
	}
	return u.String()
}
