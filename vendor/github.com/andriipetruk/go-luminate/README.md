# go-luminate
go-luminate is a Go client library for accessing the  [luminate.io](https://luminate.io/) API

go-luminate requires Go version 1.9 or greater.


## Usage ##

Before start to go you should get yours tenant name, client id and client secret.
For more detail please look  [https://luminatepublicapi.docs.apiary.io/#introduction/authorization](https://luminatepublicapi.docs.apiary.io/#introduction/authorization)

```go
  import   "github.com/andriipetruk/go-luminate/luminate"
```

Construct a new luminate client, then use the various services on the client to
access different parts of the  API. For example:

```go
client := goluminate.NewClient(ctx, ClientID, ClientSecret, TenantName)

// list all application
applist, _, err := client.ListApp(ctx, "")
```

More examples you will found in file [example.go](./example.go)


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
