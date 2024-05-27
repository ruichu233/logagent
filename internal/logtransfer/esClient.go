package logtransfer

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ruichu233/logagent/pgk/es"
	"github.com/ruichu233/logagent/pgk/options"
	"github.com/sirupsen/logrus"
)

type ESClient struct {
	es    *elasticsearch.TypedClient
	index string
}

func NewESClient(opts *options.ESOptions, index string) *ESClient {
	return &ESClient{
		es:    es.NewClient(opts),
		index: index,
	}
}

func (e *ESClient) sentToES(ctx context.Context, ch chan interface{}) {
	for {
		select {
		case <-ctx.Done():
			return
		case obj := <-ch:
			resp, err := e.es.Index(e.index).
				Request(obj).
				Do(ctx)
			if err != nil {
				logrus.Warnf("indexing document failed,err:%v\n", err)
				return
			}
			logrus.Infof("Indexed user %s to index %s\n", resp.Id_, resp.Index_)
		}
	}
}
