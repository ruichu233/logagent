package options

import "time"

type EtcdOptions struct {
	Endpoints   []string      `mapstructure:"endpoints"`
	Username    string        `mapstructure:"username"`
	Password    string        `mapstructure:"password"`
	DialTimeout time.Duration `mapstructure:"dial-timeout"`
}

func NewEtcdOptions() *EtcdOptions {
	return &EtcdOptions{
		Endpoints: []string{"127.0.0.1:2379"},
		Username:  "",
		Password:  "",
	}

}
