package read

import (
	"context"
	"encoding/json"
	"github.com/ruichu233/logagent/pgk/options"
	"github.com/segmentio/kafka-go"
)

type Reader struct {
	kr *kafka.Reader
}

func NewReader(opts *options.KafkaOptions) *Reader {
	rc := kafka.ReaderConfig{
		Brokers: opts.Brokers,

		Topic:          opts.Topic,
		GroupID:        opts.ReaderOptions.GroupID,
		MaxBytes:       opts.ReaderOptions.MaxBytes,
		CommitInterval: opts.ReaderOptions.CommitInterval,
		StartOffset:    opts.ReaderOptions.StartOffset,
	}
	r := kafka.NewReader(rc)
	return &Reader{
		kr: r,
	}
}

func (r *Reader) Read(ctx context.Context, ch chan interface{}) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			message, err := r.kr.ReadMessage(ctx)
			if err != nil {
				return
			}
			var mp map[string]interface{}
			if err := json.Unmarshal(message.Value, &mp); err != nil {
				return
			}
			ch <- mp
		}
	}
}

func (r *Reader) Close() {
	_ = r.kr.Close()
}
