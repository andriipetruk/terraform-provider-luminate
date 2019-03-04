# Terraform "luminate" provider

The *luminate* Terraform provider based on [go-luminate](https://github.com/andriipetruk/go-luminate/) Go client library for accessing the luminate.io API



## Installation

Note: the [Go](https://golang.org/) language compiler is required for building the provider.

Clone the repository into your `$GOPATH`:

```
$ mkdir -p $GOPATH/src/github.com/andriipetruk
$ git clone https://github.com/andriipetruk/terraform-provider-luminate $GOPATH/src/github.com/andriipetruk/terraform-provider-luminate
```

Enter the provider directory and build the provider

```
$ cd $GOPATH/src/github.com/andriipetruk/terraform-provider-luminate
$ make build
```

## Using

We have two way to use provider

1) Initialize the stack using option for terraform init

```
$ cd /path/to/terraform/stack
$ terraform init -plugin-dir=$GOPATH/bin
$ terraform plan
```

or

2) Installing provider to local plugin dir
```
$  cp ~/go/bin/terraform-provider-luminate ~/.terraform.d/plugins/

```


## Configuration

### Global provider options


*  `tenant_name`   (required – type string): Tenant name
*  `client_id`     (required – type string): Client ID
*  `client_secret` (required – type string): Client Secret

For more detail please look  [https://luminatepublicapi.docs.apiary.io/#introduction/authorization](https://luminatepublicapi.docs.apiary.io/#introduction/authorization)


### Resource "luminate_site"

* `site_name` (required – type string): Site name to be created


### Resource "luminate_connector"

* `connector_name` (required – type string): Connector name  to be created
* `site_id` (required – type string): Site id for link connecter on it

### Resource "luminate_app_http"

* `app_name` (required – type string): Application name to be created
* `internal_address` (required – type string): Application internal address format {protocol}://host.domain , example - http://test.local.com
* `site_id` (required – type string): Site id for link connecter on it

### Resource "luminate_app_ssh"

* `app_name` (required – type string): Application name  to be created
* `internal_address` (required – type string): Application internal address format tcp://host.domain:{port} , example - tcp://test.local.com:22
* `ssh_login` (required – type string): Login on host for ssh access 
* `site_id` (required – type string): Site id for link connecter on it

### Resource "luminate_app_tcp"

* `app_name` (required – type string): Application name  to be created
* `internal_address` (required – type string): Application internal address format host.domain , example - test.local.com
* `tcp_port` (required – type string): tcp port number for ssh access 
* `site_id` (required – type string): Site id for link connecter on it

## Example Usage

Using the following Terraform configuration will be created new site, connector and http application. Also will be deployed connector to kubernets cluster :

```

provider "luminate" {
  tenant_name   = "${var.luminate["tenant"]}"
  client_id     = "${var.luminate["client_id"]}"
  client_secret = "${var.luminate["client_secret"]}"
}


resource "luminate_site" "newSite" {
  depends_on = ["module.cluster_poc1"]
  site_name = "${var.luminate["site_name"]}"
}

resource "luminate_connector" "newConnector" {
  depends_on = ["module.cluster_poc1"]
  connector_name = "${var.luminate["connector_name"]}"
  site_id = "${luminate_site.newSite.id}"
}

resource "null_resource" "connector_install" {
  depends_on = ["module.cluster_poc1"]
  provisioner "local-exec" {
    command = "${luminate_connector.newConnector.install}"
  }
}

resource "luminate_app_http" "stackstorm-st2web" {
  app_name =  "${var.luminate-st2web["app_name"]}"
  internal_address = "${var.luminate-st2web["internal_address"]}"
  site_id = "${luminate_site.newSite.id}"
}

```

###  Testing ###

Not ready yet.

## Roadmap ##

Will be added

## Contributing ##
Any PR are welcome.

## Versioning ##

In general, go-luminate follows [semver](https://semver.org/) as closely as we
can for tagging releases of the package. 

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.