#!/usr/bin/python3
"""
 Generates 2 config maps for BorDNS
    1. Corefile for CoreDNS
    2. BorDNS config

Note:
    - The etcd_hosts will always be set to localhost
    - If there is an issue with your initial BorDNS
      configuration, this likely will not detect or fix it
    - Secrets will both be for the "bordns" namespace.
    - If you take issue with any of this, change the
      constants below

Procedure:
    1. Ensure argument was supplied to a valid file, and load yaml file
    2. Generate corefile based on zones specified in bordns config
    3. Add corefile secret conf to secret file
    4. Set the ETCD host to go to localhost (etcd will be in the same pod as coredns)
    5. Add BorDNS secret to secret file
    6. Write output to ./conf/bordns-secrets.yml

"""
import yaml
import os
import sys
import base64

try:
    from yaml import CLoader as Loader, CDumper as Dumper
except ImportError:
    from yaml import Loader, Dumper

# 1.
if len(sys.argv) != 2 or not os.path.isfile(sys.argv[1]):
    print("Invalid arguments: please supply path to valid BorDNS config.yml")
    exit(1)

# ---------------------------------------------
#               CONSTANTS
# ---------------------------------------------
CONF_INPUT = sys.argv[1]
OVERRIDE_ETCD_HOST = '127.0.0.1:2379'
NAMESPACE = "bordns"
OUTPUT_FILE = "./conf/bordns-secrets.yml"
COREFILE_HEADER = """
.:53 {
        log
        errors
}
"""
# ---------------------------------------------


def get_zone_section(zone, path):
    return (
            """
    """
            + zone
            + """ {
        log
        errors
        etcd  {
                path """
            + path
            + """ 
                endpoint http://127.0.0.1:2379
        }
    }"""
    )


def generate_secret(conf_file_name, secret_name, conf_file_content):
    """
    :param conf_file_name: the file name to store in the config map
    :param secret_name: name of the kubernetes secret
    :param conf_file_content: text data base64 encoded
    :return: text of secret conf
    """
    return f"""
apiVersion: v1
kind: Secret
metadata:
  namespace: {NAMESPACE} 
  name: {secret_name}
type: Opaque
data:
  {conf_file_name}: {str(base64.b64encode(conf_file_content.encode("ascii")))[2:-1]}
    """


secret_file_content = ""
with open(CONF_INPUT, "r") as yml_fl:
    try:
        conf = yaml.load(yml_fl, Loader=Loader)
    except Exception as e:
        print(f"Invalid config '{CONF_INPUT}': {e}")
        exit(-2)

    # 2. Generate core file secret
    corefile_conf = COREFILE_HEADER
    for zone in conf["zones"]:
        corefile_conf += get_zone_section(zone["zone"], zone["path"])

    # 3. Add corefile secret to secret file
    secret_file_content += generate_secret("Corefile", "corefile-secret", corefile_conf) + "\n---\n"

    # 4. setting ectd host
    conf['etcd_hosts'] = [OVERRIDE_ETCD_HOST]

    # 5. Add bordns secret to secret file
    secret_file_content += generate_secret("config.yml", "bordns-secret", yaml.dump(conf)) + "\n"

# 6. write kubernetes secret to file
with open(OUTPUT_FILE, 'w') as fl:
    fl.write(secret_file_content)
