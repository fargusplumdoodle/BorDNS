
# Kubernetes Config for BorDNS
Isaac Thiessen

Sat Apr 11 17:28:38 PDT 2020


### Contents
- [Description](##Description)
- [Generate secrets/configs](##Generate secrets/configs)
- [Simple setup](##Simple setup)

---------------------------------------------------


## Description

Simple deployment for BorDNS, CoreDNS, and ETCD. Helm chart is probably on the way

## Defaults
2 external services are created with type LoadBalancer:
- 10.0.1.253:53 CoreDNS DNS server
- 10.0.1.252:80 BorDNS HTTP server (make requests here)

## Generate secrets/configs

Run `./generate_config.py <path to bordns config>`
to generate config maps for CoreDNS and BorDNS.

The output will be
- Corefile: `./conf/bordns-secrets.yml`

## Simple setup

1. Generate your secrets
2. Create namespace
`kubectl create ns bordns`
3. Apply configuration
`kubectl apply -f ./conf/`
4. Apply objects
`kubectl apply -f ./objects`

Everything will be in the `bordns` namespace
