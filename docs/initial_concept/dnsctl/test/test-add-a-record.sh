#!/bin/bash

# this is a test for add-a-record.
# it assumes the "ra" zone exists

ZONE=ra
FQDN=this.site.ra
IP=10.0.0.1

COREDNS_IP=10.0.0.53
ETCD_IP=10.0.0.2

echo running add-a-record script
python3 $(pwd)/add-a-record $ZONE $FQDN $IP > /dev/null

# testing object was created
if [ $(etcdctl get /ra/ra/site/this  --endpoints=$ETCD_IP:2379 | wc -l) -ne 2 ];
then
  echo "TEST FAILED: OBJECT WAS NOT CREATED"
  exit -1
fi

# testing dns name was created
if [ $( dig @$COREDNS_IP +short $FQDN | wc -l) -ne 1 ];
then
  echo "TEST FAILED: OBJECT WAS NOT CREATED"
  exit -1
fi


