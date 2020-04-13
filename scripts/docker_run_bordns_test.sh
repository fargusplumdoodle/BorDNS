#!/bin/bash
docker run -it \
	--rm \
	--net bordns_dnsnet \
	-p 8000:8000 \
	-v $(pwd)/src/github.com/fargusplumdoodle/bordns/config.yml:/config.yml \
	bordns:0.1
