package logtransfer

import "github.com/ruichu233/logagent/pgk/options"

type Options struct {
	KafKaOptions *options.KafkaOptions `mapstructure:"kafka"`
	ESOptions    *options.ESOptions    `mapstructure:"es"`
	Index        string                `mapstructure:"index"`
	BufferSize   int                   `mapstructure:"buffer-size"`
}

func NewOptions() *Options {
	return &Options{
		KafKaOptions: options.NewKafkaOptions(),
		ESOptions:    options.NewESOptions(),
		Index:        "index",
		BufferSize:   100,
	}
}
