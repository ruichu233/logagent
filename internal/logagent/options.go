package logagent

import "github.com/ruichu233/logagent/pgk/options"

type Options struct {
	KafkaOptions    *options.KafkaOptions    `mapstructure:"kafka"`
	TailFileOptions *options.TailFileOptions `mapstructure:"tail-file"`
	EtcdOptions     *options.EtcdOptions     `mapstructure:"etcd"`
	// tailFile 与 read Writer 通信的缓存大小
	BuffSize int `mapstructure:"buff-size"`
}

// NewOptions 创建 Options zero 值
func NewOptions() *Options {
	return &Options{
		KafkaOptions:    options.NewKafkaOptions(),
		TailFileOptions: options.NewTailFileOptions(),
		EtcdOptions:     options.NewEtcdOptions(),
		BuffSize:        1000,
	}
}

var DefaultOptions = NewOptions()
