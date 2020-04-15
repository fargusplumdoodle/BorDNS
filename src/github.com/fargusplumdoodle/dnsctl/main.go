package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

/*
DNSCTL
-------
All info is in the README
*/

const (
	WHITE_SPACE = 20
	CONF_FILE      = "/etc/bordns/client_conf.yml"
	CMD_ALL        = "all"
	CMD_GET        = "get"
	CMD_SET        = "set"
	CMD_DEL        = "del"
	CMD_HELP       = "help"
	CMD_GEN_CONFIG = "generate-config"
	HELP_MSG       = `
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
    - Creates example config file in /etc/bordns/client_conf.yaml`
)

var (
	conf Conf
)

/*
Procedure:
	1. Read configuration
	2. Determine command
	3. Run command
*/
func main() {
	// validating input
	if len(os.Args) <= 1 {
		Fail("invalid arguments")
	}
	// Read conf
	conf = ReadConf(CONF_FILE)

	// Determining command
	determineCommand()
}

/*
Determine Command
-----------------

Runs through arguments and
determines which command the user
has requested.
*/
func determineCommand() {
	switch os.Args[1] {
	case CMD_GET:
		if os.Args[2] == CMD_ALL {
			GetAll()
		} else {
			Get()
		}
	case CMD_SET:
		Set()
	case CMD_DEL:
		Del()
	case CMD_GEN_CONFIG:
		GenerateConf()
	case CMD_HELP:
		Help()
	default:
		Fail("argument not recognized")
	}

}

func Fail(errMsg string) {
	/*
		Prints help, then the error message
	*/
	Help()
	fmt.Println("\nerror:", errMsg)
	os.Exit(1)
}
func Help() {
	fmt.Println(HELP_MSG)
}

func ReadConf(confFile string) Conf {
	var conf Conf

	// checking /etc/bordns directory exists
	if _, err := os.Stat("/etc/bordns"); os.IsNotExist(err) {
		Fail("/etc/bordns does not exist! You may have to run \n   sudo dnsctl generate-config\n")
	}

	ymlfl, err := ioutil.ReadFile(confFile)
	if err != nil {
		Fail(fmt.Sprintf("Unable to read conf: %q", confFile))
	}
	err = yaml.Unmarshal(ymlfl, &conf)
	if err != nil {
		Fail(fmt.Sprintf("Invalid configuration: %q", confFile))
	}
	return conf
}

type Conf struct {
	Username   string `yaml:"auth_username"`
	Password   string `yaml:"auth_password"`
	BorDNSHost string `yaml:"bordns_host"`
}
type Arecord struct {
	IP   string `json:"ip"`
	FQDN string `json:"fqdn"`
}

type Zone struct {
	Name    string    `json:"zone"`
	Domains []Arecord `json:"domains"`
}

func GetAll() {
	/*
		Shows all known dns names and their values

		Procedure:
			1. Make HTTP GET request to bordns/domain
			2. Load input
			2. For each zones, print all dns names
	*/
	resp := MakeRequest(http.MethodGet, "domain")

	var zones []Zone
	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(&zones)
	if err != nil {
		Fail("invalid response from bordns" + err.Error())
	}

	for _, zone := range zones {
		fmt.Println("zones: " + zone.Name)

		for _, domain := range zone.Domains {
			space_amount := WHITE_SPACE - len(domain.FQDN)
			space := ""
			for i := 0; i < space_amount; i++ {
				space = space + " "
			}
			fmt.Println(domain.FQDN, space, domain.IP)
		}
		fmt.Println()
	}
}

func MakeRequest(method, uri string) *http.Response {
	/*
		Make request to BorDNS

		Args:
			method: the HTTP method to make the request with
		    uri: the URI to make, with no leading forward slash

		Procedure:
			1. Get URL
			2. Make request
			3. set credentials
			4. Make request
	*/
	// 1. Get URL
	url := GetURLWithTrailingSlash(conf.BorDNSHost) + uri

	// 2. Make Request
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		Fail("failed to make request to bordns: " + err.Error())
	}

	// 3. Set Credentials
	req.SetBasicAuth(conf.Username, conf.Password)

	// 4. Make request
	resp, err := client.Do(req)
	if err != nil {
		Fail("failed to make request to bordns: " + err.Error())
	}

	return resp
}

func Get() {
	fmt.Println("get")
}
func Set() {
	fmt.Println("set")
}
func Del() {
	fmt.Println("del")
}
func GenerateConf() {
	fmt.Println("generate conf")
}

func GetURLWithTrailingSlash(url string) string {
	/*
		Takes a given URL, if it doesn't end in a slash
		we add the slash
	*/
	// if the last character is a slash, return
	if url[len(url)-1] == []byte("/")[0] {
		return url
	} else {
		return url + "/"
	}

}
