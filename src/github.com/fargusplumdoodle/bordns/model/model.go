package model

import (
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
	"go.etcd.io/etcd/clientv3"
)

// etcd client
var client *clientv3.Client

func SetupDB(etcdHosts []string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdHosts,
		DialTimeout: conf.DB_TIMEOUT,
		// TODO: AUTHENTICATION
		//Username:             "",
		//Password:             "",
	})
	if err != nil {
		panic(fmt.Errorf("unable to connect to etcd: %q", err))
	}
	defer cli.Close()

	client = cli
}
