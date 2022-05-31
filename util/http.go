package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

var (
	HttpClient *http.Client
)

func InitHttpClient() {
	HttpClient = &http.Client{
		Timeout: 30 * time.Second, // 请求超时时间
	}
}

// UploadFile 发送文件上传请求
func UploadFile(url string, nameField, filename string, file io.Reader, headers map[string]string) ([]byte, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	// 在表单中创建一个文件字段
	formFile, err := writer.CreateFormFile(nameField, filename)
	if err != nil {
		return nil, err
	}
	// 读取文件内容到表单文件字段
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}
	// 构造请求对象
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	// 发送请求
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Post(url string, params map[string]string) ([]byte, error) {
	param := ""
	for key, val := range params {
		param += key + "=" + val + "&"
	}
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(param))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
