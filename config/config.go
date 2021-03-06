package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const PROJECT_NAME string = "gin-web-base-template"

type JSONConfig struct {
	ENV         string
	ProjectName string
	Server      Server
	GORM        GORM
	Redis       Redis
}

type Server struct {
	Port int
}

type GORM struct {
	Driver  string
	Open    string
	MaxIdle int
	MaxOpen int
}

type Redis struct {
	MaxIdle  int
	Network  string
	Address  string
	Password string
	AuthKey  string
	DB       int
}

var Config JSONConfig

func init() {
	InitConfig(&Config)
}

func (config *JSONConfig) ConfigFolderPath() string {
	var projectGoPath string
	// 在GOPATH 寻找项目存在的那一条路径
	//
	goPaths := strings.Split(os.Getenv("GOPATH"), ":")
	for _, goPath := range goPaths {
		if _, err := os.Stat(filepath.Join(goPath, "src", PROJECT_NAME)); err == nil {
			projectGoPath = goPath
			break
		}
	}
	return filepath.Join(projectGoPath, "src", PROJECT_NAME, "config")
}

func (config *JSONConfig) ConfigFilePath(env string) string {
	return filepath.Join(config.ConfigFolderPath(), fmt.Sprintf("%s.json", env))
}

func InitConfig(config *JSONConfig) {
	env := os.Getenv("GOENV")
	if env == "" {
		env = "dev"
	}

	filePath := config.ConfigFilePath(env)
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		panic(err)
	}
	config.ENV = env
}
