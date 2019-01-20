package mailer

import (
	"fmt"

	gomanager "github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Mailer *MailerConfig `json:"mailer"`
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
func NewConfig() (*MailerConfig, error) {
	appConfig := &AppConfig{}
	if _, err := gomanager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig); err != nil {
		log.Error(err.Error())

		return &MailerConfig{}, err
	}

	return appConfig.Mailer, nil
}
