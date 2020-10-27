package goutil

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
	_ "sync"
)

type ConfigUtil interface {
	ReadYamlConfig(path string) *SYSConfig
}

type MysqlConfig struct {
	Connection string `yaml:"db_connection"`
	Host       string `yaml:"db_host"`
	Port       string `yaml:"db_port"`
	DataBase   string `yaml:"db_database"`
	Charset    string `yaml:"db_charset"`
	UserName   string `yaml:"db_username"`
	Password   string `yaml:"db_password"`
}
type LogConfig struct {
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}
type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}
type JsonWebTokenConfig struct {
	SecretKey string `yaml:"secret_key"`
}

type JsonCodeNumberConfig struct {
	Failed         int `yaml:"failed"`
	Success        int `yaml:"success"`
	TokenError     int `yaml:"token_error"`
	ParameterError int `yaml:"parameter_error"`
	Nothing        int `yaml:"nothing"`
}

type JsonCodeDescriptionConfig struct {
	Failed         string `yaml:"failed"`
	Success        string `yaml:"success"`
	TokenError     string `yaml:"token_error"`
	ParameterError string `yaml:"parameter_error"`
	Nothing        string `yaml:"nothing"`
}

type SYSConfig struct {
	AppName             string                    `yaml:"app_name"`
	Mysql               MysqlConfig               `yaml:"mysql"`
	Log                 LogConfig                 `yaml:"log"`
	Redis               RedisConfig               `yaml:"redis"`
	JsonWebToken        JsonWebTokenConfig        `yaml:"json_web_token"`
	JsonCodeNumber      JsonCodeNumberConfig      `yaml:"json_code_number"`
	JsonCodeDescription JsonCodeDescriptionConfig `yaml:"json_code_description"`
}

type configUtil struct {
}

func NewYamlUtil() ConfigUtil {
	return &configUtil{}
}

var Config *SYSConfig
var once sync.Once

func GetConfigInstance(path string) *SYSConfig {
	once.Do(func() {
		Config = &SYSConfig{}
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

func (u *configUtil) ReadYamlConfig(path string) *SYSConfig {
	Config = &SYSConfig{}
	yamlFile, err := ioutil.ReadFile(path)
	//fmt.Println("yamlFile:", yamlFile)
	if err != nil {
		//fmt.Printf("yamlFile.Get err #%v ", err)
		return nil
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		//fmt.Printf("Unmarshal: %v", err)
		return nil
	}
	//fmt.Println("conf", conf)
	return Config
}
