# OVH Resource Provider

The OVH Resource Provider lets you manage [OVHcloud](https://www.ovhcloud.com/en/) resources.

<a href="https://github.com/scraly/pulumi-ovh/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/scraly/pulumi-ovh?logo=github&style=flat-square"></a>
[![GoDoc](https://godoc.org/github.com/scraly/pulumi-ovh?status.svg)](https://godoc.org/github.com/scraly/pulumi-ovh)
[![NPM version](https://badge.fury.io/js/@scraly%2Fpulumi-ovh.svg)](https://badge.fury.io/js/@scraly%2Fpulumi-ovh)
[![PyPI version](https://badge.fury.io/py/scraly-pulumi-ovh.svg)](https://badge.fury.io/py/scraly-pulumi-ovh)
[![NuGet version](https://badge.fury.io/nu/Scraly.PulumiPackage.Ovh.svg)](https://badge.fury.io/nu/Scraly.PulumiPackage.Ovh)
<a href="https://gitpod.io/#https://github.com/scraly/pulumi-ovh"><img src="https://img.shields.io/badge/Contribute%20with-Gitpod-908a85?logo=gitpod" alt="Contribute with Gitpod"/></a>

## Usage

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

## Upgrading/Sync with existing Terraform provider

* install gh CLI
* Install upgrade-provider CLI
* create and retrieve your GitHub Personal access token (PAT)
* export it:

```bash
$ export GITHUB_TOKEN="<your-gh-pat>"
```

* execute the upgrade-provider CLI:

If the provider is already up to date:

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

If the provider need to be upgraded:
```bash
$ upgrade-provider scraly/pulumi-ovh
---- Setting Up Environment ----
- ✓ GOWORK="off": done
- ✓ PULUMI_MISSING_DOCS_ERROR="false": done
- ✓ PULUMI_CONVERT_EXAMPLES_CACHE_DIR="": done
---- Discovering Repository ----
- Ensure 'github.com/scraly/pulumi-ovh'
  - ✓ Expected Location: /workspaces/pulumi-ovh
  - ✓ Downloading: skipped - already exists
  - ✓ Validating: done
- pull default branch
  - ✓ /usr/bin/git ls-remote --heads origin: done
  - ✓ finding default branch: main
  - ✓ /usr/bin/git fetch: done
  - ✓ /usr/bin/git checkout main: done
  - ✓ /usr/bin/git pull origin: done
- ✓ Repo kind: plain
- ✓ Planning Provider Update: 0.32.0 -> 0.33.0
- ✓ Planning Bridge Update: v3.53.0 -> v3.57.0
- ✓ Planning Plugin SDK Update: v2.26.1 -> 2.27.0
- ✓ Planning Plugin Framework Update:  -> 0.15.2
---- Update Artifacts ----
- Ensure Branch
  - ✓ /usr/bin/git branch: done
  - ✓ Already exists: no
  - ✓ /usr/bin/git checkout -b upgrade-terraform-provider-ovh-to-v0.33.0: done
  - ✓ /usr/bin/git checkout upgrade-terraform-provider-ovh-to-v0.33.0: done
- Update TF Provider
  - ✓ Update TF Plugin SDK Fork: updated
  - ✓ /usr/local/go/bin/go mod tidy: done
  - ✓ Lookup Tag SHA: 6f0cc4b7d791d1682b613e2c2cbd9a48d4267686
  - ✓ /usr/local/go/bin/go get github.com/ovh/terraform-provider-ovh@6f0cc4b7d791d1682...: done
- ✓ /usr/local/go/bin/go get github.com/pulumi/pulumi-terraform-bridge/v3@v3.57.0: done
- ✓ /usr/local/go/bin/go get github.com/pulumi/pulumi-terraform-bridge/pf@v0.15.2: done
- ✓ /usr/local/go/bin/go mod tidy: done
- ✓ /usr/local/go/bin/go mod tidy: done
- ✓ /bin/echo Plugins not removed.: done
- X /usr/bin/make tfgen: exit status 2:
[resource plugin random-4.8.2] installing
warning: A new version of Pulumi is available. To upgrade from version '3.78.1' to '3.79.0', run 
   $ curl -sSL https://get.pulumi.com | sh
or visit https://pulumi.com/docs/install/ for manual instructions and release notes.



```

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/ovh/api-docs/).
