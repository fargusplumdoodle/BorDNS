package conf

/*
    BorDNS
-------------

For setting up and configuring the
BorDNS API
*/

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

const (
	DB_TIMEOUT  = 5 * time.Second
	DEFAULT_TTL = 60
)

var (
	Env        *Config
	configFile = os.Getenv("CONFIG")
)

// config, yaml config must match this
type Config struct {
	EtcdHosts    []string     `yaml:"etcd_hosts"`
	ListenAddr   string       `yaml:"listen_address"`
	Zones        []ZoneConfig `yaml:"zones"`
	AuthUsername string       `yaml:"auth_username"`
	AuthPassword string       `yaml:"auth_password"`
}

type ZoneConfig struct {
	Zone     string `yaml:"zone"`
	EtcdPath string `yaml:"path"`
}

func SetupConfig() {
	var conf Config

	// reading config
	ymlfl, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("unable to read conf: %q", configFile))
	}
	err = yaml.Unmarshal(ymlfl, &conf)
	if err != nil {
		panic(fmt.Errorf("invalid conf: %q", configFile))
	}

	Env = &conf
}
