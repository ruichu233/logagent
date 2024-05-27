package logagent

import (
	"context"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ConfCentre struct {
	ctx        context.Context
	key        string
	updateChan chan struct{}
	etcd       *clientv3.Client
	collects   []*collectEntry
}

func NewConfCentre(updateChan chan struct{}, etcd *clientv3.Client) (*ConfCentre, error) {
	c := &ConfCentre{
		ctx:        context.Background(),
		key:        getIp(),
		updateChan: updateChan,
		etcd:       etcd,
	}
	if err := c.GetCollectsFromEtcd(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ConfCentre) GetCollectsFromEtcd() error {
	resp, err := c.etcd.Get(c.ctx, c.key)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resp.Kvs[0].Value, &c.collects); err != nil {
		return err
	}
	return nil
}

func (c *ConfCentre) watchConf() error {
	for watchCh := range c.etcd.Watch(c.ctx, c.key) {
		if watchCh.Err() != nil {
			return watchCh.Err()
		}
		if err := json.Unmarshal(watchCh.Events[0].Kv.Value, &c.collects); err != nil {
			return err
		}
		c.updateChan <- struct{}{}
	}
	return nil
}
