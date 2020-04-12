package conf

/*
    BorDNS
-------------

For setting up and configuring the
BorDNS API
*/

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

const (
	COREFILE    = "Corefile"
	CONFIG_FILE = "config.yml"
	DB_TIMEOUT  = 5 * time.Second
)

// config, yaml config must match this
type Config struct {
	EtcdHosts  []string     `yaml:"etcd_hosts"`
	ListenAddr string       `yaml:"listen_address"`
	Zones      []ZoneConfig `yaml:"zones"`
}

type ZoneConfig struct {
	Zone     string `yaml:"zone"`
	EtcdPath string `yaml:"path"`
}

func SetupDB(etcdHosts []string) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdHosts,
		DialTimeout: DB_TIMEOUT,
		// TODO: AUTHENTICATION
		//Username:             "",
		//Password:             "",
	})
	if err != nil {
		panic(fmt.Errorf("unable to connect to etcd: %q", err))
	}
	defer cli.Close()
	return cli
}

func GetConfig() *Config {
	var conf Config

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
