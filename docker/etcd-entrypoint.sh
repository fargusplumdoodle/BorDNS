#!/bin/bash
TOKEN=token-01
CLUSTER_STATE=new
NAME_1=machine-1
HOST_1=0.0.0.0
CLUSTER=${NAME_1}=http://${HOST_1}:2380

etcd \
	--data-dir=data.etcd --name ${NAME_1} \
	--initial-advertise-peer-urls http://${HOST_1}:2380 --listen-peer-urls http://${HOST_1}:2380 \
	--advertise-client-urls http://${HOST_1}:2379 --listen-client-urls http://${HOST_1}:2379 \
	--initial-cluster ${CLUSTER} \
	--initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}


