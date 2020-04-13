package model

import (
	"errors"
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
	"go.etcd.io/etcd/clientv3"
	"strings"
)

// etcd client
var client *clientv3.Client

func SetupDB(etcdHosts []string) *clientv3.Client {
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
	// TODO: FIGURE THIS OUT
	// defer cli.Close()

	client = cli

	return client
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


Explanation:
Creates lists of the hostnames and reverses them. So
zone.bor -> "bor", "zone"

then loops from start to finish through the known zones checking each
element on the provided host and the zone we are checking for.

As soon as we find an element that is in the zone we are checking for
but NOT in the provided host, we go to the next zone to check for

We take the zone that has the most elements in common with the provided
host. Or just return an error if we couldnt find it
*/
func getZoneFromHost(host string) (conf.ZoneConfig, error) {
	zc := conf.ZoneConfig{}

	// 1. getting reversed host
	reversedHost := getReversedDomain(host)
	var longestZone = 0

	// looping through recognized zones
	for _, zone := range conf.Env.Zones {
		// getting reversed zone
		reverseDomain := getReversedDomain(zone.Zone)
		var match = false

		// looping until we find two elements that dont match
		for i := 0; i < len(reverseDomain); i++ {
			if reverseDomain[i] != reversedHost[i] {
				match = false
				break
			} else {
				match = true
			}
		}
		// setting the zone if its the longest
		if match && len(reverseDomain) > longestZone {
			longestZone = len(reverseDomain)
			zc = zone
		}
	}

	// checking if any  zones were found
	if longestZone == 0 {
		// no zones were found, informing the user of the known zones because we are nice
		var knownZones = []string{}
		for _, zone := range conf.Env.Zones {
			knownZones = append(knownZones, zone.Zone)
		}
		return zc, errors.New(fmt.Sprintf(
			"host does not match any known zones, known zones: [%v]", strings.Join(knownZones, ", ")))
	}

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

/*
Performs the inverse operation of getReversedDomain

Example:
  in: []string{"bor", "test", "longer"}
  out: "longer.test.bor"
*/
func unReverseDomain(reversedDomain []string) string {
	unreversed := reverseSlice(reversedDomain)
	return strings.Join(unreversed, ".")
}

/*
	Reverses the order of the elements of a slice:
	Credit: https://github.com/golang/go/wiki/SliceTricks
*/
func reverseSlice(a []string) []string {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func getARecordValue(ip string) string {
	return fmt.Sprintf(`{"host":"%v","ttl":%v}`, ip, conf.DEFAULT_TTL)
}

type CoreDNSARecord struct {
	Host string `json:"host"` // host is the IP, thats just how CoreDNS does it
	Ttl  int    `json:"ttl"`
}

/*

 */
func GetHostnameFromEtcdPath(zone, path string) string {
	return path
}
