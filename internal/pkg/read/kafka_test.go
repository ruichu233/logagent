package read

import (
	"context"
	"github.com/ruichu233/logagent/pgk/options"
	"log"
	"testing"
)

func TestReader(t *testing.T) {
	o := options.NewKafkaOptions()
	o.Brokers = []string{"127.0.0.1:9091"}
	o.Topic = "web_log"
	o.ReaderOptions.GroupID = "web_log_consumer"
	reader := NewReader(o)
	ch := make(chan interface{}, 100)
	go reader.Read(context.Background(), ch)

	go func() {
		for {
			v := <-ch
			log.Println(v)
		}
	}()

	select {}
}
