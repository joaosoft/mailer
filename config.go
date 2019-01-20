package mailer

import (
	"fmt"
	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Mailer MailerConfig `json:"mailer"`
}

// MailerConfig ...
type MailerConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Identity string `json:"identity"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	if err != nil {
		log.Error(err.Error())
	}

	return appConfig, simpleConfig, err
}