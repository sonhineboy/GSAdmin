package initialize

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/sonhineboy/gsadmin/service/config"
	"github.com/spf13/viper"
)

func ConfigInit(path string) *config.Config {
	var c config.Config

	myViper := viper.New()
	myViper.SetConfigFile(path + "config.yaml")
	err := myViper.ReadInConfig()

	if err != nil {
		panic(errors.WithStack(err))
	}

	err = myViper.Unmarshal(&c, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "yaml"
	})

	fmt.Println(c.Db.MaxOpenConns)
	if err != nil {
		panic(err)
	}
	return &c
}
