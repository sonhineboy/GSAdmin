package test

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"testing"
)

func TestLoad(t *testing.T) {

	myViper := viper.New()

	myViper.SetConfigFile("./c.yaml")

	err := myViper.ReadInConfig()
	if err != nil {
		t.Error(fmt.Errorf("reding err %w ", errors.WithStack(err)))
	}
	t.Logf("c--->%v", myViper.AllKeys())
}
