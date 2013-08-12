package github

import (
	"encoding/json"
	"net/http"
	"fmt"
)
type GitHub struct {
	Token    string
	Username string
	Orgs     []string
	BaseUrl  string "https://api.github.com"
}

type TokenResponse struct {
	Token string
	Id int
	Url string

}


func url(part string) string {
	return  fmt.Sprintf("https://api.github.com/%v", part)
}

func (gh GitHub) GetToken(username, password string) string {
	var r TokenResponse
	data, _ := requestData("POST", "authorize", r)

	return data.(TokenResponse).Token
}

func (gh GitHub) GetUserInfo(username string) interface{} {
	return 1
}

func requestData(method string, path string, payloadContainer interface{}) (interface{}, error) {
	httpClient := &http.Client{}

	response, err := http.Get(path)
	if err != nil {
		var f interface{}
		return parseJSON("{}", f), err
	}

	parseJSON(response.Body, payloadContainer)


}

func parseJSON(payload string, payloadContainer interface{}) interface{} {
	_ = json.Unmarshal([]byte(payload), &payloadContainer)

	return payloadContainer
}
