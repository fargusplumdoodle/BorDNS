#!/bin/bash

docker network create dns --subnet 10.0.0.0/24 &> /dev/null

docker run --name dns \
	--rm -it \
	--ip 10.0.0.53 \
	--net dns \
	-v $(pwd)/coredns/:/root/ \
	coredns/coredns \
	-conf /root/Corefile
