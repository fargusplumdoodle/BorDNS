package model

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/fargusplumdoodle/bordns/conf"
	"strings"
)

type CoreDNSARecord struct {
	Host string `json:"host"`
	Ttl  int    `json:"ttl"`
}

func AddARecord(host string, ip string) error {
	/*
		Adds a record to the ETCD database
		Procedure:
		---------
		1. Determine if host is in any of the existing zones
		2. Convert host into etcd path, swap "." for "/" then split by "/" then reverse order. Then
		   prepend the appropriate zone
			-> example if the path is /bor, and the host was test.bor. The output would be /bor/bor/test
		3. Add the A record to the databse
	*/

	// 1.
	zone, err := getZoneFromHost(host)
	if err != nil {
		return err
	}

	// 2. Converting to etcd path.
	reversedHostDomain := getReversedDomain(host)
	path := zone.EtcdPath + "/" + strings.Join(reversedHostDomain, "/")

	// 3. add to db
	ctx, cancel := context.WithTimeout(context.Background(), conf.DB_TIMEOUT)
	_, err = client.Put(ctx, path, getARecordValue(ip))
	cancel()
	if err != nil {
		return errors.New("error adding dns name to etcd, " + err.Error())
	}

	return nil
}

func GetCoreDNSRecordForHost(host string) (CoreDNSARecord, error) {
	/*
		Get CoreDNS record for Host
		-------------
		Looks in the ETCD database for the IP of the host

		1. Determine if the host is in any of the existing zones
		2. convert to etcd path for lookup
		3. Perform lookup on path
	*/
	// getting IP from CoreDNS A record
	var aRecord = CoreDNSARecord{}

	// 1.
	zone, err := getZoneFromHost(host)
	if err != nil {
		return aRecord, err
	}

	// 2. Converting to etcd path.
	reversedHostDomain := getReversedDomain(host)
	path := zone.EtcdPath + "/" + strings.Join(reversedHostDomain, "/")

	// 3.
	ctx, cancel := context.WithTimeout(context.Background(), conf.DB_TIMEOUT)
	resp, err := client.Get(ctx, path)
	cancel()

	if err != nil {
		return aRecord, errors.New("did not find IP of host " + host)
	}

	for _, ev := range resp.Kvs {

		err := json.Unmarshal(ev.Value, &aRecord)
		if err != nil {
			return aRecord, errors.New("data stored in invalid format: " + err.Error())
		}

		return aRecord, nil
	}
	return aRecord, errors.New("did not find IP of host " + host)
}
