package api

import (
	"encoding/json"
	"os"
	"uimg/util"
)

const tokenApi = "/api/v1/tokens"
const uploadApi = "/api/v1/upload"

func FetchToken(account string, password string, host string) string {
	params := map[string]string{
		"email":    account,
		"password": password,
	}
	result, _ := util.Post(host+tokenApi, params)
	token := &Response[TokenData]{}
	json.Unmarshal(result, &token)
	if token.Status {
		return token.Data.Token
	}
	panic(result)
}

func UploadImg(fileName string, token string, host string) (bool, string) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Accept":        "application/json",
	}
	file, _ := os.Open(fileName)
	util.InitHttpClient()
	result, _ := util.UploadFile(host+uploadApi, "file", fileName, file, headers)
	response := &Response[ImgData]{}
	json.Unmarshal(result, &response)
	return response.Status, response.Data.Links.Url
}
