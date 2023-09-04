---
title: OVH Installation & Configuration
meta_desc: Information on how to install the OVH provider.
layout: installation
---

## Information

Note that the [lbrlabs Pulumi OVH provider](https://github.com/lbrlabs/pulumi-ovh) is replaced by this official one.

## Installation

The Pulumi `OVH` provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@ovh-devrelteam/pulumi-ovh`](https://www.npmjs.com/package/@ovh-devrelteam/pulumi-ovh)
* Python: [`pulumi_ovh`](https://pypi.org/project/pulumi-ovh/)
* Go: [`github.com/ovh/pulumi-ovh/sdk/go/ovh`](https://pkg.go.dev/github.com/ovh/pulumi-ovh/sdk)
* .NET: [`Pulumi.Ovh`](https://www.nuget.org/packages/Pulumi.Ovh)

### Provider Binary

The OVH provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource ovh <version>
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi OVH provider, you need to have OVH credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Grafana:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export OVH_ENDPOINT="<the Ovh endpoint, for example ovh-eu>"
$ export OVH_APPLICATION_KEY="<the Ovh application key>"
$ export OVH_APPLICATION_SECRET="<the Ovh application secret>"
$ export OVH_CONSUMER_KEY="<the Ovh consumer key>"
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export OVH_ENDPOINT="<the Ovh endpoint, for example ovh-eu>"
$ export OVH_APPLICATION_KEY="<the Ovh application key>"
$ export OVH_APPLICATION_SECRET="<the Ovh application secret>"
$ export OVH_CONSUMER_KEY="<the Ovh consumer key>"
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:OVH_ENDPOINT = "<the Ovh endpoint, for example ovh-eu>"
> $env:OVH_APPLICATION_KEY = "<the Ovh application key>"
> $env:OVH_APPLICATION_SECRET = "<the Ovh application secret>"
> $env:OVH_CONSUMER_KEY = "<the Ovh consumer key>"
```

{{% /choosable %}}
{{< /chooser >}}
