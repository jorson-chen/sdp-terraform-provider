---
layout: "appgate"
page_title: "APPGATE: appgate_appliance"
sidebar_current: "docs-appgate-resource-appliance"
description: |-
   Create a new inactive Appliance.
---

# appgate_appliance

Create a new inactive Appliance..

## Example Usage

```hcl

data "appgate_site" "default_site" {
  site_name = "Default site"
}

resource "appgate_appliance" "new_gateway" {
  name     = "gateway-asd"
  hostname = "envy-10-97-168-1337.devops"

  client_interface {
    hostname       = "envy-10-97-168-1338.devops"
    proxy_protocol = true
    https_port     = 447
    dtls_port      = 445
    allow_sources {
      address = "1.3.3.8"
      netmask = 0
      nic     = "eth0"
    }
    override_spa_mode = "UDP-TCP"
  }

  peer_interface {
    hostname   = "envy-10-97-168-1338.devops"
    https_port = "1338"

    allow_sources {
      address = "1.3.3.8"
      netmask = 0
      nic     = "eth0"
    }
  }


  admin_interface {
    hostname = "envy-10-97-168-1337.devops"
    https_ciphers = [
      "ECDHE-RSA-AES256-GCM-SHA384",
      "ECDHE-RSA-AES128-GCM-SHA256"
    ]
    # allow_sources {
    #   address = "1.3.3.7"
    #   netmask = 0
    #   nic     = "eth0"
    # }
  }

  tags = [
    "terraform",
    "api-created"
  ]
  notes = "hello world"
  site  = data.appgate_site.default_site.id

  connection {
    type     = "ssh"
    user     = "cz"
    password = "cz"
    host     = "10.97.168.30"
  }

  networking {

    hosts {
      hostname = "bla"
      address  = "0.0.0.0"
    }
    hosts {
      hostname = "foo"
      address  = "127.0.0.1"
    }

    nics {
      enabled = true
      name    = "eth0"
      ipv4 {
        dhcp {
          enabled = true
          dns     = true
          routers = true
          ntp     = true
        }
      }
    }
    dns_servers = [
      "8.8.8.8",
      "1.1.1.1",
    ]
    dns_domains = [
      "aa.com"
    ]
    routes {
      address = "0.0.0.0"
      netmask = 24
      gateway = "1.2.3.4"
      nic     = "eth0"
    }
  }


  ntp {
    servers {
      hostname = "ntp.microsoft.com"
      key_type = "MD5"
      key      = "bla"
    }
    servers {
      hostname = "ntp.google.com"
      key_type = "MD5"
      key      = "bla"
    }
  }

  ssh_server {
    enabled                 = true
    port                    = 2222
    password_authentication = true
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
    # allow_sources {
    #   address = "0.0.0.0"
    #   netmask = 0
    #   nic     = "eth1"
    # }
  }

  snmp_server {
    enabled    = false
    tcp_port   = 161
    udp_port   = 161
    snmpd_conf = "foo"
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }

  healthcheck_server {
    enabled = true
    port    = 5555
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }
  prometheus_exporter {
    enabled = true
    port    = 1234
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }

  ping {
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }

  log_forwarder {
    enabled = true
    elasticsearch {
      url                      = "https://aws.com/elasticsearch/instance/asdaxllkmda64"
      aws_id                   = "string"
      aws_region               = "eu-west-2"
      use_instance_credentials = true
      retention_days           = 3
    }

    tcp_clients {
      name    = "Company SIEM"
      host    = "siem.company.com"
      port    = 8888
      format  = "json"
      use_tls = true
    }
    sites = [
      data.appgate_site.default_site.id
    ]
  }

  iot_connector {
    enabled = true
    clients {
      name      = "Printers"
      device_id = "12699e27-b584-464a-81ee-5b4784b6d425"

      sources {
        address = "1.3.3.7"
        netmask = 24
        nic     = "eth0"
      }
      snat = true
    }
  }

  rsyslog_destinations {
    selector    = "*.*"
    template    = "%HOSTNAME% %msg%"
    destination = "10.10.10.2"
  }
  rsyslog_destinations {
    selector    = ":msg, contains, \"[AUDIT]\""
    template    = "%msg:9:$%"
    destination = "10.30.20.3"
  }

  hostname_aliases = [
    "appgatealias.company.com",
    "alias2.appgate.company.com"
  ]

  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=controller-a
  controller {
    enabled = true
  }

  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=logserver-a
  log_server {
    enabled = false
    # retention_days = 2
  }
  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=gateway-a
  # gateway {
  #   enabled = true
  #   vpn {
  #     weight = 60
  #     allow_destinations {
  #       address = "127.0.0.1"
  #       netmask = 0
  #       nic     = "eth0"
  #     }
  #   }
  # }



}

```

## Argument Reference

The following arguments are supported:


* `activated`: (Optional) Whether the Appliance is activated or not. If it is not activated, it won't be accessible by the Clients.
* `pending_certificate_renewal`: (Optional) Whether the Appliance is pending certificate renewal or not. Should be true for a very short period on certificate renewal.
* `version`: (Optional) Peer version of the Appliance.
* `hostname`: (Required) Generic hostname of the appliance. Used as linux hostname and to identify within logs.
* `site`: (Optional) Site served by the Appliance. Entitlements on this Site will be included in the Entitlement Token for this Appliance. Not useful if Gateway role is not enabled.
* `site_name`: (Optional) Name of the Site for this Appliance. For convenience only.
* `customization`: (Optional) Customization assigned to this Appliance.
* `connect_to_peers_using_client_port_with_spa`: (Optional) Makes the Appliance to connect to Controller/LogServer/LogForwarders using their clientInterface.httpsPort instead of peerInterface.httpsPort. The Appliance uses SPA to connect.
* `client_interface`: (Required) The details of the Client connection interface.
* `peer_interface`: (Required) The details of peer connection interface. Used by other appliances and administrative UI.
* `admin_interface`: (Optional) The details of the admin connection interface. If null, admin interface will be accessible via peerInterface.
* `networking`: (Required) Networking configuration of the system.
* `ntp_servers`: (Optional) Deprecated as of 4.3.0, use 'ntp' field instead. NTP servers to synchronize time.
* `ntp`: (Optional) NTP configuration.
* `ssh_server`: (Optional) SSH server configuration.
* `snmp_server`: (Optional) SNMP Server configuration.
* `healthcheck_server`: (Optional) Healthcheck Server configuration.
* `prometheus_exporter`: (Optional) Prometheus Exporter configuration.
* `ping`: (Optional) Rules for allowing ping.
* `log_server`: (Optional) Log Server settings. Log Server collects audit logs from all the appliances and stores them.
* `controller`: (Optional) Controller settings.
* `gateway`: (Optional) Gateway settings.
* `log_forwarder`: (Optional) LogForwarder settings. LogForwarder collects audit logs from the appliances in the given sites and sends them to the given endpoints.
* `connector`: (Optional) Connector settings.
* `rsyslog_destinations`: (Optional) Rsyslog destination settings to forward appliance logs.
* `hostname_aliases`: (Optional) Hostname aliases. They are added to the Appliance certificate as Subject Alternative Names so it is trusted using different IPs or hostnames. Requires manual certificate renewal to apply changes to the certificate.
* `id`: (Required) ID of the object.
* `name`: (Required) Name of the object.
* `notes`: (Optional) Notes for the object. Used for documentation purposes.
* `tags`: (Optional) Array of tags.


### client_interface
The details of the Client connection interface.

* `proxy_protocol`:  (Optional)  default value `false` To enable/disable Proxy protocol on this Appliance.
* `hostname`: (Required) Hostname to connect by the Clients. It will be used to validate the Appliance Certificate. Example: appgate.company.com.
* `https_port`:  (Optional)  default value `443` Port to connect for the Client specific services.
* `dtls_port`:  (Optional)  default value `443` Port to connect for the Clients that connects to vpnd on DTLS if enabled.
* `allow_sources`:  (Optional) Source configuration to allow via iptables.
* `override_spa_mode`:  (Optional)  Enum values: `Disabled,TCP,UDP-TCP`Override SPA mode for this appliance.
#### allow_sources
Source configuration to allow via iptables.
* `address`: (Optional) IP address to allow connection. Example: 0.0.0.0,::.
* `netmask`: (Optional) Netmask to use with address for allowing connections. Example: 0.
* `nic`: (Optional) NIC name to accept connections on. Example: eth0.
### peer_interface
The details of peer connection interface. Used by other appliances and administrative UI.

* `hostname`: (Required) Hostname to connect by the peers. It will be used to validate the appliance certificate. Example: appgate.company.com.
* `https_port`:  (Optional)  default value `444` Port to connect for peer specific services.
* `allow_sources`:  (Optional) Source configuration to allow via iptables.
#### allow_sources
Source configuration to allow via iptables.
* `address`: (Optional) IP address to allow connection. Example: 0.0.0.0,::.
* `netmask`: (Optional) Netmask to use with address for allowing connections. Example: 0.
* `nic`: (Optional) NIC name to accept connections on. Example: eth0.
### admin_interface
The details of the admin connection interface. If null, admin interface will be accessible via peerInterface.

* `hostname`: (Required) Hostname to connect to the admin interface. This hostname will be used to validate the appliance certificate. Example: appgate.company.com.
* `https_port`:  (Optional)  default value `8443` Port to connect for admin services.
* `https_ciphers`: (Required)  default value `ECDHE-RSA-AES256-GCM-SHA384,ECDHE-RSA-AES128-GCM-SHA256` The type of TLS ciphers to allow. See: https://www.openssl.org/docs/man1.0.2/apps/ciphers.html for all supported ciphers.
* `allow_sources`:  (Optional) Source configuration to allow via iptables.
#### https_ciphers
The type of TLS ciphers to allow. See: https:&#x2F;&#x2F;www.openssl.org&#x2F;docs&#x2F;man1.0.2&#x2F;apps&#x2F;ciphers.html for all supported ciphers.
#### allow_sources
Source configuration to allow via iptables.
* `address`: (Optional) IP address to allow connection. Example: 0.0.0.0,::.
* `netmask`: (Optional) Netmask to use with address for allowing connections. Example: 0.
* `nic`: (Optional) NIC name to accept connections on. Example: eth0.
### networking
Networking configuration of the system.

* `hosts`:  (Optional) /etc/hosts configuration
* `nics`:  (Optional) System NIC configuration
* `dns_servers`:  (Optional) DNS Server addresses. Example: 172.17.18.19,192.100.111.31.
* `dns_domains`:  (Optional) DNS Search domains. Example: internal.company.com.
* `routes`:  (Optional) System route settings.
#### hosts
&#x2F;etc&#x2F;hosts configuration
* `hostname`: (Required) Hostname to map IP to. Example: internal.service.company.com.
* `address`: (Required) IP for the given hostname for appliance to resolve. Example: 10.10.10.10.
#### nics
System NIC configuration
* `enabled`: (Optional) Whether the NIC is active or not. Example: true.
* `name`: (Required) NIC name Example: eth0.
* `ipv4`: (Optional) IPv4 settings for this NIC.
* `ipv6`: (Optional) IPv6 settings for this NIC.
* `mtu`: (Optional) MTU setting for the NIC. If left empty, appliance default will be used. Example: 1500.
##### dhcp
IPv4 DHCP configuration for the NIC.
* `enabled`: Whether DHCP for IPv4 is enabled.
* `dns`: Whether to use DHCP for setting IPv4 DNS settings on the appliance.
* `routers`: Whether to use DHCP for setting IPv4 default gateway on the appliance.
* `ntp`: Whether to use DHCP for setting NTP on the appliance.
* `mtu`: Whether to use DHCP for setting MTU on the appliance.
##### static
IPv4 static NIC configuration for the NIC.
* `address`: IPv4 Address of the network interface.
* `netmask`: Netmask of the network interface.
* `hostname`: NIC hostname.
* `snat`: Enable SNAT on this IP.
##### virtual_ip
Virtual IP to use for IPv4.
##### dhcp
IPv6 DHCP configuration for the NIC.
* `enabled`: Whether DHCP for IPv6 is enabled.
* `dns`: Whether to use DHCP for setting IPv6 DNS settings on the Appliance.
* `ntp`: Whether to use DHCP for setting NTP on the appliance.
* `mtu`: Whether to use DHCP for setting MTU on the appliance.
##### static
IPv6 static NIC configuration for the NIC.
* `address`: IPv6 Address of the network interface.
* `netmask`: Netmask of the network interface.
* `hostname`: NIC hostname.
* `snat`: Enable SNAT on this IP.
##### virtual_ip
Virtual IP to use for IPv6.
#### dns_servers
DNS Server addresses.
#### dns_domains
DNS Search domains.
#### routes
System route settings.
* `address`: (Required) Address to route. Example: 10.0.0.0.
* `netmask`: (Required) Netmask for the subnet to route. Example: 24.
* `gateway`: (Optional) Gateway to use for routing. Example: 10.0.0.254.
* `nic`: (Optional) NIC name to use for routing. Example: eth0.
### ntp_servers
Deprecated as of 4.3.0, use 'ntp' field instead. NTP servers to synchronize time.

### ntp
NTP configuration.

* `servers`:  (Optional) undefined
#### servers

* `hostname`: (Required) Hostname or IP of the NTP server. Example: 0.ubuntu.pool.ntp.org.
* `key_type`: (Optional) Type of key to use for secure NTP communication. ENUM: MD5,SHA,SHA1,SHA256,SHA512,RMD160.
* `key`: (Optional) Key to use for secure NTP communication.
### ssh_server
SSH server configuration.

* `enabled`:  (Optional)  default value `false` Whether the SSH Server is enabled on this appliance or not.
* `port`:  (Optional)  default value `22` SSH port.
* `allow_sources`:  (Optional) Source configuration to allow via iptables.
* `password_authentication`:  (Optional)  default value `true` Whether SSH allows password authentication or not.
#### allow_sources
Source configuration to allow via iptables.
* `address`: (Optional) IP address to allow connection. Example: 0.0.0.0,::.
* `netmask`: (Optional) Netmask to use with address for allowing connections. Example: 0.
* `nic`: (Optional) NIC name to accept connections on. Example: eth0.
### snmp_server
SNMP Server configuration.

* `enabled`:  (Optional)  default value `false` Whether the SNMP Server os enabled on this appliance or not.
* `tcp_port`:  (Optional) TCP port for SNMP Server. Example: 161.
* `udp_port`:  (Optional) UDP port for SNMP Server. Example: 161.
* `snmpd.conf`:  (Optional) Raw SNMP configuration.
* `allow_sources`:  (Optional) Source configuration to allow via iptables.
#### allow_sources
Source configuration to allow via iptables.
* `address`: (Optional) IP address to allow connection. Example: 0.0.0.0,::.
* `netmask`: (Optional) Netmask to use with address for allowing connections. Example: 0.
* `nic`: (Optional) NIC name to accept connections on. Example: eth0.
### healthcheck_server
Healthcheck Server configuration.

* `enabled`:  (Optional)  default value `false` Whether the Healthcheck Server is enabled on this appliance or not.
* `port`:  (Optional)  default value `5555` Port to connect for Healthcheck Server.
* `allow_sources`:  (Optional) Source configuration to allow via iptables.
#### allow_sources
Source configuration to allow via iptables.
* `address`: (Optional) IP address to allow connection. Example: 0.0.0.0,::.
* `netmask`: (Optional) Netmask to use with address for allowing connections. Example: 0.
* `nic`: (Optional) NIC name to accept connections on. Example: eth0.
### prometheus_exporter
Prometheus Exporter configuration.

* `enabled`:  (Optional)  default value `false` Whether the Prometheus Exporter is enabled on this appliance or not.
* `port`:  (Optional)  default value `5556` Port to connect for Prometheus Exporter.
* `allow_sources`:  (Optional) Source configuration to allow via iptables.
#### allow_sources
Source configuration to allow via iptables.
* `address`: (Optional) IP address to allow connection. Example: 0.0.0.0,::.
* `netmask`: (Optional) Netmask to use with address for allowing connections. Example: 0.
* `nic`: (Optional) NIC name to accept connections on. Example: eth0.
### ping
Rules for allowing ping.

* `allow_sources`:  (Optional) Source configuration to allow via iptables.
#### allow_sources
Source configuration to allow via iptables.
* `address`: (Optional) IP address to allow connection. Example: 0.0.0.0,::.
* `netmask`: (Optional) Netmask to use with address for allowing connections. Example: 0.
* `nic`: (Optional) NIC name to accept connections on. Example: eth0.
### log_server
Log Server settings. Log Server collects audit logs from all the appliances and stores them.

* `enabled`:  (Optional)  default value `false` Whether the Log Server is enabled on this appliance or not.
* `retention_days`:  (Optional)  default value `30` How many days of audit logs will be kept.
### controller
Controller settings.

* `enabled`:  (Optional)  default value `false` Whether the Controller is enabled on this appliance or not.
### gateway
Gateway settings.

* `enabled`:  (Optional)  default value `false` Whether the Gateway is enabled on this appliance or not.
* `vpn`:  (Optional) VPN configuration.
#### vpn
VPN configuration.
* `weight`: (Optional) Load balancing weight.
* `allow_destinations`: (Optional) Destinations to allow tunnels to.
### log_forwarder
LogForwarder settings. LogForwarder collects audit logs from the appliances in the given sites and sends them to the given endpoints.

* `enabled`:  (Optional)  default value `false` Whether the LogForwarder is enabled on this appliance or not.
* `elasticsearch`:  (Optional) undefined
* `tcp_clients`:  (Optional) TCP endpoints to connect and send the audit logs with the given format.
* `aws_kineses`:  (Optional) AWS Kinesis endpoints to connect and send the audit logs with the given format.
* `sites`:  (Optional) The sites to collect logs from and forward.
#### tcp_clients
TCP endpoints to connect and send the audit logs with the given format.
* `name`: (Required) Name of the endpoint. Example: Company SIEM.
* `host`: (Required) Hostname or the IP address of the endpoint. Example: siem.company.com.
* `port`: (Required) Port of the endpoint. Example: 8888.
* `format`: (Required) The format to send the audit logs. ENUM: json,syslog.
* `use_t_l_s`: (Optional) Whether to use TLS to connect to endpoint or not. If enabled, make sure the LogForwarder appliance trusts the certificate of the endpoint.
* `filter`: (Optional) JMESPath expression to filter audit logs to forward. Example: event_type=='authentication_succeeded'.
#### aws_kineses
AWS Kinesis endpoints to connect and send the audit logs with the given format.
#### sites
The sites to collect logs from and forward.
### connector
Connector settings.

* `enabled`:  (Optional)  default value `false` Whether the Connector is enabled on this appliance or not.
* `express_clients`:  (Optional) A list of Clients to run on the appliance with the given configuration. The Clients will get the necessary tokens automatically according to the Site assigned to this Appliance. Currently only one allowed.
* `advanced_clients`:  (Optional) A list of Clients to run on the appliance with the given configuration. Requires manual Policy configuration.
#### express_clients
A list of Clients to run on the appliance with the given configuration. The Clients will get the necessary tokens automatically according to the Site assigned to this Appliance. Currently only one allowed.
* `name`: (Required) Name for the Client. It will be mapped to the user claim 'clientName'. Example: Printers.
* `device_id`: (Optional) The device ID to assign to this Client. It will be used to generate device distinguished name. Example: 12699e27-b584-464a-81ee-5b4784b6d425.
* `allow_resources`: (Optional) A list of subnets to allow access via Client.
* `snat_to_resources`: (Optional) Use Source NAT for the resources.
#### advanced_clients
A list of Clients to run on the appliance with the given configuration. Requires manual Policy configuration.
* `name`: (Required) Name for the Client. It will be mapped to the user claim 'clientName'. Example: Printers.
* `device_id`: (Optional) The device ID to assign to this Client. It will be used to generate device distinguished name. Example: 12699e27-b584-464a-81ee-5b4784b6d425.
* `allow_resources`: (Optional) Source configuration to allow via iptables.
* `snat_to_tunnel`: (Optional) Use Source NAT for the Client tunnel.
### rsyslog_destinations
Rsyslog destination settings to forward appliance logs.

* `selector`:  (Optional)  default value `*.*` Rsyslog selector. Example: :msg, contains, "[AUDIT]".
* `template`:  (Optional)  default value `%HOSTNAME% %msg%` Rsyslog template to forward logs with. Example: %msg:9:$%
.
* `destination`:  (Optional) Rsyslog server destination. Example: @@10.10.10.2.
### hostname_aliases
Hostname aliases. They are added to the Appliance certificate as Subject Alternative Names so it is trusted using different IPs or hostnames. Requires manual certificate renewal to apply changes to the certificate.

### tags
Array of tags.




## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import appgate_appliance d3131f83-10d1-4abc-ac0b-7349538e8300
```
