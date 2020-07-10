---
layout: "appgate"
page_title: "APPGATE: appgate_condition"
sidebar_current: "docs-appgate-resource-condition"
description: |-
   Create a new Condition.
---

# appgate_condition

Create a new Condition..

## Example Usage

```hcl

resource "appgate_condition" "test_condition" {
  name = "teraform-example-condition"
  tags = [
    "terraform",
    "api-created"
  ]

  expression = <<-EOF
var result = false;
/*password*/
if (claims.user.hasPassword('test', 60)) {
  return true;
}
/*end password*/
return result;
EOF

  repeat_schedules = [
    "1h",
    "13:32"
  ]


}

```

## Argument Reference

The following arguments are supported:


* `expression`: (Required) Boolean expression in JavaScript.
* `repeat_schedules`: (Optional) A list of schedules that decides when to reevaluate the Condition. All the scheduled times will be effective. One will not override the other. - It can be a time of the day, e.g. 13:00, 10:25, 2:10 etc. - It can be one of the predefined
  intervals, e.g. 1m, 5m, 15m, 1h. These intervals
  will be always rounded up, i.e. if it's 15m and the
  time is 12:07 when the Condition is evaluated
  first, then the next evaluation will occur at
  12:15, and the next one will be at
  12:30 and so on.
* `remedy_methods`: (Optional) The remedy methods that will be triggered if the evaluation fails.
* `id`: (Required) ID of the object.
* `name`: (Required) Name of the object.
* `notes`: (Optional) Notes for the object. Used for documentation purposes.
* `tags`: (Optional) Array of tags.


### repeat_schedules
A list of schedules that decides when to reevaluate the Condition. All the scheduled times will be effective. One will not override the other. - It can be a time of the day, e.g. 13:00, 10:25, 2:10 etc. - It can be one of the predefined
  intervals, e.g. 1m, 5m, 15m, 1h. These intervals
  will be always rounded up, i.e. if it's 15m and the
  time is 12:07 when the Condition is evaluated
  first, then the next evaluation will occur at
  12:15, and the next one will be at
  12:30 and so on.

### remedy_methods
The remedy methods that will be triggered if the evaluation fails.

* `type`:  (Optional)  Enum values: `DisplayMessage,OtpAuthentication,PasswordAuthentication,Reason`remedy method type.
* `message`:  (Optional) Message to be shown to the user. Required for all remedy method. Example: This resoure requires you to enter your password again.
* `claim_suffix`:  (Optional) Suffix to be added to the claim. Required for OtpAuthentication, PasswordAuthentication and Reason remedy methods. Example: test.
* `provider_id`:  (Optional) MFA Provider Id. Required for OtpAuthentication remedy method. Example: 4c07bc67-57ea-42dd-b702-c2d6c45419fc.
### tags
Array of tags.




## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import appgate_condition d3131f83-10d1-4abc-ac0b-7349538e8300
```
