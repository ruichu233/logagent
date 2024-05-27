package logtransfer

import (
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(path.Dir(filename)))

	viper.AddConfigPath(root + "\\configs")
	viper.SetConfigName("logtransfer")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	logrus.Infoln(viper.ConfigFileUsed())
}

func GetOptions() (*Options, error) {
	opts := NewOptions()
	err := viper.Unmarshal(opts)
	if err != nil {
		return nil, err
	}
	return opts, nil
}
