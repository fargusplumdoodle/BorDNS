package model

import (
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
	"go.etcd.io/etcd/clientv3"
	"strings"
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

/*
Determine if host is in existing zone
---------------

Looks through all of the zones from the config and
finds the zone that matches the provided host.

Returns an error if none were found.

In the event that two zones are similar, we will match the
zone that is the longest.
e.g:
Zones: zone.bor, .bor
Input: test.zone.bor
Match: zone.bor

Procedure:
	1. Convert all inputs to reversed slices, seperated by "."
*/
func getZoneFromHost(host string) (conf.ZoneConfig, error) {
	zc := conf.ZoneConfig{}

	return zc, nil
}

/*
Converts a domain into a slice  without the "."
then reverses the order

Example:
  in: "longer.test.bor"
  out: []string{"bor", "test", "longer"}

Procedure:
	1. Convert to slice
	2. Reverse order of slice
*/
func getReversedDomain(host string) []string {
	split := strings.Split(host, ".")
	return reverseSlice(split)
}

func reverseSlice(a []string) []string {
	/*
		Reverses the order of the elements of a slice:
		Credit: https://github.com/golang/go/wiki/SliceTricks
	*/
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
