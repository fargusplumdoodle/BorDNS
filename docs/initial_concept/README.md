# Simple DNS Server
------------

**NOTE** THIS IS NOT A PRODUCTION DNS SERVER
	 AUTHENTICATION IS NOT CONFIGURED FOR
	 THE ETCD DATABASE 

This is a simple DNS server using CoreDNS and etcd. It is configured and
installed with the "bordns" helmchart. DNS records can be added with the 
"dnsctl" tool.

CoreDNS is the standard Kubernetes DNS server.
etcd is the standard Kubernetes database.

This dns server only supports a single zone at this time.

CoreDNS integrates very well with etcd, allowing to dynamicallyupdate A records, by adding them directly to the etcd database.


# Production setup

What this needs to be production ready:
	- TLS authentication for ETCD server
	- ETCD Cluster instead of single node
	- Add client authentication for scripts/add-a-record

# Setup: bordns

This is a guide for getting bordns on your Kubernetes cluster.

Requirements:
	- Kubernetes cluster
	- Helm installed locally
	- Decide the IP that you want the service to be available from 
	- Decide the zone name you want
	- Decide the namespace you want to install the bordns service to

Variables:
	serviceIP:
		 description: The external IP for the service
		 default:192.168.0.53 
	storageClass: 
		 description: the storage class for etcd
		 default: nfs4
	defaultZone: 
		 description: The DNS zone to create
		 default: bor 

Usage Examples:
	```bash
	helm install bordns ./bordns \
		--set serviceIP=192.168.0.53 \
		--set storageClass=nfs4 \
		--set defaultZone=bor \
		-n dns
	```

# Adding A records manually

This guide is for manually adding A records to
your dns server. It requires you to have 
etcdctl installed. 

Assuming the DNS server is configured like this:
  zone: bor
  bordns host: 10.0.0.2
And you want to add this A record
  fqdn: isaac.bor
  host ip: 192.168.0.24

```
# create A record
etcdctl put /bor/bor/isaac {host: 192.168.0.24, ttl: 60} --endpoints=10.0.0.2:2379

# validates to ensure it worked
dig +short isaac.bor @10.0.0.2
```

# dnsctl: adding A records from the CLI

I wrote the python script "./dnsctl/dnsctl" for making
this process easier.

First edit the conf at the top of "./dnsctl/dnsctl"
to match your enviroment.

Assuming the DNS server is configured like this:
  zone: bor
  etcd host: 10.0.0.2
  coredns host: 10.0.0.53
And you want to add this A record
  fqdn: isaac.bor
  host ip: 192.168.0.24

```bash
python3 ./scripts/add-a-record bor isaac.bor 192.168.0.24

# validates to ensure it worked
dig +short isaac.bor @10.0.0.53
```
