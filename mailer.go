package mailer

import (
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type Mailer struct {
	config        *MailerConfig
	auth          Auth
	logger        logger.ILogger
	isLogExternal bool
	mux           sync.Mutex
	pm            *manager.Manager
}

// NewMailer ...
func NewMailer(options ...MailerOption) *Mailer {
	config, simpleConfig, err := NewConfig()

	service := &Mailer{
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		config: &config.Mailer,
		logger: logger.NewLogDefault("mailer", logger.WarnLevel),
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Mailer.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	service.auth = PlainAuth(service.config.Identity, service.config.Username, service.config.Password, service.config.Host)

	service.Reconfigure(options...)

	return service
}

func (e *Mailer) SendMessage() *SendMessageService {
	return NewSendMessageService(e)
}
