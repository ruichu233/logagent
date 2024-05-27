package options

type ESOptions struct {
	Addresses []string `mapstructure:"addresses"`
}

func NewESOptions() *ESOptions {
	return &ESOptions{
		Addresses: []string{"http://127.0.0.1:9200"},
	}
}
