---
layout: "appgate"
page_title: "APPGATE: appgate_entitlement"
sidebar_current: "docs-appgate-resource-entitlement"
description: |-
   Create a new Entitlement.
---

# appgate_entitlement

Create a new Entitlement..

## Example Usage

```hcl

data "appgate_site" "default_site" {
  site_name = "Default site"
}
resource "appgate_entitlement" "ping_entitlement" {
  name = "test entitlement"
  site = data.appgate_site.default_site.id
  # site = appgate_site.gbg_site.id
  conditions = [
    data.appgate_condition.always.id
  ]

  actions {
    subtype = "icmp_up"
    action  = "allow"
    # https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml#icmp-parameters-types
    types = ["0-16"]
    hosts = [
      "10.0.0.1",
      "10.0.0.0/24",
      "hostname.company.com",
      "dns://hostname.company.com",
      "aws://security-group:accounting"
    ]
  }

}


```

## Argument Reference

The following arguments are supported:


* `disabled`: (Optional) If true, the Entitlement will be disregarded during authorization.
* `site`: (Required) ID of the Site for this Entitlement.
* `site_name`: (Optional) Name of the Site for this Entitlement. For convenience only.
* `condition_logic`: (Optional) Whether all the Conditions must succeed to have access to this Entitlement or just one.
* `conditions`: (Required) List of Condition IDs applies to this Entitlement.
* `actions`: (Required) List of all IP Access actions in this Entitlement.
* `app_shortcuts`: (Optional) Array of App Shortcuts.
* `app_shortcut_scripts`: (Optional) List of Entitlement Script IDs used for creating App Shortcuts dynamically.
* `id`: (Required) ID of the object.
* `name`: (Required) Name of the object.
* `notes`: (Optional) Notes for the object. Used for documentation purposes.
* `tags`: (Optional) Array of tags.


### conditions
List of Condition IDs applies to this Entitlement.

### actions
List of all IP Access actions in this Entitlement.

* `subtype`:  (Optional)  Enum values: `icmp_up,icmp_down,icmpv6_up,icmpv6_down,udp_up,udp_down,tcp_up,tcp_down,ah_up,ah_down,esp_up,esp_down,gre_up,gre_down,http_up`Type of the IP Access action.
* `action`:  (Optional)  Enum values: `allow,block,alert`Applied action to the traffic.
* `hosts`:  (Optional) Hosts to apply the action to. See admin manual for possible values. Example: 10.0.0.1,10.0.0.0/24,hostname.company.com,dns://hostname.company.com,aws://security-group:accounting.
* `ports`:  (Optional) Destination port. Multiple ports can be entered comma separated. Port ranges can be entered dash separated. Only valid for tcp and udp subtypes Example: 80,1024-65535.
* `types`:  (Optional) ICMP type. Only valid for icmp subtypes. Example: 0-255.
* `monitor`:  (Optional) Only available for tcp_up subtype. If enabled, Gateways will monitor this action for responsiveness and act accordingly. See admin manual for more details.
#### hosts
Hosts to apply the action to. See admin manual for possible values.
#### ports
Destination port. Multiple ports can be entered comma separated. Port ranges can be entered dash separated. Only valid for tcp and udp subtypes
#### types
ICMP type. Only valid for icmp subtypes.
#### monitor
Only available for tcp_up subtype. If enabled, Gateways will monitor this action for responsiveness and act accordingly. See admin manual for more details.
* `enabled`: (Optional) undefined
* `timeout`: (Optional) Number of seconds the action needs to be unresponsive to mark the action unhealthy.
### app_shortcuts
Array of App Shortcuts.

* `name`:  (Optional) Name for the App Shortcut which will be visible on the Client UI. Example: Accounting Software.
* `description`:  (Optional) Description for the App Shortcut which will be visible on the Client UI. Example: This app will get you to the Accounting Software..
* `url`:  (Optional) The URL that will be triggered on the OS to be handled. For example, an HTTPS URL will start the browser for the given URL. Example: https://service.company.com.
* `color_code`:  (Optional)  default value `1` The code of the published app on the client.
- 1: Light Green
- 2: Green
- 3: Indigo
- 4: Deep Purple
- 5: Yellow
- 6: Lime
- 7: Light Blue
- 8: Blue
- 9: Amber
- 10: Orange
- 11: Cyan
- 12: Teal
- 13: Deep Orange
- 14: Red
- 15: Gray
- 16: Brown
- 17: Pink
- 18: Purple
- 19: Blue Gray
- 20: Near Black

### app_shortcut_scripts
List of Entitlement Script IDs used for creating App Shortcuts dynamically.

### tags
Array of tags.




## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import appgate_entitlement d3131f83-10d1-4abc-ac0b-7349538e8300
```
