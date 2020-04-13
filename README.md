
# BorDNS
Isaac Thiessen

Sun Apr 12 09:48:06 PDT 2020


### Contents
- [Description](##Description)
- [Configuring](##Configuring)
- [Kubernetes](##Kubernetes)
- [Docker and docker-compose](##Docker and docker-compose)

---------------------------------------------------


## Description
  Boreal DNS (BorDNS). A REST API management system for CoreDNS.

## Configuring
Environment Variable Required: 

  `CONFIG = (path to config)`

Config example in src/github.com/fargusplumdoodle/bordns/config.yml

## Kubernetes:

Example configuration in `./kubernetes`. Instructions in `./kubernetes/README.md`

## Docker and docker-compose
An example script for running the BorDNS API alone: `scripts/docker_run_bordns_test.sh`

But the BorDNS API is fairly useless without etcd and CoreDNS. 
You can start the 3 of them with `docker-compose up`
from the root of this repository

