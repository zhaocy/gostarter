package gostarter

import (
	"errors"
	"github.com/spf13/viper"
)

type Application struct {
	conf           *viper.Viper
	starterContext *StarterContext
}

func New(conf *viper.Viper) (*Application, error) {
	app := &Application{
		starterContext: &StarterContext{},
	}

	if err := conf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		} else {
			return nil, err
		}
	} else {
		err := conf.Unmarshal(app.starterContext)
		if err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (app *Application) GetStartContext() *StarterContext {
	return app.starterContext
}

func (app *Application) EnableDb() {

}
