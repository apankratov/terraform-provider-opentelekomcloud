---
subcategory: "APIGW"
---

Up-to-date reference of API arguments for API Gateway group service you can get at
`https://docs.otc.t-systems.com/api-gateway/api-ref/dedicated_gateway_apis_v2/group_response_management/index.html`.

# opentelekomcloud_apigw_response_v2

Manages an APIGW (API) custom response resource within OpenTelekomCloud.

## Example Usage

```hcl
variable "gateway_id" {}
variable "group_id" {}
variable "response_name" {}

resource "opentelekomcloud_apigw_response_v2" "test" {
  gateway_id = var.gateway_id
  group_id   = var.group_id
  name       = var.response_name

  rule {
    error_type  = "AUTHORIZER_FAILURE"
    body        = "{\"code\":\"$context.authorizer.frontend.code\",\"message\":\"$context.authorizer.frontend.message\"}"
    status_code = 401
  }
}
```

## Argument Reference

The following arguments are supported:
* `gateway_id` - (Required, String, ForceNew) Specifies the ID of the dedicated instance to which the API group and the
  API custom response belong.
  Changing this will create a new resource.

* `group_id` - (Required, String, ForceNew) Specifies the ID of the API group to which the API custom response
  belongs.
  Changing this will create a new resource.

* `name` - (Required, String) Specifies the name of the API custom response.
  The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.

* `rule` - (Optional, List) Specifies the API custom response rules definition.
  The [object](#custom_response_rule) structure is documented below.

<a name="custom_response_rule"></a>
The `rule` block supports:

* `error_type` - (Required, String) Specifies the error type of the API response rule.
  The valid values and the related default status code are as follows:
  + `ACCESS_DENIED`: (`403`) Access denied.
  + `AUTH_FAILURE`: (`401`) Authentication failed.
  + `AUTH_HEADER_MISSING`: (`401`) The identity source is missing.
  + `AUTHORIZER_CONF_FAILURE`: (`500`) There has been a custom authorizer error.
  + `AUTHORIZER_FAILURE`: (`500`) Custom authentication failed.
  + `AUTHORIZER_IDENTITIES_FAILURE`: (`401`) The identity source of the custom authorizer is invalid.
  + `BACKEND_TIMEOUT`: (`504`) Communication with the backend service timed out.
  + `BACKEND_UNAVAILABLE`: (`502`) The backend service is unavailable.
  + `NOT_FOUND`: (`404`) No API is found.
  + `REQUEST_PARAMETERS_FAILURE`: (`400`) The request parameters are incorrect.
  + `THROTTLED`: (`429`) The request was rejected due to request throttling.
  + `UNAUTHORIZED`: (`401`) The app you are using has not been authorized to call the API.
  + `DEFAULT_4XX`: (`NONE`) Another 4XX error occurred.
  + `DEFAULT_5XX`: (`NONE`) Another 5XX error occurred.
  + `THIRD_AUTH_CONF_FAILURE`: (`500`) Third-party authorizer configuration error.
  + `THIRD_AUTH_FAILURE`: (`401`) Third-party authentication failed.
  + `THIRD_AUTH_IDENTITIES_FAILURE`: (`401`) Identity source of the third-party authorizer is invalid.

* `body` - (Required, String) Specifies the body template of the API response rule, e.g.
  `{\"code\":\"$context.authorizer.frontend.code\",\"message\":\"$context.authorizer.frontend.message\"}`

* `status_code` - (Optional, Int) Specifies the HTTP status code of the API response rule.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the API custom response.
* `region` -  The region where the API custom response is located.
* `created_at` - The creation time of the API custom response.
* `updated_at` - The latest update time of the API custom response.

## Import

API Responses can be imported using their `name` and IDs of the APIGW dedicated instances and API groups to which the API
response belongs, separated by slashes, e.g.

```shell
$ terraform import opentelekomcloud_apigw_response_v2.test <gateway_id>/<group_id>/<name>
```
