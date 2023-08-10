# OVH Resource Provider

The OVH Resource Provider lets you manage [OVHcloud](https://www.ovhcloud.com/) resources.

<a href="https://github.com/dirien/pulumi-ovh/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/dirien/pulumi-ovh?logo=github&style=flat-square"></a>
[![GoDoc](https://godoc.org/github.com/dirien/pulumi-ovh?status.svg)](https://godoc.org/github.com/dirien/pulumi-ovh)
[![Go Report Card](https://goreportcard.com/badge/github.com/dirien/pulumi-ovh)](https://goreportcard.com/report/github.com/dirien/pulumi-ovh)
<a href="https://gitpod.io/#https://github.com/dirien/pulumi-ovh"><img src="https://img.shields.io/badge/Contribute%20with-Gitpod-908a85?logo=gitpod" alt="Contribute with Gitpod"/></a>

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @dirien/pulumi-ovh
```

or `yarn`:

```bash
yarn add @dirien/pulumi-ovh
```

### Python

To use from Python, install using `pip`:

```bash
pip install lbrlabs-pulumi-ovh
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/lbrlabs/pulumi-ovh/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Lbrlabs.PulumiPackage.Ovh
```

## Configuration

The following configuration points are available for the `Ovh` provider:

- `ovh:endpoint` (environment: `OVH_ENDPOINT`) - the Ovh endpoint, such `ovh-us`
- `ovh:applicationKey` (environment: `OVH_APPLICATION_KEY`) - the Ovh application key
- `ovh:applicationSecret` (environment: `OVH_APPLICATION_SECRET`) - the Ovh application secret
- `ovh:consumerKey` (environment: `OVH_CONSUMER_KEY`) - the Ovh consumer key

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/ovh/api-docs/).
