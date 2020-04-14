package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

/*
DNSCTL
-------
All info is in the README
 */

const (
	CONF_FILE = "/etc/bordns/client_conf.yaml"
	CMD_ALL	= "all"
	CMD_GET = "get"
	CMD_SET = "set"
	CMD_DEL = "del"
	CMD_HELP = "help"
	CMD_GEN_CONFIG = "generate-config"
	HELP_MSG = `
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

func GetAll() {
	fmt.Println("getting all: ", conf.Username)
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
