
# BorDNS
Isaac Thiessen

Sun Apr 12 09:48:06 PDT 2020


### Contents
- [Description](##Description)
- [Configuring](##Configuring)
- [TODO](##TODO)

---------------------------------------------------


## Description
  Boreal DNS (BorDNS). A REST API bolted onto CoreDNS and etcd.   

## Configuring
Environment Variable Required: 

  `CONFIG = (path to config)`

Config example in src/github.com/fargusplumdoodle/bordns/config.yml

## TODO:
- Script that will generate simple Corefile from config.yml zones
- Helm chart

## Docker 
An example script for running the BorDNS API alone: `scripts/docker_run_bordns_test.sh`

But the BorDNS API is fairly useless without etcd and CoreDNS. 
You can start the 3 of them with `docker-compose up`
from the root of this repository

