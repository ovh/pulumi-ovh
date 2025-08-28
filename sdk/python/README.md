# OVH Resource Provider

The OVH Resource Provider lets you manage [OVHcloud](https://www.ovhcloud.com/en/) resources.

<a href="https://github.com/ovh/pulumi-ovh/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/ovh/pulumi-ovh?logo=github&style=flat-square"></a>
[![GoDoc](https://godoc.org/github.com/ovh/pulumi-ovh?status.svg)](https://pkg.go.dev/github.com/ovh/pulumi-ovh/sdk/v2)
[![NPM version](https://badge.fury.io/js/@ovhcloud%2Fpulumi-ovh.svg)](https://badge.fury.io/js/@ovhcloud%2Fpulumi-ovh)
[![](https://img.shields.io/npm/dm/@ovhcloud/pulumi-ovh)](https://www.npmjs.com/package/@ovhcloud/pulumi-ovh)
[![PyPI version](https://badge.fury.io/py/pulumi-ovh.svg)](https://badge.fury.io/py/pulumi-ovh)
[![](https://img.shields.io/pypi/dm/pulumi-ovh)](https://pypi.org/project/pulumi-ovh/)
[![NuGet version](https://badge.fury.io/nu/Pulumi.Ovh.svg)](https://badge.fury.io/nu/Pulumi.Ovh)
[![](https://img.shields.io/nuget/dt/Pulumi.Ovh)](https://www.nuget.org/packages/Pulumi.Ovh/)
[![Maven Central](https://img.shields.io/maven-central/v/com.ovhcloud.pulumi.ovh/pulumi-ovh)](https://central.sonatype.com/artifact/com.ovhcloud.pulumi.ovh/pulumi-ovh)
<a href="https://gitpod.io/#https://github.com/ovh/pulumi-ovh"><img src="https://img.shields.io/badge/Contribute%20with-Gitpod-908a85?logo=gitpod" alt="Contribute with Gitpod"/></a>

## Usage

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @ovhcloud/pulumi-ovh
```

or `yarn`:

```bash
yarn add @ovhcloud/pulumi-ovh
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-ovh
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/ovh/pulumi-ovh/sdk/v2/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Ovh
```

### Java

To use from Java, add the dependency below to your `pom.xml` file:

```bash
<dependency>
    <groupId>com.ovhcloud.pulumi.ovh</groupId>
    <artifactId>pulumi-ovh</artifactId>
    <version>[2.0.0,)</version>
</dependency>
```

## Configuration

The following configuration points are available for the `Ovh` provider:

- `ovh:endpoint` (environment: `OVH_ENDPOINT`) - the Ovh endpoint, such `ovh-eu`
- `ovh:clientId` (environment: `OVH_CLIENT_ID`) - the Ovh OAuth2 client ID
- `ovh:clientSecret` (secret) (environment: `OVH_CLIENT_SECRET`) - the Ovh OAuth2 client secret
- `ovh:accessToken` (environment: `OVH_ACCESS_TOKEN`) - the Ovh access token
- `ovh:applicationKey` (environment: `OVH_APPLICATION_KEY`) - the Ovh application key
- `ovh:applicationSecret` (environment: `OVH_APPLICATION_SECRET`) - the Ovh application secret
- `ovh:consumerKey` (environment: `OVH_CONSUMER_KEY`) - the Ovh consumer key

## Reference

For further information, visit [OVH in the Pulumi Registry](https://www.pulumi.com/registry/packages/ovh/)
or for detailed API reference documentation, visit [OVH API Docs in the Pulumi Registry](https://www.pulumi.com/registry/packages/ovh/api-docs/).
