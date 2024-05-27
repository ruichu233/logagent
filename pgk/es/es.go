package es

import (
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/ruichu233/logagent/pgk/options"
)

func NewClient(opts *options.ESOptions) *elasticsearch8.TypedClient {
	cfg := elasticsearch8.Config{
		Addresses: opts.Addresses,
	}
	client, err := elasticsearch8.NewTypedClient(cfg)
	if err != nil {
		return nil
	}
	return client
}
