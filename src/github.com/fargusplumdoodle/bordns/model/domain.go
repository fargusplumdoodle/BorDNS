package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
	"github.com/fargusplumdoodle/bordns/viewmodel"
	"go.etcd.io/etcd/clientv3"
)

func GetAllDomains() ([]viewmodel.Zone, error) {
	zoneMap := map[string][]viewmodel.Arecord{}

	for _, zone := range conf.Env.Zones {
		zoneMap[zone.Zone] = []viewmodel.Arecord{}

		ctx, cancel := context.WithTimeout(context.Background(), conf.DB_TIMEOUT)

		// unfortunately I have had a hard time figuring out how to get ETCD to return all children.
		// this seems to work, but I would prefer a better way. Basically it will keep recursing until
		// either it reaches all children, or finds this specific key. Since this key doesn't exist
		// it will return all children
		resp, err := client.Get(ctx, zone.EtcdPath, clientv3.WithRange(zone.EtcdPath+"/this/key/is/highly/unlikely/to/exist"))
		cancel()

		if err != nil {
			return nil, err
		}

		fmt.Println("# keys: ", len(resp.Kvs))
		for _, ev := range resp.Kvs {

			// getting attributes from etcd. This will contain  {host: ip, ttl: 60}
			aRecord := CoreDNSARecord{}
			err := json.Unmarshal(ev.Value, &aRecord)
			if err != nil {
				return nil, err
			}

			zoneMap[zone.Zone] = append(zoneMap[zone.Zone], viewmodel.Arecord{
				IP:   aRecord.Host,
				FQDN: GetHostnameFromEtcdPath(zone.EtcdPath, string(ev.Key)),
			})
		}
	}

	// Creating list of zones
	result := []viewmodel.Zone{}
	for k, v := range zoneMap {
		zone := viewmodel.Zone{Name: k, Domains: []viewmodel.Arecord{}}

		// looping over records to add to zone
		for _, record := range v {
			zone.Domains = append(zone.Domains, record)
		}

		result = append(result, zone)
	}
	return result, nil
}
