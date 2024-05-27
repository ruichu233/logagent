package logagent

import (
	"context"
	"github.com/magiconair/properties/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
)

func TestOptions(t *testing.T) {
	opts, err := GetOptions()
	if err != nil {
		return
	}
	assert.Equal(t, opts.KafkaOptions.WriterOptions.MaxAttempts, 3)
}

func TestGetIp(t *testing.T) {
	ip := getIp()
	assert.Equal(t, ip, "10.105.29.225")
}

func TestAddConf(t *testing.T) {
	etcd, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:20000", "127.0.0.1:20002", "127.0.0.1:20004"},
	})
	addConf(etcd)
	get, _ := etcd.Get(context.Background(), "10.105.29.225")
	bytess := get.Kvs[0].Value
	s := string(bytess)
	assert.Equal(t, s, "[{\"path\":\"E:\\\\test.log\",\"topic\":\"web_log\"}]")
}
