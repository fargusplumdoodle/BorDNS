
# Kubernetes Config for BorDNS
Isaac Thiessen

Sat Apr 11 17:28:38 PDT 2020


### Contents
- [Description](##Description)
- [Generate secrets/configs](##Generate secrets/configs)
- [Simple setup](##Simple setup)

---------------------------------------------------


## Description

Simple deployment for BorDNS, CoreDNS, and ETCD. 

## Generate secrets/configs

Run `./generate_config.py <path to bordns config>`
to generate config maps for CoreDNS and BorDNS.

The output will be 
- Corefile: `./conf/bordns-secrets.yml`

## Simple setup

1. Generate your secrets
2. Apply configuration
`kubectl apply -f ./conf`
3. Apply objects
`kubectl apply -f ./objects`

Everything will be in the `bordns` namespace
