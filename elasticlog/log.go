package elasticlog

import (
	"fmt"
	"github.com/SergeiSadov/positions/config"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

var Log *logrus.Logger

func StartLog() error {
	Log = logrus.New()
	client, err := elastic.NewSimpleClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%s", config.Config.Elastic.Host,
			config.Config.Elastic.Port)),
		elastic.SetBasicAuth(config.Config.Elastic.Username, config.Config.Elastic.Password),
		elastic.SetSniff(false),
	)
	if err != nil {
		return fmt.Errorf("server.startLog() error #1: \t\n %w \r\n", err)
	}

	hook, err := elogrus.NewAsyncElasticHook(client, config.Config.Elastic.Host, logrus.ErrorLevel, "error")
	if err != nil {
		return fmt.Errorf("server.startLog() error #2: \t\n %w \r\n", err)
	}

	Log.Hooks.Add(hook)

	Log.Info("log started")
	Log.Error("log started")

	return nil
}
