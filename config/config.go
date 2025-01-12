package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Env              string `json:"env"`                // 环境
	OssAddr          string `json:"oss_addr"`           // OSS endpoint
	OssID            string `json:"oss_id"`             // OSS key
	OssSecret        string `json:"oss_secret"`         // OSS secret
	Bucket           string `json:"bucket"`             // OSS bucket
	MysqlUser        string `json:"mysql_user"`         // mysql user
	MysqlPassword    string `json:"mysql_password"`     // mysql password
	MysqlIP          string `json:"mysql_ip"`           // mysql ip
	StreamServerPort string `json:"stream_server_port"` // stream_server port
	ApiPort          string `json:"api_port"`           // api port
	SchedulerPort    string `json:"scheduler_port"`     // scheduler port
	WebPort          string `json:"web_port"`           // web port
	Address          string `json:"address"`            // 本地地址
}

var DefaultConfig *Configuration

func InitConfig(ConfigFile string) {
	file, _ := os.Open(ConfigFile)
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)
	DefaultConfig = &Configuration{}
	if err := json.NewDecoder(file).Decode(DefaultConfig); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
