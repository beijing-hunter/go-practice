package worker

import (
	"encoding/json"
	"io/ioutil"
)

//程序配置
type Config struct {
	EtcdEndpoints   []string `json:"etcdEndpoints"`
	EtcdDialTimeout int      `json:"etcdDialTimeout"`
}

var (
	G_config *Config
)

//初始化配置
func InitConfig(filename string) (err error) {

	var (
		content []byte
		config  Config
	)

	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = json.Unmarshal(content, &config); err != nil {
		return
	}

	G_config = &config
	return
}
