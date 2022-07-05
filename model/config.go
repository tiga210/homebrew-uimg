package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"uimg/api"
	"uimg/util"
)

const aesKey = "1234567890123456"

type Config struct {
	Account  string
	Password string
	Host     string
	Token    string
}

func getDirAndPath() (string, string) {
	sysUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	configDir := sysUser.HomeDir + "/.config"
	configPath := configDir + "/uimg.json"
	return configDir, configPath
}

func InitConfig() {
	configDir, configPath := getDirAndPath()

	fmt.Println("please input host:")
	host, _ := util.GetInput()

	fmt.Println("please input account:")
	account, _ := util.GetInput()

	fmt.Println("please input password:")
	password, _ := util.GetPassword()

	c := &Config{}

	if account != "" {
		c.Account = util.AesEncrypt(account, aesKey)
	}
	if password != "" {
		c.Password = util.AesEncrypt(password, aesKey)
	}

	if host != "" {
		c.Host = host
	}

	if c.Account != "" && c.Password != "" && c.Host != "" {
		token := api.FetchToken(account, password, c.Host)
		c.Token = util.AesEncrypt(token, aesKey)
	}

	jsonConfig, _ := json.Marshal(c)
	os.MkdirAll(configDir, 0755)
	ioutil.WriteFile(configPath, jsonConfig, 0666)
}

func GetConfig() *Config {
	_, path := getDirAndPath()
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	c := &Config{}
	json.Unmarshal(bytes, &c)
	c.Account = util.AesDecrypt(c.Account, aesKey)
	c.Password = util.AesDecrypt(c.Password, aesKey)
	c.Token = util.AesDecrypt(c.Token, aesKey)
	return c
}
