# OVH Resource Provider

The OVH Resource Provider lets you manage [OVHcloud](https://www.ovhcloud.com/en/) resources.

<a href="https://github.com/scraly/pulumi-ovh/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/scraly/pulumi-ovh?logo=github&style=flat-square"></a>
[![GoDoc](https://godoc.org/github.com/scraly/pulumi-ovh?status.svg)](https://godoc.org/github.com/scraly/pulumi-ovh)
[![Go Report Card](https://goreportcard.com/badge/github.com/scraly/pulumi-ovh)](https://goreportcard.com/report/github.com/scraly/pulumi-ovh)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fdatabricks.svg)](https://www.npmjs.com/package/@pulumi/databricks)
[![Python version](https://badge.fury.io/py/pulumi-databricks.svg)](https://pypi.org/project/pulumi-databricks)
[![NuGet version](https://badge.fury.io/nu/pulumi.databricks.svg)](https://badge.fury.io/nu/pulumi.databricks)
<a href="https://gitpod.io/#https://github.com/scraly/pulumi-ovh"><img src="https://img.shields.io/badge/Contribute%20with-Gitpod-908a85?logo=gitpod" alt="Contribute with Gitpod"/></a>

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @scraly/pulumi-ovh
```

or `yarn`:

```bash
yarn add @scraly/pulumi-ovh
```

### Python

To use from Python, install using `pip`:

```bash
pip install scraly-pulumi-ovh
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/scraly/pulumi-ovh/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Scraly.PulumiPackage.Ovh
```

## Configuration

The following configuration points are available for the `Ovh` provider:

- `ovh:endpoint` (environment: `OVH_ENDPOINT`) - the Ovh endpoint, such `ovh-eu`
- `ovh:applicationKey` (environment: `OVH_APPLICATION_KEY`) - the Ovh application key
- `ovh:applicationSecret` (environment: `OVH_APPLICATION_SECRET`) - the Ovh application secret
- `ovh:consumerKey` (environment: `OVH_CONSUMER_KEY`) - the Ovh consumer key

## Upgrading

* install gh CLI
* Install upgrade-provider CLI
* create and retrieve your GitHub Personal access token (PAT)
* export it:

```bash
$ export GITHUB_TOKEN="<your-gh-pat>"
```

* execute the upgrade-provider CLI:

```bash
$ upgrade-provider scraly/pulumi-ovh
---- Setting Up Environment ----
- ✓ GOWORK="off": done
- ✓ PULUMI_MISSING_DOCS_ERROR="false": done
- ✓ PULUMI_CONVERT_EXAMPLES_CACHE_DIR="": done
---- Discovering Repository ----
- Ensure 'github.com/scraly/pulumi-ovh'
  - ✓ Expected Location: /workspace/pulumi-ovh
  - ✓ Downloading: skipped - already exists
  - ✓ Validating: done
- pull default branch
  - ✓ /usr/bin/git ls-remote --heads origin: done
  - ✓ finding default branch: main
  - ✓ /usr/bin/git fetch: done
  - ✓ /usr/bin/git checkout main: done
  - ✓ /usr/bin/git pull origin: done
- ✓ Repo kind: plain
- X Planning Provider Update: current upstream version 0.32.0 is greater than/ equal to the target version 0.32.0
```

In this case, OK nothing to do :)

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/ovh/api-docs/).
