package model

import (
	"context"
	"errors"
	"github.com/fargusplumdoodle/bordns/conf"
	"strings"
)

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
