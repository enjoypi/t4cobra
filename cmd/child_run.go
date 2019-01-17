package cmd

import (
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// config for sub
type cfgChild struct {
	Config string
	Etcd   string
	Child  child
}

type child struct {
	Bool bool
	Test bool
	Str  string
}

func childRun(v *viper.Viper) error {
	var c cfgChild
	if err := mapstructure.Decode(v.AllSettings(), &c); err != nil {
		return err
	}
	logrus.Infof("settings on child: %+v", c)
	return nil
}
