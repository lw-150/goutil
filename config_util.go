package goutil

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
	_ "sync"
)

type mysqlConfig struct {
	Connection string `yaml:"db_connection"`
	Host       string `yaml:"db_host"`
	Port       string `yaml:"db_port"`
	DataBase   string `yaml:"db_database"`
	Charset    string `yaml:"db_charset"`
	UserName   string `yaml:"db_username"`
	Password   string `yaml:"db_password"`
}
type logConfig struct {
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}
type redisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}
type jsonWebTokenConfig struct {
	SecretKey string `yaml:"secret_key"`
}

type jsonCodeNumberConfig struct {
	Failed         int `yaml:"failed"`
	Success        int `yaml:"success"`
	TokenError     int `yaml:"token_error"`
	ParameterError int `yaml:"parameter_error"`
	Nothing        int `yaml:"nothing"`
}

type jsonCodeDescriptionConfig struct {
	Failed         string `yaml:"failed"`
	Success        string `yaml:"success"`
	TokenError     string `yaml:"token_error"`
	ParameterError string `yaml:"parameter_error"`
	Nothing        string `yaml:"nothing"`
}

type sysConfig struct {
	AppName             string                    `yaml:"app_name"`
	Mysql               mysqlConfig               `yaml:"mysql"`
	Log                 logConfig                 `yaml:"log"`
	Redis               redisConfig               `yaml:"redis"`
	JsonWebToken        jsonWebTokenConfig        `yaml:"json_web_token"`
	JsonCodeNumber      jsonCodeNumberConfig      `yaml:"json_code_number"`
	JsonCodeDescription jsonCodeDescriptionConfig `yaml:"json_code_description"`
}

var Config *sysConfig
var once sync.Once

func GetConfigInstance(path string) *sysConfig {
	once.Do(func() {
		Config = &sysConfig{}
		yamlFile, err := ioutil.ReadFile(path)
		//fmt.Println("yamlFile:", yamlFile)
		if err != nil {
			fmt.Printf("yamlFile.Get err #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, Config)
		if err != nil {
			fmt.Printf("Unmarshal: %v", err)
		}
	})
	return Config
}
