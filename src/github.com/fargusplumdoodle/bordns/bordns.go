package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

/*
    BorDNS
-------------

For setting up and configuring the
BorDNS API
*/

// config, yaml config must match this
type config struct {
	EtcdHosts  []string `yaml:"etcd_hosts"`
	ListenAddr string   `yaml:"listen_address"`
}

const (
	CONFIG_FILE = "config.yml"
)

func setupDB(etcdHosts []string) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdHosts,
		DialTimeout: 5 * time.Second,
		//DialKeepAliveTime:    0,
		//DialKeepAliveTimeout: 0,
		//MaxCallSendMsgSize:   0,
		//MaxCallRecvMsgSize:   0,
		//TLS:                  nil,
		//Username:             "",
		//Password:             "",
		//RejectOldCluster:     false,
		//DialOptions:          nil,
		//LogConfig:            nil,
		//Context:              nil,
		//PermitWithoutStream:  false,
	})
	if err != nil {
		panic(fmt.Errorf("unable to connect to etcd: %q", err))
	}
	defer cli.Close()
	return cli
}

func getConfig() *config {
	var conf config

	// reading config
	ymlfl, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		panic(fmt.Errorf("unable to read conf: %q", CONFIG_FILE))
	}
	err = yaml.Unmarshal(ymlfl, &conf)
	if err != nil {
		panic(fmt.Errorf("invalid conf: %q", CONFIG_FILE))
	}

	return &conf
}
