package mailer

import (
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type Mailer struct {
	config        *MailerConfig
	auth          Auth
	isLogExternal bool
	mux           sync.Mutex
	pm            *manager.Manager
}

// NewMailer ...
func NewMailer(options ...MailerOption) *Mailer {
	config, simpleConfig, err := NewConfig()
	mailer := &Mailer{
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		config: &config.Mailer,
	}

	if mailer.isLogExternal {
		mailer.pm.Reconfigure(manager.WithLogger(log))
	}

	if err != nil {
		log.Error(err.Error())
	} else {
		mailer.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Mailer.Log.Level)
		log.Debugf("setting log level to %s", level)
		log.Reconfigure(logger.WithLevel(level))
	}

	mailer.auth = PlainAuth(mailer.config.Identity, mailer.config.Username, mailer.config.Password, mailer.config.Host)

	mailer.Reconfigure(options...)

	return mailer
}

func (e *Mailer) SendMessage() *SendMessageService {
	return NewSendMessageService(e)
}
