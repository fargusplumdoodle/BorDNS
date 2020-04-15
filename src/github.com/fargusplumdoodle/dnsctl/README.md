# dnsctl 

( DNS control )

For controlling BorDNS from the command line

### Usage

`dnsctl <command> <args>`

Commands:
- `dnsctl get all`
  - returns all A records in database
- `dnsctl get example.bor`
  - returns ip address of `example.bor`
- `dnsctl set example.bor 192.168.0.1`
  - creates an A record for `example.bor` with the IP of `192.168.0.1`
- `dnsctl del example.bor`
  - deletes A record for `example.bor` 
- `sudo dnsctl generate-config`
  - Creates example config file in /etc/bordns/client_conf.yml
- `dnsctl help`
  - prints commands

### Configuration
Config file location:
    /etc/bordns/client_conf.yml
   
Example config:
```yaml
auth_username: dev
auth_password: dev
bordns_host: http://localhost:8000
```

To generate the above configuration, run:

`sudo dnsctl generate-config`

Sudo privileges are required to write in `/etc`
### Example output

#### dnsctl get all
```
$ dnsctl get all
zone: bor
example.bor             10.0.0.1
web.site.bor            10.0.2.1

zone: sekhnet.ra 
example.sekhnet.ra      10.2.2.8
dns.sekhnet.ra          192.168.0.1
```
  
#### dnsctl get example.bor
```
$ dnsctl get example.bor
example.bor 192.168.0.45
```
When a domain isn't found
```
$ dnsctl get not.found.domain
not found
```

#### dnsctl set example.bor 192.168.0.1
```
$ dnsctl set example.bor 192.168.0.1
ok
```
#### dnsctl del example.bor 
```
$ dnsctl del example.bor 
boop... gone
```

#### sudo dnsctl generate-config
```
$ sudo dnsctl generate-config
generated config in: /etc/bordns/config.yml
```
```
$ dnsctl generate-config
error: requires elevated privileges to generate config
```
#### help
```
$ dnsctl help
dnsctl commands:
    dnsctl get all
      - returns all A records in database
    dnsctl get example.bor
      - returns ip address of "example.bor"
    dnsctl set example.bor 192.168.0.1
      - creates an A record for "example.bor" with the IP of "192.168.0.1"
    dnsctl del example.bor
      - deletes A record for "example.bor"
    dnsctl help
      - prints help message
    sudo dnsctl generate-config
      - Creates example config file in /etc/bordns/client_conf.yml
```
#### Error occured
If something goes wrong, or invalid arguments were supplied
The help message will be displayed with the specific error
after
