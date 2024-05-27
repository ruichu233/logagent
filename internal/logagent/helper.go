package logagent

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(path.Dir(filename)))
	viper.AddConfigPath(root + "/configs")
	viper.SetConfigName("logagent")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("load config failed,err:%v\n", err)
		panic(err)
	}
	logrus.Infof("conf file is %s\n", filepath.Join(viper.ConfigFileUsed()))
}

func GetOptions() (*Options, error) {
	opts := NewOptions()
	if err := viper.Unmarshal(opts); err != nil {
		return nil, err
	}
	return opts, nil
}

func getIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		logrus.Warnf("get ip failed,err:%v\n", err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr).IP.String()

	addr := strings.Split(localAddr, ":")[0]
	return addr
}

func addConf(etcd *clientv3.Client) {
	_, _ = etcd.Put(context.Background(), "10.105.29.225", "[{\"path\":\"E:\\\\test.log\",\"topic\":\"web_log\"}]")
}
