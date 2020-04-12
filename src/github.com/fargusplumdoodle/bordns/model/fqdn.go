package model

import (
	"context"
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
)

func AddARecord(host string, ip string) error {
	/*
		Adds a record to the ETCD database
		Procedure:
		---------
		1. Determine if host is in any of the existing zones
		2. Convert host into etcd path, swap "." for "/" then split by "/" then reverse order. Then
		   prepend the appropriate zone
		3. Add the
	*/
	ctx, cancel := context.WithTimeout(context.Background(), conf.DB_TIMEOUT)
	resp, err := client.Put(ctx, "/bor/bor/test", `{"host":"10.0.0.1","ttl":60}`)
	cancel()
	if err != nil {
		fmt.Errorf("Error adding dns name to etcd, %q", err)
	}
	fmt.Printf("response:, %q", resp)

	return nil
}
